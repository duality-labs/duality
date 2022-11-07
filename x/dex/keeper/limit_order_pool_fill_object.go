package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolFillObject set a specific limitOrderPoolFillObject in the store from its index
func (k Keeper) SetLimitOrderPoolFillObject(ctx sdk.Context, limitOrderPoolFillObject types.LimitOrderPoolFillObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillObjectKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolFillObject)
	store.Set(types.LimitOrderPoolFillObjectKey(
		limitOrderPoolFillObject.PairId,
		limitOrderPoolFillObject.TickIndex,
		limitOrderPoolFillObject.Token,
		limitOrderPoolFillObject.Count,
	), b)
}

// GetLimitOrderPoolFillObject returns a limitOrderPoolFillObject from its index
func (k Keeper) GetLimitOrderPoolFillObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) (val types.LimitOrderPoolFillObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillObjectKeyPrefix))

	b := store.Get(types.LimitOrderPoolFillObjectKey(
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

// RemoveLimitOrderPoolFillObject removes a limitOrderPoolFillObject from the store
func (k Keeper) RemoveLimitOrderPoolFillObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillObjectKeyPrefix))
	store.Delete(types.LimitOrderPoolFillObjectKey(
		pairId,
		tickIndex,
		token,
		count,
	))
}

// GetAllLimitOrderPoolFillObject returns all limitOrderPoolFillObject
func (k Keeper) GetAllLimitOrderPoolFillObject(ctx sdk.Context) (list []types.LimitOrderPoolFillObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolFillObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolFillObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
