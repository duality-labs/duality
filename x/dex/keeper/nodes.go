package keeper

import (
	"encoding/binary"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetNodesCount get the total number of nodes
func (k Keeper) GetNodesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NodesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetNodesCount set the total number of nodes
func (k Keeper) SetNodesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NodesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendNodes appends a nodes in the store with a new id and update the count
func (k Keeper) AppendNodes(
	ctx sdk.Context,
	nodes types.Nodes,
) uint64 {
	// Create the nodes
	count := k.GetNodesCount(ctx)

	// Set the ID of the appended value
	nodes.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKey))
	appendedValue := k.cdc.MustMarshal(&nodes)
	store.Set(GetNodesIDBytes(nodes.Id), appendedValue)

	// Update nodes count
	k.SetNodesCount(ctx, count+1)

	return count
}

// SetNodes set a specific nodes in the store
func (k Keeper) SetNodes(ctx sdk.Context, nodes types.Nodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKey))
	b := k.cdc.MustMarshal(&nodes)
	store.Set(GetNodesIDBytes(nodes.Id), b)
}

// GetNodes returns a nodes from its id
func (k Keeper) GetNodes(ctx sdk.Context, id uint64) (val types.Nodes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKey))
	b := store.Get(GetNodesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNodes removes a nodes from the store
func (k Keeper) RemoveNodes(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKey))
	store.Delete(GetNodesIDBytes(id))
}

// GetAllNodes returns all nodes
func (k Keeper) GetAllNodes(ctx sdk.Context) (list []types.Nodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Nodes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNodesIDBytes returns the byte representation of the ID
func GetNodesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetNodesIDFromBytes returns ID in uint64 format from a byte array
func GetNodesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
