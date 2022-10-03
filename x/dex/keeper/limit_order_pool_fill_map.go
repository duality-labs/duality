package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolFillMap set a specific limitOrderPoolFillMap in the store from its index
func (k Keeper) SetLimitOrderPoolFillMap(ctx sdk.Context, pairId string, tickIndex int64, token string, limitOrderPoolFillMap types.LimitOrderPoolFillMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderFillMapPrefix(pairId, tickIndex, token))
	b := k.cdc.MustMarshal(&limitOrderPoolFillMap)
	store.Set(types.LimitOrderPoolFillMapKey(
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderFillMapPrefix(pairId, tickIndex, token))

	b := store.Get(types.LimitOrderPoolFillMapKey(
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LimitOrderFillMapPrefix(pairId, tickIndex, token))
	store.Delete(types.LimitOrderPoolFillMapKey(
		count,
	))
}
