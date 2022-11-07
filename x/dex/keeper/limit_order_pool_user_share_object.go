package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUserShareObject set a specific limitOrderPoolUserShareObject in the store from its index
func (k Keeper) SetLimitOrderPoolUserShareObject(ctx sdk.Context, limitOrderPoolUserShareObject types.LimitOrderPoolUserShareObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserShareObjectKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolUserShareObject)
	store.Set(types.LimitOrderPoolUserShareObjectKey(
		limitOrderPoolUserShareObject.PairId,
		limitOrderPoolUserShareObject.TickIndex,
		limitOrderPoolUserShareObject.Token,
		limitOrderPoolUserShareObject.Count,
		limitOrderPoolUserShareObject.Address,
	), b)
}

// GetLimitOrderPoolUserShareObject returns a limitOrderPoolUserShareObject from its index
func (k Keeper) GetLimitOrderPoolUserShareObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUserShareObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserShareObjectKeyPrefix))

	b := store.Get(types.LimitOrderPoolUserShareObjectKey(
		pairId,
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

// RemoveLimitOrderPoolUserShareObject removes a limitOrderPoolUserShareObject from the store
func (k Keeper) RemoveLimitOrderPoolUserShareObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserShareObjectKeyPrefix))
	store.Delete(types.LimitOrderPoolUserShareObjectKey(
		pairId,
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderPoolUserShareObject returns all limitOrderPoolUserShareObject
func (k Keeper) GetAllLimitOrderPoolUserShareObject(ctx sdk.Context) (list []types.LimitOrderPoolUserShareObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserShareObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolUserShareObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
