package keeper

import (
	"context"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/duality-labs/duality/utils"
	"github.com/duality-labs/duality/x/dex/types"
)

///////////////////////////////////////////////////////////////////////////////
//                           GETTERS & INITIALIZERS                          //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) TokenInit(ctx sdk.Context, address string) {
	_, found := k.GetTokenMap(ctx, address)
	if !found {
		tokenIndex := k.GetTokensCount(ctx)
		newTokenCount := tokenIndex + 1
		// TODO: Consolidate TokenMap and Tokens into one type
		k.SetTokenMap(ctx, types.TokenMap{Address: address, Index: int64(tokenIndex)})
		k.AppendTokens(ctx, types.Tokens{Address: address, Id: tokenIndex})
		k.SetTokensCount(ctx, newTokenCount)
	}
}

// Handles initializing a new pair (token0/token1) if not found, adds token0, token1 to global list of tokens active on the dex
func (k Keeper) GetOrInitPair(ctx sdk.Context, token0 string, token1 string) types.TradingPair {
	k.TokenInit(ctx, token0)
	k.TokenInit(ctx, token1)
	pairId := CreatePairId(token0, token1)
	pair, found := k.GetTradingPair(ctx, pairId)
	if !found {
		pair = types.TradingPair{
			PairId:          &types.PairId{Token0: token0, Token1: token1},
			CurrentTick0To1: math.MaxInt64,
			CurrentTick1To0: math.MinInt64,
		}
		k.SetTradingPair(ctx, pair)
	}
	return pair
}

func (k Keeper) GetOrInitPoolReserves(ctx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64, fee uint64) (*types.PoolReserves, error) {
	tickLiq, tickFound := k.GetPoolReserves(
		ctx,
		pairId,
		tokenIn,
		tickIndex,
		fee,
	)
	if tickFound {
		return tickLiq, nil
	} else if IsTickOutOfRange(tickIndex) {
		return nil, types.ErrTickOutsideRange
	} else {
		return &types.PoolReserves{
			PairId:    pairId,
			TokenIn:   tokenIn,
			TickIndex: tickIndex,
			Fee:       fee,
			Reserves:  sdk.ZeroInt(),
		}, nil
	}

}

func NewLimitOrderTranche(pairId *types.PairId, tokenIn string, tickIndex int64, trancheIndex uint64) (types.LimitOrderTranche, error) {
	if IsTickOutOfRange(tickIndex) {
		return types.LimitOrderTranche{}, types.ErrTickOutsideRange
	}
	return types.LimitOrderTranche{
		PairId:           pairId,
		TokenIn:          tokenIn,
		TickIndex:        tickIndex,
		TrancheIndex:     trancheIndex,
		ReservesTokenIn:  sdk.ZeroInt(),
		ReservesTokenOut: sdk.ZeroInt(),
		TotalTokenIn:     sdk.ZeroInt(),
		TotalTokenOut:    sdk.ZeroInt(),
	}, nil

}

func (k Keeper) GetOrInitLimitOrderTrancheUser(
	goCtx context.Context,
	pairId *types.PairId,
	tickIndex int64,
	tokenIn string,
	currentLimitOrderKey uint64,
	receiver string,
) types.LimitOrderTrancheUser {
	ctx := sdk.UnwrapSDKContext(goCtx)

	UserShareData, UserShareDataFound := k.GetLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, currentLimitOrderKey, receiver)

	if !UserShareDataFound {
		return types.LimitOrderTrancheUser{
			Count:           currentLimitOrderKey,
			Address:         receiver,
			SharesOwned:     sdk.ZeroInt(),
			SharesWithdrawn: sdk.ZeroInt(),
			SharesCancelled: sdk.ZeroInt(),
			TickIndex:       tickIndex,
			Token:           tokenIn,
			PairId:          pairId,
		}
	}

	return UserShareData
}

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetCurrTick1To0(ctx sdk.Context, pairId *types.PairId) (tickIdx int64, found bool) {

	ti := k.NewTickIterator(ctx, pairId, pairId.Token0)

	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}
	return math.MinInt64, false

}

func (k Keeper) GetCurrTick0To1(ctx sdk.Context, pairId *types.PairId) (tickIdx int64, found bool) {
	ti := k.NewTickIterator(ctx, pairId, pairId.Token1)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return tick.TickIndex(), true
		}
	}

	return math.MaxInt64, false
}

func (k Keeper) IsBehindEnemyLines(ctx sdk.Context, pairId *types.PairId, tokenIn string, tickIndex int64) bool {
	if tokenIn == pairId.Token0 {
		curr0To1, _ := k.GetCurrTick0To1(ctx, pairId)
		if tickIndex >= curr0To1 {
			return true
		}
	} else {

		curr1To0, _ := k.GetCurrTick1To0(ctx, pairId)
		if tickIndex <= curr1To0 {
			return true
		}
	}
	return false
}

// Balance trueAmount1 to the pool ratio
func CalcTrueAmounts(
	centerTickPrice1To0 sdk.Dec,
	lowerReserve0 sdk.Int,
	upperReserve1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
	totalShares sdk.Int,
) (trueAmount0 sdk.Int, trueAmount1 sdk.Int, sharesMinted sdk.Int) {
	lowerReserve0Dec := lowerReserve0.ToDec()
	upperReserve1Dec := upperReserve1.ToDec()
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()

	// See spec: https://www.notion.so/dualityxyz/Autoswap-Spec-e856fa7b2438403c95147010d479b98c
	if upperReserve1Dec.GT(sdk.ZeroDec()) {
		trueAmount0 = sdk.MinDec(
			amount0Dec,
			amount1Dec.Mul(lowerReserve0Dec).Quo(upperReserve1Dec)).TruncateInt()
	} else {
		trueAmount0 = amount0
	}

	if lowerReserve0Dec.GT(sdk.ZeroDec()) {
		trueAmount1 = sdk.MinDec(
			amount1Dec,
			amount0Dec.Mul(upperReserve1Dec).Quo(lowerReserve0Dec)).TruncateInt()
	} else {
		trueAmount1 = amount1
	}

	valueMintedToken0 := CalcShares(trueAmount0, trueAmount1, centerTickPrice1To0)
	valueExistingToken0 := CalcShares(lowerReserve0, upperReserve1, centerTickPrice1To0)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).MulInt(totalShares).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
//                            TOKENIZER UTILS                                //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) MintShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	// mint share tokens
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// transfer them to addr
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sharesCoins); err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnShares(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Int, sharesId string) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesId, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	return nil
}
