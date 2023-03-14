package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetGoodTillRecord set a specific goodTillRecord in the store from its index
func (k Keeper) SetGoodTillRecord(ctx sdk.Context, goodTillRecord types.GoodTillRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))
	b := k.cdc.MustMarshal(&goodTillRecord)
	store.Set(types.GoodTillRecordKey(
		goodTillRecord.GoodTillDate,
		goodTillRecord.TrancheRef,
	), b)
}

// GetGoodTillRecord returns a goodTillRecord from its index
func (k Keeper) GetGoodTillRecord(
	ctx sdk.Context,
	goodTillDate string,
	trancheRef []byte,

) (val types.GoodTillRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))

	b := store.Get(types.GoodTillRecordKey(
		goodTillDate,
		trancheRef,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGoodTillRecord removes a goodTillRecord from the store
func (k Keeper) RemoveGoodTillRecord(
	ctx sdk.Context,
	goodTillDate string,
	trancheRef string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))
	store.Delete(types.GoodTillRecordKey(
		goodTillDate,
		trancheRef,
	))
}

// GetAllGoodTillRecord returns all goodTillRecord
func (k Keeper) GetAllGoodTillRecord(ctx sdk.Context) (list []types.GoodTillRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GoodTillRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
