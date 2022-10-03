package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolReserveMap set a specific limitOrderPoolReserveMap in the store from its index
func (k Keeper) SetLimitOrderPoolReserveMap(ctx sdk.Context, pairId string, tickIndex int64, token string, limitOrderPoolReserveMap types.LimitOrderPoolReserveMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderReseveMapPrefix(pairId, tickIndex, token))
	b := k.cdc.MustMarshal(&limitOrderPoolReserveMap)
	store.Set(types.LimitOrderPoolReserveMapKey(
		limitOrderPoolReserveMap.Count,
	), b)
}

// GetLimitOrderPoolReserveMap returns a limitOrderPoolReserveMap from its index
func (k Keeper) GetLimitOrderPoolReserveMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) (val types.LimitOrderPoolReserveMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderReseveMapPrefix(pairId, tickIndex, token))

	b := store.Get(types.LimitOrderPoolReserveMapKey(
		count,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderPoolReserveMap removes a limitOrderPoolReserveMap from the store
func (k Keeper) RemoveLimitOrderPoolReserveMap(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderReseveMapPrefix(pairId, tickIndex, token))
	store.Delete(types.LimitOrderPoolReserveMapKey(
		count,
	))
}

// GetAllLimitOrderPoolReserveMap returns all limitOrderPoolReserveMap
func (k Keeper) GetAllLimitOrderPoolReserveMap(ctx sdk.Context) (list []types.LimitOrderPoolReserveMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolReserveMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolReserveMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
