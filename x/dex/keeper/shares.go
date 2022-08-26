package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetShares set a specific shares in the store from its index
func (k Keeper) SetShares(ctx sdk.Context, token0 string, token1 string, shares types.Shares) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SharesPrefix(token0, token1))
	b := k.cdc.MustMarshal(&shares)
	store.Set(types.SharesKey(
		shares.Address,
		shares.Price,
		shares.Fee,
		shares.OrderType,
	), b)
}

// GetShares returns a shares from its index
func (k Keeper) GetShares(
	ctx sdk.Context,
	token0 string,
	token1 string,
	address string,
	price string,
	fee string,
	orderType string,

) (val types.Shares, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SharesPrefix(token0, token1))

	b := store.Get(types.SharesKey(
		address,
		price,
		fee,
		orderType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShares removes a shares from the store
func (k Keeper) RemoveShares(
	ctx sdk.Context,
	token0 string,
	token1 string,
	address string,
	price string,
	fee string,
	orderType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SharesPrefix(token0, token1))
	store.Delete(types.SharesKey(
		address,
		price,
		fee,
		orderType,
	))
}

// GetAllShares returns all shares
func (k Keeper) GetAllSharesByPair(ctx sdk.Context, token0 string, token1 string) (list []types.Shares) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SharesPrefix(token0, token1))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shares
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllShares returns all shares
func (k Keeper) GetAllShares(ctx sdk.Context) (list []types.Shares) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseSharesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shares
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
