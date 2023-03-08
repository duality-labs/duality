package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/incentives/types"
)

func (k msgServer) CreateIncentivePlan(goCtx context.Context, msg *types.MsgCreateIncentivePlan) (*types.MsgCreateIncentivePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetIncentivePlan(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var incentivePlan = types.IncentivePlan{
		Creator:     msg.Creator,
		Index:       msg.Index,
		StartDate:   msg.StartDate,
		EndDate:     msg.EndDate,
		TradingPair: msg.TradingPair,
		TotalAmount: msg.TotalAmount,
		StartTick:   msg.StartTick,
		EndTick:     msg.EndTick,
	}

	k.SetIncentivePlan(
		ctx,
		incentivePlan,
	)
	return &types.MsgCreateIncentivePlanResponse{}, nil
}

func (k msgServer) UpdateIncentivePlan(goCtx context.Context, msg *types.MsgUpdateIncentivePlan) (*types.MsgUpdateIncentivePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetIncentivePlan(
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

	var incentivePlan = types.IncentivePlan{
		Creator:     msg.Creator,
		Index:       msg.Index,
		StartDate:   msg.StartDate,
		EndDate:     msg.EndDate,
		TradingPair: msg.TradingPair,
		TotalAmount: msg.TotalAmount,
		StartTick:   msg.StartTick,
		EndTick:     msg.EndTick,
	}

	k.SetIncentivePlan(ctx, incentivePlan)

	return &types.MsgUpdateIncentivePlanResponse{}, nil
}

func (k msgServer) DeleteIncentivePlan(goCtx context.Context, msg *types.MsgDeleteIncentivePlan) (*types.MsgDeleteIncentivePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetIncentivePlan(
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

	k.RemoveIncentivePlan(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteIncentivePlanResponse{}, nil
}
