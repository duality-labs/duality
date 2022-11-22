package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPairMap set a specific pairMap in the store from its index
func (k Keeper) SetPairMap(ctx sdk.Context, pairMap types.PairMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairMapKeyPrefix))
	b := k.cdc.MustMarshal(&pairMap)
	k.Logger(ctx).Error("SetPairMap", "pairMap", pairMap)
	store.Set(types.PairMapKey(
		pairMap.PairId,
	), b)
}

// GetPairMap returns a pairMap from its index
func (k Keeper) GetPairMap(
	ctx sdk.Context,
	pairId string,
) (val types.PairMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairMapKeyPrefix))

	b := store.Get(types.PairMapKey(
		pairId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairMap removes a pairMap from the store
func (k Keeper) RemovePairMap(
	ctx sdk.Context,
	pairId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairMapKeyPrefix))
	store.Delete(types.PairMapKey(
		pairId,
	))
}

// GetAllPairMap returns all pairMap
func (k Keeper) GetAllPairMap(ctx sdk.Context) (list []types.PairMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PairMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
