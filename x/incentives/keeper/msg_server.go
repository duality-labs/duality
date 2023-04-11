package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.MsgServer = msgServer{}

// msgServer provides a way to reference keeper pointer in the message server interface.
type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl returns an instance of MsgServer for the provided keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

// CreateGauge creates a gauge and sends coins to the gauge.
// Emits create gauge event and returns the create gauge response.
func (server msgServer) CreateGauge(goCtx context.Context, msg *types.MsgCreateGauge) (*types.MsgCreateGaugeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	gauge, err := server.keeper.CreateGauge(
		ctx,
		msg.IsPerpetual,
		owner,
		msg.Coins,
		msg.DistributeTo,
		msg.StartTime,
		msg.NumEpochsPaidOver,
		msg.PricingTick,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtCreateGauge,
			sdk.NewAttribute(types.AttributeGaugeID, strconv.FormatUint(gauge.Id, 10)),
		),
	})

	return &types.MsgCreateGaugeResponse{}, nil
}

// AddToGauge adds coins to gauge.
// Emits add to gauge event and returns the add to gauge response.
func (server msgServer) AddToGauge(goCtx context.Context, msg *types.MsgAddToGauge) (*types.MsgAddToGaugeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	err = server.keeper.AddToGaugeRewards(ctx, owner, msg.Rewards, msg.GaugeId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtAddToGauge,
			sdk.NewAttribute(types.AttributeGaugeID, strconv.FormatUint(msg.GaugeId, 10)),
		),
	})

	return &types.MsgAddToGaugeResponse{}, nil
}

// LockTokens locks tokens in either two ways.
// 1. Add to an existing lock if a lock with the same owner and same duration exists.
// 2. Create a new lock if not.
// A sanity check to ensure given tokens is a single token is done in ValidateBaic.
// That is, a lock with multiple tokens cannot be created.
func (server msgServer) LockTokens(goCtx context.Context, msg *types.MsgLockTokens) (*types.MsgLockTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	epochInfo := server.keeper.GetEpochInfo(ctx)
	epochDuration := epochInfo.Duration

	// check if there's an existing lock from the same owner with the same duration.
	// If so, simply add tokens to the existing lock.
	lockExists := server.keeper.HasFullLock(ctx, owner)
	if lockExists {
		lockID, err := server.keeper.AddToExistingLock(ctx, owner, msg.Coins)
		if err != nil {
			return nil, err
		}

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.TypeEvtAddTokensToLock,
				sdk.NewAttribute(types.AttributeLockID, strconv.FormatUint(lockID, 10)),
				sdk.NewAttribute(types.AttributeLockOwner, msg.Owner),
				sdk.NewAttribute(types.AttributeLockAmount, msg.Coins.String()),
			),
		})
		return &types.MsgLockTokensResponse{ID: lockID}, nil
	}

	// if the owner + duration combination is new, create a new lock.
	lock, err := server.keeper.CreateLock(ctx, owner, msg.Coins, epochDuration)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtLockTokens,
			sdk.NewAttribute(types.AttributeLockID, strconv.FormatUint(lock.ID, 10)),
			sdk.NewAttribute(types.AttributeLockOwner, lock.Owner),
			sdk.NewAttribute(types.AttributeLockAmount, lock.Coins.String()),
			sdk.NewAttribute(types.AttributeLockDuration, lock.Duration.String()),
			sdk.NewAttribute(types.AttributeLockUnlockTime, lock.EndTime.String()),
		),
	})

	return &types.MsgLockTokensResponse{ID: lock.ID}, nil
}

// BeginUnlocking begins unlocking of the specified lock.
// The lock would enter the unlocking queue, with the endtime of the lock set as block time + duration.
func (server msgServer) BeginUnlocking(goCtx context.Context, msg *types.MsgBeginUnlocking) (*types.MsgBeginUnlockingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lock, err := server.keeper.GetLockByID(ctx, msg.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	if msg.Owner != lock.Owner {
		return nil, sdkerrors.Wrap(types.ErrNotLockOwner, fmt.Sprintf("msg sender (%s) and lock owner (%s) does not match", msg.Owner, lock.Owner))
	}

	unlockingLock, err := server.keeper.BeginUnlock(ctx, lock.ID, msg.Coins)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// N.B. begin unlock event is emitted downstream in the keeper method.

	return &types.MsgBeginUnlockingResponse{Success: true, UnlockingLockID: unlockingLock}, nil
}

// BeginUnlockingAll begins unlocking for all the locks that the account has by iterating all the not-unlocking locks the account holds.
func (server msgServer) BeginUnlockingAll(goCtx context.Context, msg *types.MsgBeginUnlockingAll) (*types.MsgBeginUnlockingAllResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	unlocks := server.keeper.getLocksFromIterator(
		ctx,
		server.keeper.iterator(ctx, types.GetKeyLockIndexByAccount(false, owner)))

	for _, lock := range unlocks {
		_, err := server.keeper.BeginUnlock(ctx, lock.ID, nil)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the events for this message
	unlockedCoins := unlocks.GetCoins()
	events := sdk.Events{
		sdk.NewEvent(
			types.TypeEvtBeginUnlockAll,
			sdk.NewAttribute(types.AttributeLockOwner, msg.Owner),
			sdk.NewAttribute(types.AttributeUnlockedCoins, unlockedCoins.String()),
		),
	}
	for _, lock := range unlocks {
		lock := lock
		events = events.AppendEvent(createBeginUnlockEvent(lock))
	}
	ctx.EventManager().EmitEvents(events)

	return &types.MsgBeginUnlockingAllResponse{}, nil
}

func createBeginUnlockEvent(lock *types.Lock) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtBeginUnlock,
		sdk.NewAttribute(types.AttributeLockID, strconv.FormatUint(lock.ID, 10)),
		sdk.NewAttribute(types.AttributeLockOwner, lock.Owner),
		sdk.NewAttribute(types.AttributeLockDuration, lock.Duration.String()),
		sdk.NewAttribute(types.AttributeLockUnlockTime, lock.EndTime.String()),
	)
}
