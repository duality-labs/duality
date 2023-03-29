package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetLimitOrderTrancheUser set a specific LimitOrderTrancheUser in the store from its index
func (k Keeper) SetLimitOrderTrancheUser(ctx sdk.Context, LimitOrderTrancheUser types.LimitOrderTrancheUser) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))
	b := k.cdc.MustMarshal(&LimitOrderTrancheUser)
	store.Set(types.LimitOrderTrancheUserKey(
		LimitOrderTrancheUser.Address,
		LimitOrderTrancheUser.TrancheKey,
	), b)
}

// GetLimitOrderTrancheUser returns a LimitOrderTrancheUser from its index
func (k Keeper) GetLimitOrderTrancheUser(
	ctx sdk.Context,
	address string,
	trancheKey string,
) (val types.LimitOrderTrancheUser, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))

	b := store.Get(types.LimitOrderTrancheUserKey(
		address,
		trancheKey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLimitOrderTrancheUserByKey removes a LimitOrderTrancheUser from the store
func (k Keeper) RemoveLimitOrderTrancheUserByKey(
	ctx sdk.Context,
	pairId *types.PairId,
	tickIndex int64,
	token string,
	trancheKey string,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LimitOrderTrancheUserKeyPrefix))
	store.Delete(types.LimitOrderTrancheUserKey(
		address,
		trancheKey,
	))
}

func (k Keeper) RemoveLimitOrderTrancheUser(ctx sdk.Context, trancheUser types.LimitOrderTrancheUser) {
	k.RemoveLimitOrderTrancheUserByKey(
		ctx,
		trancheUser.PairId,
		trancheUser.TickIndex,
		trancheUser.Token,
		trancheUser.TrancheKey,
		trancheUser.Address,
	)
}

func (k Keeper) SaveTrancheUser(ctx sdk.Context, trancheUser types.LimitOrderTrancheUser) {
	if trancheUser.IsEmpty() {
		k.RemoveLimitOrderTrancheUser(ctx, trancheUser)
	} else {
		k.SetLimitOrderTrancheUser(ctx, trancheUser)
	}
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

func (k Keeper) GetAllLimitOrderTrancheUserForAddress(ctx sdk.Context, address sdk.AccAddress) (list []types.LimitOrderTrancheUser) {
	addressPrefix := types.LimitOrderTrancheUserAddressPrefix(address.String())
	store := prefix.NewStore(ctx.KVStore(k.storeKey), addressPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LimitOrderTrancheUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
