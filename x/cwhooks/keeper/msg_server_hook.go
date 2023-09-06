package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/cwhooks/types"
)

func (k msgServer) CreateHook(goCtx context.Context, msg *types.MsgCreateHook) (*types.MsgCreateHookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hook := types.Hook{
		Creator:         msg.Creator,
		ContractAddress: msg.ContractID,
		Args:            msg.Args,
		Persistent:      msg.Persistent,
		TriggerKey:      msg.TriggerKey,
		TriggerValue:    msg.TriggerValue,
	}

	id := k.AppendHook(
		ctx,
		hook,
	)

	return &types.MsgCreateHookResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteHook(goCtx context.Context, msg *types.MsgDeleteHook) (*types.MsgDeleteHookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHookByID(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHookByID(ctx, msg.Id)

	return &types.MsgDeleteHookResponse{}, nil
}
