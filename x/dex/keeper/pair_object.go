package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPairObject set a specific pairObject in the store from its index
func (k Keeper) SetPairObject(ctx sdk.Context, pairObject types.PairObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairObjectKeyPrefix))
	b := k.cdc.MustMarshal(&pairObject)
	store.Set(types.PairObjectKey(
		pairObject.PairId,
	), b)
}

// GetPairObject returns a pairObject from its index
func (k Keeper) GetPairObject(
	ctx sdk.Context,
	pairId string,

) (val types.PairObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairObjectKeyPrefix))

	b := store.Get(types.PairObjectKey(
		pairId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairObject removes a pairObject from the store
func (k Keeper) RemovePairObject(
	ctx sdk.Context,
	pairId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairObjectKeyPrefix))
	store.Delete(types.PairObjectKey(
		pairId,
	))
}

// GetAllPairObject returns all pairObject
func (k Keeper) GetAllPairObject(ctx sdk.Context) (list []types.PairObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PairObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
