package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/mev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawFunds(goCtx context.Context, msg *types.MsgWithdrawFunds) (*types.MsgWithdrawFundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amt := sdk.Coins{sdk.Coin{
		Denom:  msg.TokenOut,
		Amount: msg.AmountOut,
	}}

	accAddressCreator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddressCreator, amt)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgWithdrawFundsResponse{}, nil
}
