package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/mevdummy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, amt)

	_ = ctx

	return &types.MsgSendResponse{}, nil
}
