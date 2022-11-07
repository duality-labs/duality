package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTokenObject set a specific tokenObject in the store from its index
func (k Keeper) SetTokenObject(ctx sdk.Context, tokenObject types.TokenObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenObjectKeyPrefix))
	b := k.cdc.MustMarshal(&tokenObject)
	store.Set(types.TokenObjectKey(
		tokenObject.Address,
	), b)
}

// GetTokenObject returns a tokenObject from its index
func (k Keeper) GetTokenObject(
	ctx sdk.Context,
	address string,

) (val types.TokenObject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenObjectKeyPrefix))

	b := store.Get(types.TokenObjectKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenObject removes a tokenObject from the store
func (k Keeper) RemoveTokenObject(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenObjectKeyPrefix))
	store.Delete(types.TokenObjectKey(
		address,
	))
}

// GetAllTokenObject returns all tokenObject
func (k Keeper) GetAllTokenObject(ctx sdk.Context) (list []types.TokenObject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenObjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenObject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
