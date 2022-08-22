package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNodes set a specific nodes in the store from its index
func (k Keeper) SetNodes(ctx sdk.Context, nodes types.Nodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKeyPrefix))
	b := k.cdc.MustMarshal(&nodes)
	store.Set(types.NodesKey(
		nodes.Node,
		nodes.OutgoingEdges,
	), b)
}

// GetNodes returns a nodes from its index
func (k Keeper) GetNodes(
	ctx sdk.Context,
	node string,
	outgoingEdges string,

) (val types.Nodes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKeyPrefix))

	b := store.Get(types.NodesKey(
		node,
		outgoingEdges,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNodes removes a nodes from the store
func (k Keeper) RemoveNodes(
	ctx sdk.Context,
	node string,
	outgoingEdges string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKeyPrefix))
	store.Delete(types.NodesKey(
		node,
		outgoingEdges,
	))
}

// GetAllNodes returns all nodes
func (k Keeper) GetAllNodes(ctx sdk.Context) (list []types.Nodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Nodes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
