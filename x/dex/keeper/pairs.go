package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPairs set a specific pairs in the store from its index
func (k Keeper) SetPairs(ctx sdk.Context, pairs types.Pairs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PairsPrefix())
	b := k.cdc.MustMarshal(&pairs)
	store.Set(types.PairsKey(
		pairs.Token0,
		pairs.Token1,
	), b)
}

// GetPairs returns a pairs from its index
func (k Keeper) GetPairs(
	ctx sdk.Context,
	token0 string,
	token1 string,

) (val types.Pairs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PairsPrefix())

	b := store.Get(types.PairsKey(
		token0,
		token1,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairs removes a pairs from the store
func (k Keeper) RemovePairs(
	ctx sdk.Context,
	token0 string,
	token1 string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PairsPrefix())
	store.Delete(types.PairsKey(
		token0,
		token1,
	))
}

// GetAllPairs returns all pairs
func (k Keeper) GetAllPairs(ctx sdk.Context) (list []types.Pairs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PairsPrefix())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pairs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
