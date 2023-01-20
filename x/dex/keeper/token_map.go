package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetTokenMap set a specific tokenMap in the store from its index
func (k Keeper) SetTokenMap(ctx sdk.Context, tokenMap types.TokenMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenMapKeyPrefix))
	b := k.cdc.MustMarshal(&tokenMap)
	store.Set(types.TokenMapKey(
		tokenMap.Address,
	), b)
}

// GetTokenMap returns a tokenMap from its index
func (k Keeper) GetTokenMap(
	ctx sdk.Context,
	address string,

) (val types.TokenMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenMapKeyPrefix))

	b := store.Get(types.TokenMapKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenMap removes a tokenMap from the store
func (k Keeper) RemoveTokenMap(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenMapKeyPrefix))
	store.Delete(types.TokenMapKey(
		address,
	))
}

// GetAllTokenMap returns all tokenMap
func (k Keeper) GetAllTokenMap(ctx sdk.Context) (list []types.TokenMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
