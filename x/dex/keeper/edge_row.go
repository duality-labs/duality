package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEdgeRowCount get the total number of edgeRow
func (k Keeper) GetEdgeRowCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EdgeRowCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEdgeRowCount set the total number of edgeRow
func (k Keeper) SetEdgeRowCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EdgeRowCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEdgeRow appends a edgeRow in the store with a new id and update the count
func (k Keeper) AppendEdgeRow(
	ctx sdk.Context,
	edgeRow types.EdgeRow,
) uint64 {
	// Create the edgeRow
	count := k.GetEdgeRowCount(ctx)

	// Set the ID of the appended value
	edgeRow.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EdgeRowKey))
	appendedValue := k.cdc.MustMarshal(&edgeRow)
	store.Set(GetEdgeRowIDBytes(edgeRow.Id), appendedValue)

	// Update edgeRow count
	k.SetEdgeRowCount(ctx, count+1)

	return count
}

// SetEdgeRow set a specific edgeRow in the store
func (k Keeper) SetEdgeRow(ctx sdk.Context, edgeRow types.EdgeRow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EdgeRowKey))
	b := k.cdc.MustMarshal(&edgeRow)
	store.Set(GetEdgeRowIDBytes(edgeRow.Id), b)
}

// GetEdgeRow returns a edgeRow from its id
func (k Keeper) GetEdgeRow(ctx sdk.Context, id uint64) (val types.EdgeRow, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EdgeRowKey))
	b := store.Get(GetEdgeRowIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEdgeRow removes a edgeRow from the store
func (k Keeper) RemoveEdgeRow(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EdgeRowKey))
	store.Delete(GetEdgeRowIDBytes(id))
}

// GetAllEdgeRow returns all edgeRow
func (k Keeper) GetAllEdgeRow(ctx sdk.Context) (list []types.EdgeRow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EdgeRowKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EdgeRow
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEdgeRowIDBytes returns the byte representation of the ID
func GetEdgeRowIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEdgeRowIDFromBytes returns ID in uint64 format from a byte array
func GetEdgeRowIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
