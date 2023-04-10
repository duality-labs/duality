package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	"github.com/duality-labs/duality/x/incentives/types"
)

// GetLastLockID returns ID used last time.
func (k Keeper) GetLastLockID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyLastLockID)
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetLastLockID save ID used by last lock.
func (k Keeper) SetLastLockID(ctx sdk.Context, ID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastLockID, sdk.Uint64ToBigEndian(ID))
}

// WithdrawAllMaturedLocks withdraws every lock thats in the process of unlocking, and has finished unlocking by
// the current block time.
func (k Keeper) WithdrawAllMaturedLocks(ctx sdk.Context) {
	coins := sdk.Coins{}
	locks := k.getUnlockingLocksMatured(ctx)
	for _, lock := range locks {
		err := k.UnlockMaturedLock(ctx, lock.ID)
		if err != nil {
			panic(err)
		}
		// sum up all coins unlocked
		coins = coins.Add(lock.Coins...)
	}
}

// AddToExistingLock adds the given coin to the existing full lock with the same owner.
// Returns the updated lock ID if successfully added coin, returns 0 and error when a lock with
// given condition does not exist, or if fails to add to lock.
func (k Keeper) AddToExistingLock(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins) (uint64, error) {
	lock := k.getLockByRefKey(ctx, types.GetKeyLockIndexByAccount(false, owner))

	// if no lock exists for the given owner + denom, return an error
	if lock == nil {
		return 0, sdkerrors.Wrapf(types.ErrLockupNotFound, "lock with owner %s does not exist", owner)
	}

	// if existing lock with same duration and denom exists, add to the existing lock
	// there should only be a single lock with the same duration + token, thus we take the first lock
	_, err := k.AddTokensToLockByID(ctx, lock.ID, owner, coins)
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return lock.ID, nil
}

// HasLock returns true if lock with the given condition exists
func (k Keeper) HasFullLock(ctx sdk.Context, owner sdk.AccAddress) bool {
	// locks := k.GetAccountFullyLocked(ctx, owner, denom)
	lock := k.getLockByRefKey(ctx, types.GetKeyLockIndexByAccount(false, owner))
	return lock != nil
}

// AddTokensToLock locks additional tokens into an existing lock with the given ID.
// Tokens locked are sent and kept in the module account.
// This method alters the lock state in store, thus we do a sanity check to ensure
// lock owner matches the given owner.
func (k Keeper) AddTokensToLockByID(ctx sdk.Context, lockID uint64, owner sdk.AccAddress, tokensToAdd sdk.Coins) (*types.Lock, error) {
	lock, err := k.GetLockByID(ctx, lockID)
	if err != nil {
		return nil, err
	}

	if lock.GetOwner() != owner.String() {
		return nil, types.ErrNotLockOwner
	}

	lock.Coins = lock.Coins.Add(tokensToAdd...)
	err = k.Lock(ctx, lock, tokensToAdd)
	if err != nil {
		return nil, err
	}

	if k.hooks == nil {
		return lock, nil
	}

	k.hooks.AfterAddTokensToLock(ctx, lock.OwnerAddress(), lock.GetID(), tokensToAdd)

	return lock, nil
}

// lock is an internal utility to lock coins and set corresponding states.
// This is only called by either of the two possible entry points to lock tokens.
// 1. CreateLock
// 2. AddTokensToLockByID
func (k Keeper) Lock(ctx sdk.Context, lock *types.Lock, tokensToLock sdk.Coins) error {
	owner, err := sdk.AccAddressFromBech32(lock.Owner)
	if err != nil {
		return err
	}

	err = lock.ValidateBasic()
	if err != nil {
		return err
	}

	if err := k.bk.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, tokensToLock); err != nil {
		return err
	}

	// store lock object into the store
	err = k.setLock(ctx, lock)
	if err != nil {
		return err
	}

	k.hooks.OnTokenLocked(ctx, owner, lock.ID, lock.Coins, lock.Duration, lock.EndTime)
	return nil
}

// BeginUnlock is a utility to start unlocking coins from NotUnlocking queue.
func (k Keeper) BeginUnlock(ctx sdk.Context, lockID uint64, coins sdk.Coins) (uint64, error) {
	lock, err := k.GetLockByID(ctx, lockID)
	if err != nil {
		return 0, err
	}

	if !coins.IsAllLTE(lock.Coins) {
		return 0, fmt.Errorf("requested amount to unlock exceeds locked tokens")
	}

	if lock.IsUnlocking() {
		return 0, fmt.Errorf("trying to unlock a lock that is already unlocking")
	}

	// If the amount were unlocking is empty, or the entire coins amount, unlock the entire lock.
	// Otherwise, split the lock into two locks, and fully unlock the newly created lock.
	// (By virtue, the newly created lock we split into should have the unlock amount)
	if len(coins) != 0 && !coins.IsEqual(lock.Coins) {
		splitLock, err := k.splitLock(ctx, lock, coins)
		if err != nil {
			return 0, err
		}
		lock = splitLock
	}

	// remove existing lock refs from not unlocking queue
	err = k.deleteLockRefs(ctx, lock)
	if err != nil {
		return 0, err
	}

	// store lock with the end time set to current block time + duration
	lock.EndTime = ctx.BlockTime().Add(lock.Duration)
	err = k.setLock(ctx, lock)
	if err != nil {
		return 0, err
	}

	// add lock refs into unlocking queue
	err = k.addLockRefs(ctx, lock)
	if err != nil {
		return 0, err
	}

	if k.hooks != nil {
		k.hooks.OnStartUnlock(ctx, lock.OwnerAddress(), lock.ID, lock.Coins, lock.Duration, lock.EndTime)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		createBeginUnlockEvent(lock),
	})

	return lock.ID, err
}

// UnlockMaturedLock finishes unlocking by sending back the locked tokens from the module accounts
// to the owner. This method requires lock to be matured, having passed the endtime of the lock.
func (k Keeper) UnlockMaturedLock(ctx sdk.Context, lockID uint64) error {
	lock, err := k.GetLockByID(ctx, lockID)
	if err != nil {
		return err
	}

	// validation for current time and unlock time
	curTime := ctx.BlockTime()
	if !lock.IsUnlocking() {
		return fmt.Errorf("lock hasn't started unlocking yet")
	}
	if curTime.Before(lock.EndTime) {
		return fmt.Errorf("lock is not unlockable yet: %s >= %s", curTime.String(), lock.EndTime.String())
	}

	owner, err := sdk.AccAddressFromBech32(lock.Owner)
	if err != nil {
		return err
	}

	// send coins back to owner
	if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, lock.Coins); err != nil {
		return err
	}

	k.deleteLock(ctx, lock.ID)

	// delete lock refs from the unlocking queue
	err = k.deleteLockRefs(ctx, lock)
	if err != nil {
		return err
	}

	k.hooks.OnTokenUnlocked(ctx, owner, lock.ID, lock.Coins, lock.Duration, lock.EndTime)
	return nil
}

// setLock is a utility to store lock object into the store.
func (k Keeper) setLock(ctx sdk.Context, lock *types.Lock) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := proto.Marshal(lock)
	if err != nil {
		return err
	}
	store.Set(types.GetLockStoreKey(lock.ID), bz)
	return nil
}

// deleteLock removes the lock object from the state.
func (k Keeper) deleteLock(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLockStoreKey(id))
}

// splitLock splits a lock with the given amount, and stores split new lock to the state.
// Returns the new lock after modifying the state of the old lock.
func (k Keeper) splitLock(ctx sdk.Context, lock *types.Lock, coins sdk.Coins) (*types.Lock, error) {
	if lock.IsUnlocking() {
		return nil, fmt.Errorf("cannot split unlocking lock")
	}

	// TODO: Manage removing coin refs
	lock.Coins = lock.Coins.Sub(coins)

	err := k.setLock(ctx, lock)
	if err != nil {
		return nil, err
	}

	// create a new lock
	return k.CreateLock(ctx, lock.OwnerAddress(), coins, lock.Duration)
}

// GetLockByID Returns lock from lockID.
func (k Keeper) GetLockByID(ctx sdk.Context, lockID uint64) (*types.Lock, error) {
	lock := types.Lock{}
	store := ctx.KVStore(k.storeKey)
	lockKey := types.GetLockStoreKey(lockID)
	if !store.Has(lockKey) {
		return nil, sdkerrors.Wrap(types.ErrLockupNotFound, fmt.Sprintf("lock with ID %d does not exist", lockID))
	}
	bz := store.Get(lockKey)
	err := proto.Unmarshal(bz, &lock)
	return &lock, err
}

// GetAccountLocks Returns the period locks associated to an account.
func (k Keeper) GetLocksByQueryCondition(ctx sdk.Context, distrTo *types.QueryCondition) types.Locks {
	unlockings := k.getLocksFromIterator(
		ctx,
		k.LockIteratorPairTick(
			ctx,
			true,
			distrTo.PairID,
			distrTo.StartTick,
			distrTo.EndTick,
		),
	)
	notUnlockings := k.getLocksFromIterator(
		ctx,
		k.LockIteratorPairTick(
			ctx,
			false,
			distrTo.PairID,
			distrTo.StartTick,
			distrTo.EndTick,
		),
	)
	return append(notUnlockings, unlockings...)
}

func (k Keeper) getUnlockingLocksMatured(ctx sdk.Context) types.Locks {
	start := types.GetKeyLockIndexUnlockingByTimestamp(time.Time{})
	end := sdk.PrefixEndBytes(types.GetKeyLockIndexUnlockingByTimestamp(ctx.BlockTime()))
	return k.getLocksFromIterator(ctx, k.iteratorStartEnd(ctx, start, end))
}

func (k Keeper) getUnlockingLocksNotMatured(ctx sdk.Context) types.Locks {
	start := sdk.PrefixEndBytes(types.GetKeyLockIndexUnlockingByTimestamp(ctx.BlockTime()))
	// TODO: make more elegant, this breaks out of the established pattern
	end := sdk.PrefixEndBytes(types.CombineKeys(
		types.GetPrefixLockStatus(true),
		types.KeyPrefixLockIndexTimestamp,
	))
	return k.getLocksFromIterator(ctx, k.iteratorStartEnd(ctx, start, end))
}

func (k Keeper) getUnlockingLocks(ctx sdk.Context) types.Locks {
	return k.getLocksFromIterator(ctx, k.iterator(ctx, types.GetKeyLockIndex(true)))
}

func (k Keeper) getFullLocks(ctx sdk.Context) types.Locks {
	return k.getLocksFromIterator(ctx, k.iterator(ctx, types.GetKeyLockIndex(false)))
}

// GetLocks Returns the period locks on pool.
func (k Keeper) GetLocks(ctx sdk.Context) (types.Locks, error) {
	unlockings := k.getUnlockingLocks(ctx)
	notUnlockings := k.getFullLocks(ctx)
	return append(notUnlockings, unlockings...), nil
}

// GetAccountLocks Returns the period locks associated to an account.
func (k Keeper) GetLocksByAccount(ctx sdk.Context, addr sdk.AccAddress) types.Locks {
	unlockings := k.getLocksFromIterator(ctx, k.iterator(ctx, types.GetKeyLockIndexByAccount(true, addr)))
	notUnlockings := k.getLocksFromIterator(ctx, k.iterator(ctx, types.GetKeyLockIndexByAccount(false, addr)))
	return append(notUnlockings, unlockings...)
}

func (k Keeper) CreateLock(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins, duration time.Duration) (*types.Lock, error) {
	ID := k.GetLastLockID(ctx) + 1

	// unlock time is initially set without a value, gets set as unlock start time + duration
	// when unlocking starts.
	lock := types.NewLock(ID, owner, duration, time.Time{}, coins)
	err := k.Lock(ctx, lock, lock.Coins)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// add lock refs into not unlocking queue
	err = k.addLockRefs(ctx, lock)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	k.SetLastLockID(ctx, lock.ID)
	return lock, nil
}
