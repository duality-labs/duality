package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// GetGaugeByID returns gauge from gauge ID.
func (k Keeper) GetPoolMetadataByID(
	ctx sdk.Context,
	poolMetadataID uint64,
) (*types.PoolMetadata, error) {
	store := ctx.KVStore(k.storeKey)
	poolKey := types.PoolMetadataKey(poolMetadataID)
	bz := store.Get(poolKey)
	if bz == nil {
		return nil, fmt.Errorf("poolMetadata with ID %d does not exist", poolMetadataID)
	}

	poolMetadata := &types.PoolMetadata{}
	k.cdc.MustUnmarshal(bz, poolMetadata)
	return poolMetadata, nil
}

// setGauge set the gauge inside store.
func (k Keeper) SetPoolMetadata(ctx sdk.Context, poolMetadata *types.PoolMetadata) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(poolMetadata)
	store.Set(types.PoolMetadataKey(poolMetadata.Id), bz)
}

// GetLastGaugeID returns the last used gauge ID.
func (k Keeper) GetLastPoolMetadataID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.LastPoolIDKey))
	if bz == nil {
		return 1
	}

	return sdk.BigEndianToUint64(bz)
}

// SetLastGaugeID sets the last used gauge ID to the provided ID.
func (k Keeper) SetLastPoolID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.LastPoolIDKey), sdk.Uint64ToBigEndian(id))
}

// getNextPoolIdAndIncrement returns the next poolMetadata Id, and increments the corresponding state entry.
func (k Keeper) getNextPoolIdAndIncrement(ctx sdk.Context) uint64 {
	lastPoolID := k.GetLastPoolMetadataID(ctx)
	nextPoolID := lastPoolID + 1
	k.SetLastPoolID(ctx, nextPoolID)
	return nextPoolID
}
