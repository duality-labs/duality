package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolFillMap set a specific limitOrderPoolFillMap in the store from its index
func (k Keeper) SetLimitOrderPoolFillMap(ctx sdk.Context, limitOrderPoolFillMap types.LimitOrderPoolFillMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillMapKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolFillMap)
	store.Set(types.LimitOrderPoolFillMapKey(
		limitOrderPoolFillMap.PairId,
		limitOrderPoolFillMap.TickIndex,
		limitOrderPoolFillMap.Token,
		limitOrderPoolFillMap.Count,
	), b)
}

// GetLimitOrderPoolFillMap returns a limitOrderPoolFillMap from its index
func (k Keeper) GetLimitOrderPoolFillMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) (val types.LimitOrderPoolFillMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillMapKeyPrefix))

	b := store.Get(types.LimitOrderPoolFillMapKey(
		pairId,
		tickIndex,
		token,
		count,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderPoolFillMap removes a limitOrderPoolFillMap from the store
func (k Keeper) RemoveLimitOrderPoolFillMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillMapKeyPrefix))
	store.Delete(types.LimitOrderPoolFillMapKey(
		pairId,
		tickIndex,
		token,
		count,
	))
}

// GetAllLimitOrderPoolFillMap returns all limitOrderPoolFillMap
func (k Keeper) GetAllLimitOrderPoolFillMap(ctx sdk.Context) (list []types.LimitOrderPoolFillMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolFillMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
