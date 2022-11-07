package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolTotalSharesObject set a specific limitOrderPoolTotalSharesObject in the store from its index
func (k Keeper) SetLimitOrderPoolTotalSharesObject(ctx sdk.Context, limitOrderPoolTotalSharesObject types.LimitOrderPoolTotalSharesObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolTotalSharesObjectKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolTotalSharesObject)
	store.Set(types.LimitOrderPoolTotalSharesObjectKey(
		limitOrderPoolTotalSharesObject.PairId,
		limitOrderPoolTotalSharesObject.TickIndex,
		limitOrderPoolTotalSharesObject.Token,
		limitOrderPoolTotalSharesObject.Count,
	), b)
}

// GetLimitOrderPoolTotalSharesObject returns a limitOrderPoolTotalSharesObject from its index
func (k Keeper) GetLimitOrderPoolTotalSharesObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) (val types.LimitOrderPoolTotalSharesObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolTotalSharesObjectKeyPrefix))

	b := store.Get(types.LimitOrderPoolTotalSharesObjectKey(
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

// RemoveLimitOrderPoolTotalSharesObject removes a limitOrderPoolTotalSharesObject from the store
func (k Keeper) RemoveLimitOrderPoolTotalSharesObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolTotalSharesObjectKeyPrefix))
	store.Delete(types.LimitOrderPoolTotalSharesObjectKey(
		pairId,
		tickIndex,
		token,
		count,
	))
}

// GetAllLimitOrderPoolTotalSharesObject returns all limitOrderPoolUserShareObject
func (k Keeper) GetAllLimitOrderPoolTotalSharesObject(ctx sdk.Context) (list []types.LimitOrderPoolTotalSharesObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolTotalSharesObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolTotalSharesObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
