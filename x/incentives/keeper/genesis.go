package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/duality-labs/duality/x/incentives/types"
)

// InitGenesis initializes the incentives module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	if err := k.InitializeAllLocks(ctx, genState.Locks); err != nil {
		return
	}
	if err := k.InitializeAllGauges(ctx, genState.Gauges); err != nil {
		return
	}
	k.SetLastLockID(ctx, genState.LastLockId)
	k.SetLastGaugeID(ctx, genState.LastGaugeId)
}

// ExportGenesis returns the x/incentives module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	locks, err := k.GetLocks(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		Params:      k.GetParams(ctx),
		Gauges:      k.GetNotFinishedGauges(ctx),
		LastGaugeId: k.GetLastGaugeID(ctx),
		LastLockId:  k.GetLastLockID(ctx),
		Locks:       locks,
	}
}

// InitializeAllLocks takes a set of locks, and initializes state to be storing
// them all correctly.
func (k Keeper) InitializeAllLocks(ctx sdk.Context, locks types.Locks) error {
	for i, lock := range locks {
		if i%25000 == 0 {
			msg := fmt.Sprintf("Reset %d lock refs, cur lock ID %d", i, lock.ID)
			ctx.Logger().Info(msg)
		}
		err := k.setLock(ctx, lock)
		if err != nil {
			return err
		}

		err = k.addLockRefs(ctx, lock)
		if err != nil {
			return err
		}
	}

	return nil
}

// InitializeAllGauges takes a set of gauges, and initializes state to be storing
// them all correctly.
func (k Keeper) InitializeAllGauges(ctx sdk.Context, gauges types.Gauges) error {
	for _, gauge := range gauges {
		err := k.setGauge(ctx, gauge)
		if err != nil {
			return err
		}
		err = k.setGaugeRefs(ctx, gauge)
		if err != nil {
			return err
		}
	}
	return nil
}
