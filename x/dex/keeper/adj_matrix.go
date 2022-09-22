package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAdjMatrixCount get the total number of adjMatrix
func (k Keeper) GetAdjMatrixCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdjMatrixCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAdjMatrixCount set the total number of adjMatrix
func (k Keeper) SetAdjMatrixCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdjMatrixCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAdjMatrix appends a adjMatrix in the store with a new id and update the count
func (k Keeper) AppendAdjMatrix(
	ctx sdk.Context,
	adjMatrix types.AdjMatrix,
) uint64 {
	// Create the adjMatrix
	count := k.GetAdjMatrixCount(ctx)

	// Set the ID of the appended value
	adjMatrix.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjMatrixKey))
	appendedValue := k.cdc.MustMarshal(&adjMatrix)
	store.Set(GetAdjMatrixIDBytes(adjMatrix.Id), appendedValue)

	// Update adjMatrix count
	k.SetAdjMatrixCount(ctx, count+1)

	return count
}

// SetAdjMatrix set a specific adjMatrix in the store
func (k Keeper) SetAdjMatrix(ctx sdk.Context, adjMatrix types.AdjMatrix) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjMatrixKey))
	b := k.cdc.MustMarshal(&adjMatrix)
	store.Set(GetAdjMatrixIDBytes(adjMatrix.Id), b)
}

// GetAdjMatrix returns a adjMatrix from its id
func (k Keeper) GetAdjMatrix(ctx sdk.Context, id uint64) (val types.AdjMatrix, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjMatrixKey))
	b := store.Get(GetAdjMatrixIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAdjMatrix removes a adjMatrix from the store
func (k Keeper) RemoveAdjMatrix(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjMatrixKey))
	store.Delete(GetAdjMatrixIDBytes(id))
}

// GetAllAdjMatrix returns all adjMatrix
func (k Keeper) GetAllAdjMatrix(ctx sdk.Context) (list []types.AdjMatrix) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdjMatrixKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AdjMatrix
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAdjMatrixIDBytes returns the byte representation of the ID
func GetAdjMatrixIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAdjMatrixIDFromBytes returns ID in uint64 format from a byte array
func GetAdjMatrixIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
