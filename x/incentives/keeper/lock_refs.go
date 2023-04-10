package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// addLockRefs adds appropriate reference keys preceded by a prefix.
// A prefix indicates whether the lock is unlocking or not.
func (k Keeper) addLockRefs(ctx sdk.Context, lock *types.Lock) error {
	refKeys, err := getLockRefKeys(lock)
	if err != nil {
		return err
	}
	for _, refKey := range refKeys {
		if err := k.addRefByKey(ctx, refKey, lock.ID); err != nil {
			return err
		}
	}
	return nil
}

// deleteLockRefs deletes all the lock references of the lock with the given lock prefix.
func (k Keeper) deleteLockRefs(ctx sdk.Context, lock *types.Lock) error {
	refKeys, err := getLockRefKeys(lock)
	if err != nil {
		return err
	}
	for _, refKey := range refKeys {
		k.deleteRefByKey(ctx, refKey, lock.ID)
	}
	return nil
}

func getLockRefKeys(lock *types.Lock) ([][]byte, error) {
	owner, err := sdk.AccAddressFromBech32(lock.Owner)
	if err != nil {
		return nil, err
	}

	refKeys := [][]byte{}
	lockRefPrefix := types.GetPrefixLockStatus(lock.IsUnlocking())
	refKeys = append(refKeys, types.CombineKeys(lockRefPrefix, types.KeyPrefixLockIndex))
	refKeys = append(refKeys, types.CombineKeys(lockRefPrefix, types.KeyPrefixLockIndexAccount, owner))

	for _, coin := range lock.Coins {
		depositDenom, err := dextypes.NewDepositDenomFromString(coin.Denom)
		if err != nil {
			panic("Only valid LP tokens should be locked")
		}
		denomBz := []byte(coin.Denom)
		pairIdBz := []byte(depositDenom.PairID.Stringify())
		tickBz := dextypes.TickIndexToBytes(depositDenom.Tick, depositDenom.PairID, depositDenom.PairID.Token1)
		refKeys = append(refKeys, types.CombineKeys(lockRefPrefix, types.KeyPrefixLockIndexDenom, denomBz))
		refKeys = append(refKeys, types.CombineKeys(lockRefPrefix, types.KeyPrefixLockIndexPairTick, pairIdBz, tickBz))
		refKeys = append(refKeys, types.CombineKeys(lockRefPrefix, types.KeyPrefixLockIndexAccountDenom, owner, denomBz))
	}

	if lock.IsUnlocking() {
		moreRefKeys, err := unlockingLockRefKeys(lock)
		if err != nil {
			return nil, err
		}
		refKeys = append(refKeys, moreRefKeys...)
	}

	return refKeys, nil
}

func unlockingLockRefKeys(lock *types.Lock) ([][]byte, error) {
	timeKey := types.GetTimeKey(lock.EndTime)

	owner, err := sdk.AccAddressFromBech32(lock.Owner)
	if err != nil {
		return nil, err
	}

	refKeys := [][]byte{}
	refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixLockIndexUnlocking, types.KeyPrefixLockIndexTimestamp, timeKey))
	refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixLockIndexUnlocking, types.KeyPrefixLockIndexAccountTimestamp, owner, timeKey))

	for _, coin := range lock.Coins {
		denomBz := []byte(coin.Denom)
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixLockIndexUnlocking, types.KeyPrefixLockIndexDenomTimestamp, denomBz, timeKey))
		refKeys = append(refKeys, types.CombineKeys(types.KeyPrefixLockIndexUnlocking, types.KeyPrefixLockIndexAccountDenomTimestamp, owner, denomBz, timeKey))
	}
	return refKeys, nil
}
