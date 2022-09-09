package keeper

import (
	"context"

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

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, createrAddr sdk.AccAddress, amount0 sdk.Dec, amount1 sdk.Dec) error {

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

	if amount0.GT(sdk.ZeroDec()) && (msg.PriceIndex-fee) >= pair.TokenPair.CurrentTick1To0 {

	}

	if amount1.GT(sdk.ZeroDec()) && (msg.PriceIndex-fee) >= pair.TokenPair.CurrentTick0To1 {

	}

	_ = goCtx
	return nil
}

func (k Keeper) MultiDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}
