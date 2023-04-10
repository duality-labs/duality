package keeper

import (
	"encoding/json"

	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
	db "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// iterator returns an iterator over all gauges in the {prefix} space of state.
func (k Keeper) iterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

// iterator returns an iterator over all gauges in the {prefix} space of state.
func (k Keeper) iteratorStartEnd(ctx sdk.Context, start []byte, end []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(start, end)
}

func UnmarshalRefArray(bz []byte) []uint64 {
	ids := []uint64{}
	err := json.Unmarshal(bz, &ids)
	if err != nil {
		panic(err)
	}
	return ids
}

// getLocksFromIterator returns an array of single lock units by period defined by the x/lockup module.
func (k Keeper) getLocksFromIterator(ctx sdk.Context, iterator db.Iterator) types.Locks {
	locks := types.Locks{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		lockIDs := UnmarshalRefArray(iterator.Value())
		for _, lockID := range lockIDs {
			lock, err := k.GetLockByID(ctx, lockID)
			if err != nil {
				panic(err)
			}
			locks = append(locks, lock)
		}
	}
	return locks
}

func (k Keeper) getLockByRefKey(ctx sdk.Context, key []byte) *types.Lock {
	store := ctx.KVStore(k.storeKey)
	lockRefArrayBz := store.Get(key)
	if lockRefArrayBz == nil {
		return nil
	}
	lockIDs := UnmarshalRefArray(lockRefArrayBz)
	if len(lockIDs) > 1 {
		panic("not expecting more than one here")
	}
	if len(lockIDs) == 0 {
		return nil
	}
	lock, err := k.GetLockByID(ctx, lockIDs[0])
	if err != nil {
		panic(err)
	}
	return lock
}

// LockIteratorDenom returns the iterator used for getting all locks by denom.
func (k Keeper) LockIteratorPairTick(ctx sdk.Context, isUnlocking bool, pairId *dextypes.PairID, startTickIndexInc int64, endTickIndexInc int64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	unlockingPrefix := types.GetPrefixLockStatus(isUnlocking)
	startTickIndexBz := dextypes.TickIndexToBytes(startTickIndexInc, pairId, pairId.Token1)
	endTickIndexBz := dextypes.TickIndexToBytes(endTickIndexInc+1, pairId, pairId.Token1)
	pairIdBz := []byte(pairId.Stringify())
	startPrefix := types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexPairTick, pairIdBz, startTickIndexBz)
	endPrefix := types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexPairTick, pairIdBz, endTickIndexBz)
	return store.Iterator(startPrefix, endPrefix)
}

// // AccountLockIteratorBeforeOrAtTime returns the iterator to get unlockable coins by account.
// func (k Keeper) AccountLockIteratorBeforeOrAtTime(ctx sdk.Context, addr sdk.AccAddress, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetPrefixLockStatus(true)
// 	// return k.iteratorBeforeOrAtTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccountTimestamp, addr), time)
// 	return k.iterator(prefix, storetypes.PrefixEndBytes(key))
// }

// // UpcomingGaugesIteratorAfterTime returns the iterator to get all upcoming gauges that start distribution after a specific time.
// func (k Keeper) UpcomingGaugesIteratorAfterTime(ctx sdk.Context, time time.Time) sdk.Iterator {
// 	return k.iteratorAfterTime(ctx, types.KeyPrefixGaugeIndexUpcoming, time)
// }

// // UpcomingGaugesiteratorBeforeOrAtTime returns the iterator to get all upcoming gauges that have already started distribution before a specific time.
// func (k Keeper) UpcomingGaugesIteratorBeforeOrAtTime(ctx sdk.Context, time time.Time) sdk.Iterator {
// 	return k.iteratorBeforeOrAtTime(ctx, types.KeyPrefixGaugeIndexUpcoming, time)
// }

// lock iterators

// // LockIteratorAfterTime returns the iterator to get locked coins.
// func (k Keeper) LockIteratorAfterTime(ctx sdk.Context, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorAfterTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexTimestamp), time)
// }

// // LockiteratorBeforeOrAtTime returns the iterator to get unlockable coins.
// func (k Keeper) LockiteratorBeforeOrAtTime(ctx sdk.Context, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorBeforeOrAtTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexTimestamp), time)
// }

// // LockIterator returns the iterator used for getting all locks.
// func (k Keeper) LockIterator(ctx sdk.Context, isUnlocking bool) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(isUnlocking)
// 	return k.iterator(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndex))
// }

// // LockIteratorAfterTimeDenom returns the iterator to get locked coins by denom.
// func (k Keeper) LockIteratorAfterTimeDenom(ctx sdk.Context, denom string, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorAfterTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexDenomTimestamp, []byte(denom)), time)
// }

// // LockiteratorBeforeOrAtTimeDenom returns the iterator to get unlockable coins by denom.
// func (k Keeper) LockiteratorBeforeOrAtTimeDenom(ctx sdk.Context, denom string, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorBeforeOrAtTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexDenomTimestamp, []byte(denom)), time)
// }

// // LockIteratorDenom returns the iterator used for getting all locks by denom.
// func (k Keeper) LockIteratorDenom(ctx sdk.Context, isUnlocking bool, denom string) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(isUnlocking)
// 	return k.iterator(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexDenom, []byte(denom)))
// }

// // AccountLockIteratorAfterTime returns the iterator to get locked coins by account.
// func (k Keeper) AccountLockIteratorAfterTime(ctx sdk.Context, addr sdk.AccAddress, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorAfterTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccountTimestamp, addr), time)
// }

// // AccountLockIterator returns the iterator used for getting all locks by account.
// func (k Keeper) AccountLockIterator(ctx sdk.Context, isUnlocking bool, addr sdk.AccAddress) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(isUnlocking)
// 	return k.iterator(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccount, addr))
// }

// // AccountLockIteratorAfterTimeDenom returns the iterator to get locked coins by account and denom.
// func (k Keeper) AccountLockIteratorAfterTimeDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorAfterTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccountDenomTimestamp, addr, []byte(denom)), time)
// }

// // AccountLockIteratorBeforeOrAtTimeDenom returns the iterator to get unlockable coins by account and denom.
// func (k Keeper) AccountLockIteratorBeforeOrAtTimeDenom(ctx sdk.Context, addr sdk.AccAddress, denom string, time time.Time) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(true)
// 	return k.iteratorBeforeOrAtTime(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccountDenomTimestamp, addr, []byte(denom)), time)
// }

// // AccountLockIteratorDenom returns the iterator used for getting all locks by account and denom.
// func (k Keeper) AccountLockIteratorDenom(ctx sdk.Context, isUnlocking bool, addr sdk.AccAddress, denom string) sdk.Iterator {
// 	unlockingPrefix := types.GetLockStatusPrefix(isUnlocking)
// 	return k.iterator(ctx, types.CombineKeys(unlockingPrefix, types.KeyPrefixLockIndexAccountDenom, addr, []byte(denom)))
// }

// // iteratorAfterTime iterates through keys between that use prefix, and have a time.
// func (k Keeper) iteratorAfterTime(ctx sdk.Context, prefix []byte, time time.Time) sdk.Iterator {
// 	store := ctx.KVStore(k.storeKey)
// 	timeKey := types.GetTimeKey(time)
// 	key := types.CombineKeys(prefix, timeKey)
// 	// If it’s unlockTime, then it should count as unlocked
// 	// inclusive end bytes = key + 1, next iterator
// 	return store.Iterator(storetypes.PrefixEndBytes(key), storetypes.PrefixEndBytes(prefix))
// }

// // iteratorBeforeOrAtTime iterates through keys between that use prefix, and have a time LTE max time.
// func (k Keeper) iteratorBeforeOrAtTime(ctx sdk.Context, prefix []byte, maxTime time.Time) sdk.Iterator {
// 	store := ctx.KVStore(k.storeKey)
// 	timeKey := types.GetTimeKey(maxTime)
// 	key := types.CombineKeys(prefix, timeKey)
// 	// If it’s unlockTime, then it should count as unlocked
// 	// inclusive end bytes = key + 1, next iterator
// 	return store.Iterator(prefix, storetypes.PrefixEndBytes(key))
// }
