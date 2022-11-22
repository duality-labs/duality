package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderPoolUser set a specific LimitOrderPoolUser in the store from its index
func (k Keeper) SetLimitOrderPoolUser(ctx sdk.Context, LimitOrderPoolUser types.LimitOrderPoolUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserKeyPrefix))
	b := k.cdc.MustMarshal(&LimitOrderPoolUser)
	store.Set(types.LimitOrderPoolUserKey(
		LimitOrderPoolUser.PairId,
		LimitOrderPoolUser.TickIndex,
		LimitOrderPoolUser.Token,
		LimitOrderPoolUser.Count,
		LimitOrderPoolUser.Address,
	), b)
}

// GetLimitOrderPoolUser returns a LimitOrderPoolUser from its index
func (k Keeper) GetLimitOrderPoolUser(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) (val types.LimitOrderPoolUser, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserKeyPrefix))

	b := store.Get(types.LimitOrderPoolUserKey(
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

// RemoveLimitOrderPoolUser removes a LimitOrderPoolUser from the store
func (k Keeper) RemoveLimitOrderPoolUser(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserKeyPrefix))
	store.Delete(types.LimitOrderPoolUserKey(
		pairId,
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderPoolUser returns all LimitOrderPoolUser
func (k Keeper) GetAllLimitOrderPoolUser(ctx sdk.Context) (list []types.LimitOrderPoolUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderPoolUserKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderPoolUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
