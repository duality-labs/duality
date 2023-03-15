package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

// SetGoodTilRecord set a specific goodTilRecord in the store from its index
func (k Keeper) SetGoodTilRecord(ctx sdk.Context, goodTilRecord types.GoodTilRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTilRecordKeyPrefix))
	b := k.cdc.MustMarshal(&goodTilRecord)
	store.Set(types.GoodTilRecordKey(
		goodTilRecord.GoodTilDate,
		goodTilRecord.TrancheRef,
	), b)
}

// GetGoodTilRecord returns a goodTilRecord from its index
func (k Keeper) GetGoodTilRecord(
	ctx sdk.Context,
	goodTilDate time.Time,
	trancheRef []byte,

) (val types.GoodTilRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTilRecordKeyPrefix))

	b := store.Get(types.GoodTilRecordKey(
		goodTilDate,
		trancheRef,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGoodTilRecord removes a goodTilRecord from the store
func (k Keeper) RemoveGoodTilRecord(
	ctx sdk.Context,
	goodTilDate time.Time,
	trancheRef []byte,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTilRecordKeyPrefix))
	store.Delete(types.GoodTilRecordKey(
		goodTilDate,
		trancheRef,
	))
}

func (k Keeper) RemoveGoodTilRecordByPrefixedyKey(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
}

// GetAllGoodTilRecord returns all goodTilRecord
func (k Keeper) GetAllGoodTilRecord(ctx sdk.Context) (list []types.GoodTilRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTilRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GoodTilRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) PurgeExpiredGoodTilRecords(ctx sdk.Context, curTime time.Time) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoodTilRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	inGTDSegment := false

	archivedTranches := make(map[string]bool)
	defer iterator.Close()
	gasCutoff := ctx.BlockGasMeter().Limit() - types.GoodTilPurgeGasBuffer

	for ; iterator.Valid(); iterator.Next() {
		var val types.GoodTilRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GoodTilDate.Before(curTime) {
			inGTDSegment = inGTDSegment || val.GoodTilDate != types.JITGoodTilTime
			gasConsumed := ctx.BlockGasMeter().GasConsumed()

			if inGTDSegment && gasConsumed >= gasCutoff {
				// If we hit our gas cutoff stop deleting so as not to timeout the block.
				// We can only do this if we are proccesing normal GT limitOrders
				// and not JIT limit orders, since there is not protection in place
				// to prevent JIT order from being traded on the the next block.
				// This is ok since only GT limit orders pose a meaningful attack
				// vector since there is no upper bound on how many GT limit orders can be
				// canceled in a single block.
				ctx.EventManager().EmitEvent(types.GoodTilPurgeHitLimitEvent(gasConsumed))
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
			k.RemoveGoodTilRecordByPrefixedyKey(ctx, iterator.Key())
		} else {
			return
		}
	}

	return
}
