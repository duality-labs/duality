package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) addEdges(goCtx context.Context, token0Index int64, token1Index int64) {
	x := 4
	_ = x
}

func (k Keeper) depositPairHelper(goCtx context.Context, token0 string, token1 string, price_index int64, feeIndex int64) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0Index, token0Found := k.GetTokenMap(ctx, token0)
	tokenLength := k.GetTokensCount(ctx)

	if !token0Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token0, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token0Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token0})
	}

	token1Index, token1Found := k.GetTokenMap(ctx, token1)

	if !token1Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token1, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token1Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token1})
	}

	pairId := k.createPairId(token0, token1)
	_, PairFound := k.GetPairMap(ctx, pairId)

	if !PairFound {

		feeValue, _ := k.GetFeeList(ctx, uint64(feeIndex))

		addEdges(goCtx, token0Index.Index, token1Index.Index)

		k.SetPairMap(ctx, types.PairMap{
			PairId: pairId,
			TokenPair: &types.TokenPairType{
				CurrentTick0To1: price_index + feeValue.Fee,
				CurrentTick1To0: price_index - feeValue.Fee,
			},
		})
	}

}

func calc_price(price_Index int64) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(price_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		return price, nil
	}

}

func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, callerAddr sdk.AccAddress, amount0 sdk.Dec, amount1 sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	feeValue, _ := k.GetFeeList(ctx, uint64(msg.FeeIndex))
	fee := feeValue.Fee
	k.depositPairHelper(goCtx, token0, token1, msg.PriceIndex, fee)

	pairId := k.createPairId(token0, token1)

	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	subFeeTick, subFeeTickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex-fee)
	addFeeTick, addFeeTickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex+fee)

	trueAmount0 := amount0
	trueAmount1 := amount1
	var sharesMinted sdk.Dec
	var oldReserve0 sdk.Dec
	var oldReserve1 sdk.Dec

	price, err := calc_price(msg.PriceIndex)

	if err != nil {
		return err
	}

	if !subFeeTickFound || !addFeeTickFound || addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Equal(sdk.ZeroDec()) {
		sharesMinted = trueAmount0.Add(amount1.Mul(price))

		if !subFeeTickFound || !addFeeTickFound {

			// We must init both sides even if we only use one from each.
			addFeeTick.TickData.Reserve0AndShares = make([]*types.Reserve0AndSharesType, k.GetFeeListCount(ctx))
			addFeeTick.TickData.Reserve1 = make([]sdk.Dec, k.GetFeeListCount(ctx))
			subFeeTick.TickData.Reserve0AndShares = make([]*types.Reserve0AndSharesType, k.GetFeeListCount(ctx))
			subFeeTick.TickData.Reserve1 = make([]sdk.Dec, k.GetFeeListCount(ctx))
		}

		oldReserve0 = sdk.ZeroDec()
		oldReserve1 = sdk.ZeroDec()

		addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0 = trueAmount0
		addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares = sharesMinted

		subFeeTick.TickData.Reserve1[msg.FeeIndex] = trueAmount1

	} else {
		if addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.GT(sdk.ZeroDec()) && (msg.PriceIndex-fee) >= pair.TokenPair.CurrentTick1To0 {
			trueAmount1 = k.Min(amount1, subFeeTick.TickData.Reserve1[msg.FeeIndex].Mul(amount0).Quo(addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0))
		}

		if subFeeTick.TickData.Reserve1[msg.FeeIndex].GT(sdk.ZeroDec()) && (msg.PriceIndex-fee) >= pair.TokenPair.CurrentTick0To1 {
			trueAmount0 = k.Min(amount0, addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.Mul(amount1).Quo(subFeeTick.TickData.Reserve1[msg.FeeIndex]))
		}

		sharesMinted = (trueAmount0.Quo(addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0).Mul(addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares)).Add(trueAmount1.Quo(subFeeTick.TickData.Reserve1[msg.FeeIndex]).Mul(addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares))

		oldReserve0 = addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0
		oldReserve1 = subFeeTick.TickData.Reserve1[msg.FeeIndex]
		addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0 = addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.Add(trueAmount0)
		addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares = addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Add(sharesMinted)

		subFeeTick.TickData.Reserve1[msg.FeeIndex] = subFeeTick.TickData.Reserve1[msg.FeeIndex].Add(trueAmount1)

	}

	shares, sharesFound := k.GetShares(ctx, msg.Creator, pairId, msg.PriceIndex, msg.FeeIndex)

	if !sharesFound {
		shares.SharesOwned = sharesMinted
	} else {
		shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
	}

	if amount0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Cannnot send zero amount")
	}

	if amount1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Cannnot send zero amount")
	}

	k.SetTickMap(ctx, pairId, addFeeTick)
	k.SetTickMap(ctx, pairId, subFeeTick)
	k.SetShares(ctx, shares)

	ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator,
		token0, token1, string(msg.PriceIndex), string(msg.FeeIndex),
		oldReserve0.String(), oldReserve1.String(), subFeeTick.TickData.Reserve1[msg.FeeIndex].String(), addFeeTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.String(),
		sharesMinted.String()),
	)

	_ = goCtx
	return nil
}

func (k Keeper) MultiDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}
