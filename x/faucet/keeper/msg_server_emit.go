package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/faucet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "strconv"
)

func EmitTokens(ctx sdk.Context, k msgServer, recipient sdk.AccAddress, denom string) (error){
	emitAmount := sdk.NewInt(int64(types.DefaultAmount))
	coins := sdk.NewCoins(sdk.NewCoin(denom, emitAmount))

	err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil{
		return err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, recipient, coins)
	if err != nil{
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.FaucetEmitEventType,
			sdk.NewAttribute(types.FaucetEmitRecipient, recipient.String()),
			sdk.NewAttribute(types.FaucetEmitAmount, emitAmount.String()),
			sdk.NewAttribute(types.FaucetEmitDenom, denom),
		),
	)

	return nil

}

func (k msgServer) Emit(goCtx context.Context, msg *types.MsgEmit) (*types.MsgEmitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	recipient, err := sdk.AccAddressFromBech32(msg.Creator)

	err = EmitTokens(ctx, k, recipient, "stake")
	if err != nil{
		return nil, err
	}
	EmitTokens(ctx, k, recipient, "token")
	if err != nil{
		return nil, err
	}

	if err != nil{
		return nil, err
	}



	return &types.MsgEmitResponse{Success: true}, nil
}
