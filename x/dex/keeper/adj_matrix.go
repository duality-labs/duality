package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAdjanceyMatrixCount get the total number of adjanceyMatrix
func (k Keeper) GetAdjanceyMatrixCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdjanceyMatrixCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAdjanceyMatrixCount set the total number of adjanceyMatrix
func (k Keeper) SetAdjanceyMatrixCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdjanceyMatrixCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAdjanceyMatrix appends a adjanceyMatrix in the store with a new id and update the count
func (k Keeper) AppendAdjanceyMatrix(
	ctx sdk.Context,
	adjanceyMatrix types.AdjanceyMatrix,
) uint64 {
	// Create the adjanceyMatrix
	count := k.GetAdjanceyMatrixCount(ctx)

	// Set the ID of the appended value
	adjanceyMatrix.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjanceyMatrixKey))
	appendedValue := k.cdc.MustMarshal(&adjanceyMatrix)
	store.Set(GetAdjanceyMatrixIDBytes(adjanceyMatrix.Id), appendedValue)

	// Update adjanceyMatrix count
	k.SetAdjanceyMatrixCount(ctx, count+1)

	return count
}

// SetAdjanceyMatrix set a specific adjanceyMatrix in the store
func (k Keeper) SetAdjanceyMatrix(ctx sdk.Context, adjanceyMatrix types.AdjanceyMatrix) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjanceyMatrixKey))
	b := k.cdc.MustMarshal(&adjanceyMatrix)
	store.Set(GetAdjanceyMatrixIDBytes(adjanceyMatrix.Id), b)
}

// GetAdjanceyMatrix returns a adjanceyMatrix from its id
func (k Keeper) GetAdjanceyMatrix(ctx sdk.Context, id uint64) (val types.AdjanceyMatrix, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjanceyMatrixKey))
	b := store.Get(GetAdjanceyMatrixIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAdjanceyMatrix removes a adjanceyMatrix from the store
func (k Keeper) RemoveAdjanceyMatrix(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjanceyMatrixKey))
	store.Delete(GetAdjanceyMatrixIDBytes(id))
}

// GetAllAdjanceyMatrix returns all adjanceyMatrix
func (k Keeper) GetAllAdjanceyMatrix(ctx sdk.Context) (list []types.AdjanceyMatrix) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjanceyMatrixKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AdjanceyMatrix
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAdjanceyMatrixIDBytes returns the byte representation of the ID
func GetAdjanceyMatrixIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAdjanceyMatrixIDFromBytes returns ID in uint64 format from a byte array
func GetAdjanceyMatrixIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
