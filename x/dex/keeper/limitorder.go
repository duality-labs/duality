package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type LimitOrderTranche struct {
	types.LimitOrderTranche
	priceInToOut sdk.Dec
	priceOutToIn sdk.Dec
}

// Tranche domain object constructor
func NewLimitOrderTranche(tokenIn string, token0 string, pairId string, tickIndex int64, trancheIndex uint64, reservesIn sdk.Int, reservesOut sdk.Int, totalIn sdk.Int, totalOut sdk.Int) (*LimitOrderTranche, error) {
	// memoize prices
	var priceOutToIn sdk.Dec
	var err error
	if tokenIn == token0 {
		priceOutToIn, err = CalcPrice1To0(tickIndex)
		if err != nil {
			return nil, err
		}
	} else {
		priceOutToIn, err = CalcPrice0To1(tickIndex)
		if err != nil {
			return nil, err
		}
	}
	priceInToOut := sdk.OneDec().Quo(priceOutToIn)

	return &LimitOrderTranche{
		types.LimitOrderTranche{
			TrancheIndex:     trancheIndex,
			TickIndex:        tickIndex,
			TokenIn:          tokenIn,
			PairId:           pairId,
			ReservesTokenIn:  reservesIn,
			ReservesTokenOut: reservesOut,
			TotalTokenIn:     totalIn,
			TotalTokenOut:    totalOut,
		},
		priceInToOut,
		priceOutToIn,
	}, nil
}

// Zero-initialized tranche object constructor
func NewLimitOrderTrancheDefault(tokenIn string, token0 string, pairId string, tickIndex int64, trancheIndex uint64) (*LimitOrderTranche, error) {
	return NewLimitOrderTranche(
		tokenIn,
		token0,
		pairId,
		tickIndex,
		trancheIndex,
		sdk.ZeroInt(),
		sdk.ZeroInt(),
		sdk.ZeroInt(),
		sdk.ZeroInt(),
	)
}

// Fetch tranche from keeper's KVStore and return error if not found
func NewLimitOrderTrancheFromKeeper(ctx sdk.Context, k Keeper, tokenIn string, token0 string, pairId string, tickIndex int64, trancheIndex uint64) (tranche *LimitOrderTranche, err error) {
	trancheProto, found := k.GetLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, trancheIndex)
	if !found {
		return nil, types.ErrValidLimitOrderMapsNotFound
	}
	tranche, err = NewLimitOrderTranche(
		trancheProto.TokenIn,
		token0,
		trancheProto.PairId,
		trancheProto.TickIndex,
		trancheProto.TrancheIndex,
		trancheProto.ReservesTokenIn,
		trancheProto.ReservesTokenOut,
		trancheProto.TotalTokenIn,
		trancheProto.TotalTokenOut,
	)
	if err != nil {
		return nil, err
	}
	return tranche, nil
}

// Fetch tranche from keeper's KVStore and zero-initialize if not found
func NewInitLimitOrderTrancheFromKeeper(ctx sdk.Context, k Keeper, tokenIn string, token0 string, pairId string, tickIndex int64, trancheIndex uint64) (tranche *LimitOrderTranche, err error) {
	tranche, err = NewLimitOrderTrancheFromKeeper(ctx, k, tokenIn, token0, pairId, tickIndex, trancheIndex)
	// if not found in KVStore, call zero-initializing constructor
	if err == types.ErrValidLimitOrderMapsNotFound {
		tranche, err = NewLimitOrderTrancheDefault(tokenIn, token0, pairId, tickIndex, trancheIndex)
		if err != nil {
			return nil, err
		}
		k.SetLimitOrderTranche(ctx, tranche.LimitOrderTranche)
	} else if err != nil {
		return nil, err
	}
	return tranche, nil
}

func (tranche *LimitOrderTranche) Save(ctx sdk.Context, k Keeper) {
	k.SetLimitOrderTranche(ctx, tranche.LimitOrderTranche)
}

func (tranche *LimitOrderTranche) PartiallyFilled() bool {
	return tranche.ReservesTokenIn.LT(tranche.TotalTokenIn)
}

func (tranche *LimitOrderTranche) PlaceLimitOrder(amountIn sdk.Int, trancheUser *LimitOrderTrancheUser) {
	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Add(amountIn)
	tranche.TotalTokenIn = tranche.TotalTokenIn.Add(amountIn)
	trancheUser.PlaceLimitOrder(amountIn)
}

func (tranche *LimitOrderTranche) CancelLimitOrder(sharesOut sdk.Int, trancheUser *LimitOrderTrancheUser) error {
	totalTokenInDec := sdk.NewDecFromInt(tranche.TotalTokenIn)
	totalTokenOutDec := sdk.NewDecFromInt(tranche.TotalTokenOut)
	filledAmount := tranche.priceOutToIn.Mul(totalTokenOutDec)
	ratioNotFilled := totalTokenInDec.Sub(filledAmount).Quo(totalTokenInDec)
	maxUserAllowedToCancel := trancheUser.SharesOwned.ToDec().Mul(ratioNotFilled).TruncateInt()
	totalUserAttemptingToCancel := trancheUser.SharesCancelled.Add(sharesOut)

	if totalUserAttemptingToCancel.GT(maxUserAllowedToCancel) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	if totalUserAttemptingToCancel.Add(trancheUser.SharesWithdrawn).GT(trancheUser.SharesOwned) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(sharesOut)

	tranche.ReservesTokenIn = tranche.ReservesTokenIn.Sub(sharesOut)
	return nil
}

func (tranche *LimitOrderTranche) WithdrawFilledLimitOrder() error {
	// reservesTokenOutDec := sdk.NewDecFromInt(tranche.ReservesTokenOut)
	// amountFilled := priceLimitOutToIn.MulInt(tranche.TotalTokenOut)
	// ratioFilled := amountFilled.QuoInt(tranche.TotalTokenIn)
	// maxAllowedToWithdraw := sdk.MinInt(
	// ratioFilled.MulInt(trancheUser.SharesOwned).TruncateInt(), // cannot withdraw more than what's been filled
	// trancheUser.SharesOwned.Sub(trancheUser.SharesCancelled),  // cannot withdraw more than what you own
	// )
	// amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)
	//
	// amountOutTokenOut := priceLimitInToOut.MulInt(amountOutTokenIn)
	//
	// trancheUser.SharesWithdrawn = maxAllowedToWithdraw
	// k.SetLimitOrderTrancheUser(ctx, trancheUser)
	//
	// tranche.ReservesTokenOut = reservesTokenOutDec.Sub(amountOutTokenOut).TruncateInt()
	// k.SetLimitOrderTranche(ctx, tranche)
	return nil
}

func (tranche *LimitOrderTranche) Swap() error {
	// reservesTokenOut := &tranche.ReservesTokenIn
	// fillTokenIn := &tranche.ReservesTokenOut
	// totalTokenIn := &tranche.TotalTokenOut
	// amountFilledTokenOut := priceInToOut.Mul(amountRemainingTokenIn).TruncateInt()
	//
	// if reservesTokenOut.LTE(amountFilledTokenOut) {
	// amountOut = amountOut.Add(*reservesTokenOut)
	// amountFilledTokenIn := priceOutToIn.MulInt(*reservesTokenOut)
	// amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
	// *reservesTokenOut = sdk.ZeroInt()
	// *fillTokenIn = fillTokenIn.Add(amountFilledTokenIn.TruncateInt())
	// *totalTokenIn = totalTokenIn.Add(amountFilledTokenIn.TruncateInt())
	// } else {
	// amountOut = amountOut.Add(amountFilledTokenOut)
	// *fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn.TruncateInt())
	// *totalTokenIn = totalTokenIn.Add(amountRemainingTokenIn.TruncateInt())
	// *reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
	// amountRemainingTokenIn = sdk.ZeroDec()
	// }
	return nil
}

type LimitOrderTrancheUser struct {
	types.LimitOrderTrancheUser
}

func (loUser *LimitOrderTrancheUser) PlaceLimitOrder(shares sdk.Int) {
	loUser.SharesOwned = loUser.SharesOwned.Add(shares)
}

func (loUser *LimitOrderTrancheUser) Save(ctx sdk.Context, k Keeper) {
	k.SetLimitOrderTrancheUser(ctx, loUser.LimitOrderTrancheUser)
}

func (k Keeper) GetUserShares(goCtx context.Context, lo *LimitOrderTranche, user string) LimitOrderTrancheUser {
	return LimitOrderTrancheUser{
		k.GetOrInitLimitOrderTrancheUser(goCtx,
			lo.LimitOrderTranche.PairId,
			lo.LimitOrderTranche.TickIndex,
			lo.LimitOrderTranche.TokenIn,
			lo.LimitOrderTranche.TrancheIndex,
			user),
	}
}

func (k Keeper) GetCurrentPlaceTranche(ctx sdk.Context, tokenIn string, token0 string, pairId string, tick *types.Tick, placeTrancheIndex *uint64) (tranche *LimitOrderTranche, err error) {
	tranche, err = NewInitLimitOrderTrancheFromKeeper(ctx, k, tokenIn, token0, pairId, tick.TickIndex, *placeTrancheIndex)
	if tranche.PartiallyFilled() {
		*placeTrancheIndex++
		k.SetTick(ctx, pairId, *tick)
		tranche, err = NewInitLimitOrderTrancheFromKeeper(ctx, k, tokenIn, token0, pairId, tick.TickIndex, *placeTrancheIndex)
	}
	return tranche, err
}

func GetFillAndPlaceIndexes(
	tokenIn string,
	token0 string,
	pair *types.TradingPair,
	tick *types.Tick,
) (fillTrancheIndex *uint64, placeTrancheIndex *uint64, err error) {
	if tokenIn == token0 {
		if tick.TickIndex > pair.CurrentTick0To1 {
			err = types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche0To1.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche0To1.PlaceTrancheIndex
	} else {
		if tick.TickIndex < pair.CurrentTick1To0 {
			err = types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderTranche1To0.FillTrancheIndex
		placeTrancheIndex = &tick.LimitOrderTranche1To0.PlaceTrancheIndex
	}
	return fillTrancheIndex, placeTrancheIndex, err
}
