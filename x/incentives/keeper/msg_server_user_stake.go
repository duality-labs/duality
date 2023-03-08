package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/incentives/types"
)

func (k msgServer) CreateUserStake(goCtx context.Context, msg *types.MsgCreateUserStake) (*types.MsgCreateUserStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUserStake(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var userStake = types.UserStake{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Amount:    msg.Amount,
		StartDate: msg.StartDate,
		EndDate:   msg.EndDate,
	}

	k.SetUserStake(
		ctx,
		userStake,
	)
	return &types.MsgCreateUserStakeResponse{}, nil
}

func (k msgServer) UpdateUserStake(goCtx context.Context, msg *types.MsgUpdateUserStake) (*types.MsgUpdateUserStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserStake(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var userStake = types.UserStake{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Amount:    msg.Amount,
		StartDate: msg.StartDate,
		EndDate:   msg.EndDate,
	}

	k.SetUserStake(ctx, userStake)

	return &types.MsgUpdateUserStakeResponse{}, nil
}

func (k msgServer) DeleteUserStake(goCtx context.Context, msg *types.MsgDeleteUserStake) (*types.MsgDeleteUserStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserStake(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUserStake(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteUserStakeResponse{}, nil
}
