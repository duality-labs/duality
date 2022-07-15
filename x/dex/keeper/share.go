package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetShare set a specific share in the store from its index
func (k Keeper) SetShare(ctx sdk.Context, share types.Share) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareKeyPrefix))
	b := k.cdc.MustMarshal(&share)
	store.Set(types.ShareKey(
		share.Owner,
		share.Token0,
		share.Token1,
		share.Price,
		share.Fee,
	), b)
}

// GetShare returns a share from its index
func (k Keeper) GetShare(
	ctx sdk.Context,
	owner string,
	token0 string,
	token1 string,
	price string,
	fee string,

) (val types.Share, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareKeyPrefix))

	b := store.Get(types.ShareKey(
		owner,
		token0,
		token1,
		price,
		fee,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShare removes a share from the store
func (k Keeper) RemoveShare(
	ctx sdk.Context,
	owner string,
	token0 string,
	token1 string,
	price string,
	fee string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareKeyPrefix))
	store.Delete(types.ShareKey(
		owner,
		token0,
		token1,
		price,
		fee,
	))
}

// GetAllShare returns all share
func (k Keeper) GetAllShare(ctx sdk.Context) (list []types.Share) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Share
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
