package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetFeeTierCount get the total number of FeeTier
func (k Keeper) GetFeeTierCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FeeTierCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFeeTierCount set the total number of FeeTier
func (k Keeper) SetFeeTierCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FeeTierCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFeeTier appends a FeeTier in the store with a new id and update the count
func (k Keeper) AppendFeeTier(
	ctx sdk.Context,
	FeeTier types.FeeTier,
) uint64 {
	// Create the FeeTier
	count := k.GetFeeTierCount(ctx)

	// Set the ID of the appended value
	FeeTier.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeTierKey))
	appendedValue := k.cdc.MustMarshal(&FeeTier)
	store.Set(GetFeeTierIDBytes(FeeTier.Id), appendedValue)

	// Update FeeTier count
	k.SetFeeTierCount(ctx, count+1)

	return count
}

// SetFeeTier set a specific FeeTier in the store
func (k Keeper) SetFeeTier(ctx sdk.Context, FeeTier types.FeeTier) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeTierKey))
	b := k.cdc.MustMarshal(&FeeTier)
	store.Set(GetFeeTierIDBytes(FeeTier.Id), b)
}

// GetFeeTier returns a FeeTier from its id
func (k Keeper) GetFeeTier(ctx sdk.Context, id uint64) (val types.FeeTier, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeTierKey))
	b := store.Get(GetFeeTierIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeTier removes a FeeTier from the store
func (k Keeper) RemoveFeeTier(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeTierKey))
	store.Delete(GetFeeTierIDBytes(id))
}

// GetAllFeeTier returns all FeeTier
func (k Keeper) GetAllFeeTier(ctx sdk.Context) (list []types.FeeTier) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeTierKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeTier
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFeeTierIDBytes returns the byte representation of the ID
func GetFeeTierIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFeeTierIDFromBytes returns ID in uint64 format from a byte array
func GetFeeTierIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
