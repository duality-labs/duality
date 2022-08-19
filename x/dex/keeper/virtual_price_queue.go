package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVirtualPriceQueue set a specific virtualPriceQueue in the store from its index
func (k Keeper) SetVirtualPriceQueue(ctx sdk.Context, virtualPriceQueue types.VirtualPriceQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceQueueKeyPrefix))
	b := k.cdc.MustMarshal(&virtualPriceQueue)
	store.Set(types.VirtualPriceQueueKey(
		virtualPriceQueue.VPrice,
		virtualPriceQueue.Direction,
		virtualPriceQueue.OrderType,
	), b)
}

// GetVirtualPriceQueue returns a virtualPriceQueue from its index
// TODO: Map each virtualpricequeue with respect to id instead of vPrice
func (k Keeper) GetVirtualPriceQueue(
	ctx sdk.Context,
	vPrice string,
	direction string,
	orderType string,

) (val types.VirtualPriceQueue, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceQueueKeyPrefix))

	b := store.Get(types.VirtualPriceQueueKey(
		vPrice,
		direction,
		orderType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualPriceQueue removes a virtualPriceQueue from the store
func (k Keeper) RemoveVirtualPriceQueue(
	ctx sdk.Context,
	vPrice string,
	direction string,
	orderType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceQueueKeyPrefix))
	store.Delete(types.VirtualPriceQueueKey(
		vPrice,
		direction,
		orderType,
	))
}

// GetAllVirtualPriceQueue returns all virtualPriceQueue
func (k Keeper) GetAllVirtualPriceQueue(ctx sdk.Context) (list []types.VirtualPriceQueue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceQueueKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualPriceQueue
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
