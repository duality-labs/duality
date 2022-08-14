package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBitArrCount get the total number of bitArr
func (k Keeper) GetBitArrCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BitArrCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBitArrCount set the total number of bitArr
func (k Keeper) SetBitArrCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BitArrCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBitArr appends a bitArr in the store with a new id and update the count
func (k Keeper) AppendBitArr(
	ctx sdk.Context,
	bitArr types.BitArr,
) uint64 {
	// Create the bitArr
	count := k.GetBitArrCount(ctx)

	// Set the ID of the appended value
	bitArr.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitArrKey))
	appendedValue := k.cdc.MustMarshal(&bitArr)
	store.Set(GetBitArrIDBytes(bitArr.Id), appendedValue)

	// Update bitArr count
	k.SetBitArrCount(ctx, count+1)

	return count
}

// SetBitArr set a specific bitArr in the store
func (k Keeper) SetBitArr(ctx sdk.Context, bitArr types.BitArr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitArrKey))
	b := k.cdc.MustMarshal(&bitArr)
	store.Set(GetBitArrIDBytes(bitArr.Id), b)
}

// GetBitArr returns a bitArr from its id
func (k Keeper) GetBitArr(ctx sdk.Context, id uint64) (val types.BitArr, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitArrKey))
	b := store.Get(GetBitArrIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBitArr removes a bitArr from the store
func (k Keeper) RemoveBitArr(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitArrKey))
	store.Delete(GetBitArrIDBytes(id))
}

// GetAllBitArr returns all bitArr
func (k Keeper) GetAllBitArr(ctx sdk.Context) (list []types.BitArr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitArrKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BitArr
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBitArrIDBytes returns the byte representation of the ID
func GetBitArrIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBitArrIDFromBytes returns ID in uint64 format from a byte array
func GetBitArrIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
