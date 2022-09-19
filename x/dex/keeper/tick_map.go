package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTickMap set a specific tickMap in the store from its index
func (k Keeper) SetTickMap(ctx sdk.Context, pairId string, tickMap types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	b := k.cdc.MustMarshal(&tickMap)
	store.Set(types.TickMapKey(
		tickMap.TickIndex,
	), b)
}

// GetTickMap returns a tickMap from its index
func (k Keeper) GetTickMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) (val types.TickMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))

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
	pairId string,
	tickIndex int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	store.Delete(types.TickMapKey(
		tickIndex,
	))
}

// GetAllTickMap returns all tickMap
func (k Keeper) GetAllTickMap(ctx sdk.Context) (list []types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllTickMap returns all tickMap
func (k Keeper) GetAllTickMapByPair(ctx sdk.Context, pairId string) (list []types.TickMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
