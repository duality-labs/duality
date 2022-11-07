package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolReserveObject set a specific limitOrderPoolReserveObject in the store from its index
func (k Keeper) SetLimitOrderPoolReserveObject(ctx sdk.Context, limitOrderPoolReserveObject types.LimitOrderPoolReserveObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolReserveObjectKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolReserveObject)
	store.Set(types.LimitOrderPoolReserveObjectKey(
		limitOrderPoolReserveObject.PairId,
		limitOrderPoolReserveObject.TickIndex,
		limitOrderPoolReserveObject.Token,
		limitOrderPoolReserveObject.Count,
	), b)
}

// GetLimitOrderPoolReserveObject returns a limitOrderPoolReserveObject from its index
func (k Keeper) GetLimitOrderPoolReserveObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) (val types.LimitOrderPoolReserveObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolReserveObjectKeyPrefix))

	b := store.Get(types.LimitOrderPoolReserveObjectKey(
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

// RemoveLimitOrderPoolReserveObject removes a limitOrderPoolReserveObject from the store
func (k Keeper) RemoveLimitOrderPoolReserveObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolReserveObjectKeyPrefix))
	store.Delete(types.LimitOrderPoolReserveObjectKey(
		pairId,
		tickIndex,
		token,
		count,
	))
}

// GetAllLimitOrderPoolReserveObject returns all limitOrderPoolReserveObject
func (k Keeper) GetAllLimitOrderPoolReserveObject(ctx sdk.Context) (list []types.LimitOrderPoolReserveObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolReserveObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolReserveObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
