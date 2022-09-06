package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTickMap set a specific tickMap in the store from its index
func (k Keeper) SetTickMap(ctx sdk.Context, tickMap types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickMapKeyPrefix))
	b := k.cdc.MustMarshal(&tickMap)
	store.Set(types.TickMapKey(
		tickMap.TickIndex,
	), b)
}

// GetTickMap returns a tickMap from its index
func (k Keeper) GetTickMap(
	ctx sdk.Context,
	tickIndex string,

) (val types.TickMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickMapKeyPrefix))

	b := store.Get(types.TickMapKey(
		tickIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTickMap removes a tickMap from the store
func (k Keeper) RemoveTickMap(
	ctx sdk.Context,
	tickIndex string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickMapKeyPrefix))
	store.Delete(types.TickMapKey(
		tickIndex,
	))
}

// GetAllTickMap returns all tickMap
func (k Keeper) GetAllTickMap(ctx sdk.Context) (list []types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TickMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
