package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) depositPairHelper(goCtx context.Context, token0 string, token1 string) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0Index, token0Found := k.GetTokenMap(ctx, token0)
	tokenLength := k.GetTokensCount(ctx)
	addEdge := false

	if !token0Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token0, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token0Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token0})
		addEdge = true
	}

	token1Index, token1Found := k.GetTokenMap(ctx, token1)

	if !token1Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token1, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token1Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token1})
		addEdge = true
	}

	if addEdge {

	}
}

func (k Keeper) addEdge(goCtx context.Context, token0Index int64, token1Index int64) {
	x := 4
	_ = x
}

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, createrAddr sdk.AccAddress, amount0 sdk.Dec, amount1 sdk.Dec) error {

	// DepositPairHelper(token0 token1)

	//
	_ = goCtx
	return nil
}

func (k Keeper) MultiDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}
