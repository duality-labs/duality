package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLimitOrderTrancheUser set a specific LimitOrderTrancheUser in the store from its index
func (k Keeper) SetLimitOrderTrancheUser(ctx sdk.Context, LimitOrderTrancheUser types.LimitOrderTrancheUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))
	b := k.cdc.MustMarshal(&LimitOrderTrancheUser)
	store.Set(types.LimitOrderTrancheUserKey(
		LimitOrderTrancheUser.PairId,
		LimitOrderTrancheUser.TickIndex,
		LimitOrderTrancheUser.Token,
		LimitOrderTrancheUser.Count,
		LimitOrderTrancheUser.Address,
	), b)
}

// GetLimitOrderTrancheUser returns a LimitOrderTrancheUser from its index
func (k Keeper) GetLimitOrderTrancheUser(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,
) (val types.LimitOrderTrancheUser, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))

	b := store.Get(types.LimitOrderTrancheUserKey(
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

// RemoveLimitOrderTrancheUser removes a LimitOrderTrancheUser from the store
func (k Keeper) RemoveLimitOrderTrancheUser(
	ctx sdk.Context,
	pairId string,
	tickIndex int64,
	token string,
	count uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))
	store.Delete(types.LimitOrderTrancheUserKey(
		pairId,
		tickIndex,
		token,
		count,
		address,
	))
}

// GetAllLimitOrderTrancheUser returns all LimitOrderTrancheUser
func (k Keeper) GetAllLimitOrderTrancheUser(ctx sdk.Context) (list []types.LimitOrderTrancheUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTrancheUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
