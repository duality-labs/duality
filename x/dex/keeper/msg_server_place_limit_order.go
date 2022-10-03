package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceLimitOrder(goCtx context.Context, msg *types.MsgPlaceLimitOrder) (*types.MsgPlaceLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPlaceLimitOrderResponse{}, nil
}
