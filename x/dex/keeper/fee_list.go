package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetFeeListCount get the total number of feeList
func (k Keeper) GetFeeListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FeeListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFeeListCount set the total number of feeList
func (k Keeper) SetFeeListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FeeListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFeeList appends a feeList in the store with a new id and update the count
func (k Keeper) AppendFeeList(
	ctx sdk.Context,
	feeList types.FeeList,
) uint64 {
	// Create the feeList
	count := k.GetFeeListCount(ctx)

	// Set the ID of the appended value
	feeList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeListKey))
	appendedValue := k.cdc.MustMarshal(&feeList)
	store.Set(GetFeeListIDBytes(feeList.Id), appendedValue)

	// Update feeList count
	k.SetFeeListCount(ctx, count+1)

	return count
}

// SetFeeList set a specific feeList in the store
func (k Keeper) SetFeeList(ctx sdk.Context, feeList types.FeeList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeListKey))
	b := k.cdc.MustMarshal(&feeList)
	store.Set(GetFeeListIDBytes(feeList.Id), b)
}

// GetFeeList returns a feeList from its id
func (k Keeper) GetFeeList(ctx sdk.Context, id uint64) (val types.FeeList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeListKey))
	b := store.Get(GetFeeListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeList removes a feeList from the store
func (k Keeper) RemoveFeeList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeListKey))
	store.Delete(GetFeeListIDBytes(id))
}

// GetAllFeeList returns all feeList
func (k Keeper) GetAllFeeList(ctx sdk.Context) (list []types.FeeList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFeeListIDBytes returns the byte representation of the ID
func GetFeeListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFeeListIDFromBytes returns ID in uint64 format from a byte array
func GetFeeListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
