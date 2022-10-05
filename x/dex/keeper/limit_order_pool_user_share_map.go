package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUserShareMap set a specific limitOrderPoolUserShareMap in the store from its index
func (k Keeper) SetLimitOrderPoolUserShareMap(ctx sdk.Context, pairId string, limitOrderPoolUserShareMap types.LimitOrderPoolUserShareMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesMapPrefix(pairId))
	b := k.cdc.MustMarshal(&limitOrderPoolUserShareMap)
	store.Set(types.LimitOrderPoolUserShareMapKey(
		limitOrderPoolUserShareMap.TickIndex,
		limitOrderPoolUserShareMap.Token,
		limitOrderPoolUserShareMap.Count,
		limitOrderPoolUserShareMap.Address,
	), b)
}

// GetLimitOrderPoolUserShareMap returns a limitOrderPoolUserShareMap from its index
func (k Keeper) GetLimitOrderPoolUserShareMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUserShareMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesMapPrefix(pairId))

	b := store.Get(types.LimitOrderPoolUserShareMapKey(
		tickIndex,
		token,
		count,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderPoolUserShareMap removes a limitOrderPoolUserShareMap from the store
func (k Keeper) RemoveLimitOrderPoolUserShareMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderUserSharesMapPrefix(pairId))
	store.Delete(types.LimitOrderPoolUserShareMapKey(
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderPoolUserShareMap returns all limitOrderPoolUserShareMap
func (k Keeper) GetAllLimitOrderPoolUserShareMap(ctx sdk.Context) (list []types.LimitOrderPoolUserShareMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserShareMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolUserShareMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
