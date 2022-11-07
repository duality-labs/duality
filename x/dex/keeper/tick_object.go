package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTickObject set a specific tickObject in the store from its index
func (k Keeper) SetTickObject(ctx sdk.Context, pairId string, tickObject types.TickObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickObjectKeyPrefix))
	b := k.cdc.MustMarshal(&tickObject)
	store.Set(types.TickObjectKey(
		pairId,
		tickObject.TickIndex,
	), b)
}

// GetTickObject returns a tickObject from its index
func (k Keeper) GetTickObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) (val types.TickObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickObjectKeyPrefix))

	b := store.Get(types.TickObjectKey(
		pairId,
		tickIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTickObject removes a tickObject from the store
func (k Keeper) RemoveTickObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickObjectKeyPrefix))
	store.Delete(types.TickObjectKey(
		pairId,
		tickIndex,
	))
}

// GetAllTickObject returns all tickObject
func (k Keeper) GetAllTickObject(ctx sdk.Context) (list []types.TickObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseTickObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllTickObject returns all tickObject
func (k Keeper) GetAllTickObjectByPair(ctx sdk.Context, pairId string) (list []types.TickObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TickPrefix(pairId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TickObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
