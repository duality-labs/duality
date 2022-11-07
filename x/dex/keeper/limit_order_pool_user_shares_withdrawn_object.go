package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUserSharesWithdrawnObject set a specific limitOrderPoolUserSharesWithdrawnObject in the store from its index
func (k Keeper) SetLimitOrderPoolUserSharesWithdrawnObject(ctx sdk.Context, limitOrderPoolUserSharesWithdrawnObject types.LimitOrderPoolUserSharesWithdrawnObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix))
	b := k.cdc.MustMarshal(&limitOrderPoolUserSharesWithdrawnObject)
	store.Set(types.LimitOrderPoolUserSharesWithdrawnObjectKey(
		limitOrderPoolUserSharesWithdrawnObject.PairId,
		limitOrderPoolUserSharesWithdrawnObject.TickIndex,
		limitOrderPoolUserSharesWithdrawnObject.Token,
		limitOrderPoolUserSharesWithdrawnObject.Count,
		limitOrderPoolUserSharesWithdrawnObject.Address,
	), b)
}

// GetLimitOrderPoolUserSharesWithdrawnObject returns a limitOrderPoolUserSharesWithdrawnObject from its index
func (k Keeper) GetLimitOrderPoolUserSharesWithdrawnObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUserSharesWithdrawnObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix))

	b := store.Get(types.LimitOrderPoolUserSharesWithdrawnObjectKey(
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

// RemoveLimitOrderPoolUserSharesWithdrawnObject removes a limitOrderPoolUserSharesWithdrawnObject from the store
func (k Keeper) RemoveLimitOrderPoolUserSharesWithdrawnObject(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix))
	store.Delete(types.LimitOrderPoolUserSharesWithdrawnObjectKey(
		pairId,
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderPoolUserSharesWithdrawnObject returns all limitOrderPoolUserSharesWithdrawnObject
func (k Keeper) GetAllLimitOrderPoolUserSharesWithdrawnObject(ctx sdk.Context) (list []types.LimitOrderPoolUserSharesWithdrawnObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolUserSharesWithdrawnObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
