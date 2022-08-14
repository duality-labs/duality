package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddLiquidityVerification(goCtx context.Context, msg *types.MsgAddLiquidity) (string, string, string, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", "", sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	_ = token0
	_ = token1
	_ = ctx

	return "", "", "", sdk.ZeroDec(), nil
}
