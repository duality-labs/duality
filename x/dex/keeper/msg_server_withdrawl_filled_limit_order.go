package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawlFilledLimitOrder(goCtx context.Context, msg *types.MsgWithdrawlFilledLimitOrder) (*types.MsgWithdrawlFilledLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawlFilledLimitOrderResponse{}, nil
}
