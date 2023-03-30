package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/mev/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amt := sdk.Coins{sdk.Coin{
		Denom:  msg.TokenIn,
		Amount: msg.AmountIn,
	}}

	accAddressCreator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddressCreator, types.ModuleName, amt)

	if err != nil {
		return nil, err
	}

	return &types.MsgSendResponse{}, nil
}
