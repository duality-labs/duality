package keeper

import (
	"time"

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
	goodTillDate time.Time,
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
	goodTillDate time.Time,
	trancheRef []byte,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))
	store.Delete(types.GoodTillRecordKey(
		goodTillDate,
		trancheRef,
	))
}

func (k Keeper) RemoveGoodTillRecordByPrefixedyKey(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
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

func (k Keeper) PurgeExpiredGoodTillRecords(ctx sdk.Context, curTime time.Time) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTillRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	inGTDSegment := false

	archivedTranches := make(map[string]bool)
	defer iterator.Close()
	gasCutoff := ctx.BlockGasMeter().Limit() - types.GoodTillPurgeGasBuffer

	for ; iterator.Valid(); iterator.Next() {
		var val types.GoodTillRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GoodTillDate.Before(curTime) {
			inGTDSegment = inGTDSegment || val.GoodTillDate != types.JITGoodTillTime
			gasConsumed := ctx.BlockGasMeter().GasConsumed()

			if inGTDSegment && gasConsumed >= gasCutoff {
				// If we hit our gas cutoff stop deleting so as not to timeout the block.
				// We can only do this if we are proccesing normal GT limitOrders
				// and not JIT limit orders, since there is not protection in place
				// to prevent JIT order from being traded on the the next block.
				// This is ok since only GT limit orders pose a meaningful attack
				// vector since there is no upper bound on how many GT limit orders can be
				// canceled in a single block.
				ctx.EventManager().EmitEvent(types.GoodTillPurgeHitLimitEvent(gasConsumed))
				return
			}
			if _, ok := archivedTranches[string(val.TrancheRef)]; !ok {
				tranche, found := k.GetLimitOrderTrancheByKey(ctx, val.TrancheRef)
				// Convert the tranche to a filled tranche
				if found {
					k.SetFilledLimitOrderTranche(ctx, *tranche)
					k.RemoveLimitOrderTranche(ctx, *tranche)
					archivedTranches[string(val.TrancheRef)] = true
				}
			}
			k.RemoveGoodTillRecordByPrefixedyKey(ctx, iterator.Key())
		} else {
			return
		}
	}

	return
}
