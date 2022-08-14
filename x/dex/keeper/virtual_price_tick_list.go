package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVirtualPriceTickList set a specific virtualPriceTickList in the store from its index
func (k Keeper) SetVirtualPriceTickList(ctx sdk.Context, virtualPriceTickList types.VirtualPriceTickList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickListKeyPrefix))
	b := k.cdc.MustMarshal(&virtualPriceTickList)
	store.Set(types.VirtualPriceTickListKey(
		virtualPriceTickList.VPrice,
		virtualPriceTickList.Direction,
		virtualPriceTickList.OrderType,
	), b)
}

// GetVirtualPriceTickList returns a virtualPriceTickList from its index
func (k Keeper) GetVirtualPriceTickList(
	ctx sdk.Context,
	vPrice string,
	direction string,
	orderType string,

) (val types.VirtualPriceTickList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickListKeyPrefix))

	b := store.Get(types.VirtualPriceTickListKey(
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

// RemoveVirtualPriceTickList removes a virtualPriceTickList from the store
func (k Keeper) RemoveVirtualPriceTickList(
	ctx sdk.Context,
	vPrice string,
	direction string,
	orderType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickListKeyPrefix))
	store.Delete(types.VirtualPriceTickListKey(
		vPrice,
		direction,
		orderType,
	))
}

// GetAllVirtualPriceTickList returns all virtualPriceTickList
func (k Keeper) GetAllVirtualPriceTickList(ctx sdk.Context) (list []types.VirtualPriceTickList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualPriceTickListKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualPriceTickList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
