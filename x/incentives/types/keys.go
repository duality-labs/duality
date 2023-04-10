package types

import (
	"bytes"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// ModuleName defines the module name.
	ModuleName = "incentives"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_capability"

	// KeyPrefixTimestamp defines prefix key for timestamp iterator key.
	KeyPrefixTimestamp = []byte{0x01}

	// KeyLastGaugeID defines key for setting last gauge ID.
	KeyLastGaugeID = []byte{0x02}

	// KeyPrefixGauge defines prefix key for storing gauges.
	KeyPrefixGauge = []byte{0x03}

	// KeyPrefixGaugeIndex defines prefix key for storing reference key for all gauges.
	KeyPrefixGaugeIndex = []byte{0x04}

	// KeyPrefixGaugeIndexUpcoming defines prefix key for storing reference key for upcoming gauges.
	KeyPrefixGaugeIndexUpcoming = []byte{0x04, 0x00}

	// KeyPrefixGaugeIndexActive defines prefix key for storing reference key for active gauges.
	KeyPrefixGaugeIndexActive = []byte{0x04, 0x01}

	// KeyPrefixGaugeIndexFinished defines prefix key for storing reference key for finished gauges.
	KeyPrefixGaugeIndexFinished = []byte{0x04, 0x02}

	// KeyPrefixGaugeIndexByPair defines prefix key for storing indexes of gauge IDs by denomination.
	KeyPrefixGaugeIndexByPair = []byte{0x05}

	// KeyLastLockID defines key to store lock ID used by last.
	KeyLastLockID = []byte{0x06}

	// KeyPrefixLock defines prefix to store period lock by ID.
	KeyPrefixLock = []byte{0x07}

	// KeyPrefixLockIndexNotUnlocking defines prefix to query iterators which hasn't started unlocking.
	KeyPrefixLockIndexNotUnlocking = []byte{0x08}

	// KeyPrefixLockIndexUnlocking defines prefix to query iterators which has started unlocking.
	KeyPrefixLockIndexUnlocking = []byte{0x09}

	// KeyPrefixLockIndex defines prefix for the iteration of lock IDs.
	KeyPrefixLockIndex = []byte{0x0a}

	// KeyPrefixLockIndexAccount defines prefix for the iteration of lock IDs by account.
	KeyPrefixLockIndexAccount = []byte{0x0b}

	// KeyPrefixLockIndexDenom defines prefix for the iteration of lock IDs by denom.
	KeyPrefixLockIndexDenom = []byte{0x0c}

	// KeyPrefixLockIndexPairTick defines prefix for the iteration of lock IDs by pairId and tick index.
	KeyPrefixLockIndexPairTick = []byte{0x0d}

	// KeyPrefixLockIndexAccountDenom defines prefix for the iteration of lock IDs by account, denomination.
	KeyPrefixLockIndexAccountDenom = []byte{0x0e}

	// KeyPrefixLockIndexTimestamp defines prefix for the iteration of lock IDs by timestamp.
	KeyPrefixLockIndexTimestamp = []byte{0x0f}

	// KeyPrefixLockIndexAccountTimestamp defines prefix for the iteration of lock IDs by account and timestamp.
	KeyPrefixLockIndexAccountTimestamp = []byte{0x10}

	// KeyPrefixLockIndexDenomTimestamp defines prefix for the iteration of lock IDs by denom and timestamp.
	KeyPrefixLockIndexDenomTimestamp = []byte{0x11}

	// KeyPrefixLockIndexAccountDenomTimestamp defines prefix for the iteration of lock IDs by account, denomination and timestamp.
	KeyPrefixLockIndexAccountDenomTimestamp = []byte{0x12}

	// KeyndexSeparator defines separator between keys when combine, it should be one that is not used in denom expression.
	KeyIndexSeparator = []byte{0xFF}
)

// lockStoreKey returns action store key from ID.
func GetLockStoreKey(ID uint64) []byte {
	return CombineKeys(KeyPrefixLock, sdk.Uint64ToBigEndian(ID))
}

// combineKeys combine bytes array into a single bytes.
func CombineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, KeyIndexSeparator)
}

// getTimeKey returns the key used for getting a set of period locks
// where unlockTime is after a specific time.
func GetTimeKey(timestamp time.Time) []byte {
	timeBz := sdk.FormatTimeBytes(timestamp)
	timeBzL := len(timeBz)
	prefixL := len(KeyPrefixTimestamp)

	bz := make([]byte, prefixL+8+timeBzL)

	// copy the prefix
	copy(bz[:prefixL], KeyPrefixTimestamp)

	// copy the encoded time bytes length
	copy(bz[prefixL:prefixL+8], sdk.Uint64ToBigEndian(uint64(timeBzL)))

	// copy the encoded time bytes
	copy(bz[prefixL+8:prefixL+8+timeBzL], timeBz)
	return bz
}

func CombineLocks(pl1 Locks, pl2 Locks) Locks {
	return append(pl1, pl2...)
}

// gaugeStoreKey returns the combined byte array (store key) of the provided gauge ID's key prefix and the ID itself.
func GetKeyGaugeStore(ID uint64) []byte {
	return CombineKeys(KeyPrefixGauge, sdk.Uint64ToBigEndian(ID))
}

// gaugePairStoreKey returns the combined byte array (store key) of the provided gauge denom key prefix and the denom itself.
func GetKeyGaugeIndexByPair(pairID string) []byte {
	return CombineKeys(KeyPrefixGaugeIndexByPair, []byte(pairID))
}

func GetPrefixLockStatus(isUnlocking bool) []byte {
	if isUnlocking {
		return KeyPrefixLockIndexUnlocking
	}
	return KeyPrefixLockIndexNotUnlocking
}

func GetKeyLockIndex(isUnlocking bool) []byte {
	return CombineKeys(
		GetPrefixLockStatus(isUnlocking),
		KeyPrefixLockIndex,
	)
}

func GetKeyLockIndexByAccount(isUnlocking bool, account sdk.AccAddress) []byte {
	return CombineKeys(
		GetPrefixLockStatus(isUnlocking),
		KeyPrefixLockIndexAccount,
		account,
	)
}

// TODO: revisit whether denom index is necessary
func GetKeyLockIndexByDenom(isUnlocking bool, denom string) []byte {
	return CombineKeys(
		GetPrefixLockStatus(isUnlocking),
		KeyPrefixLockIndexDenom,
		[]byte(denom),
	)
}

func GetKeyLockIndexByAccountDenom(isUnlocking bool, account sdk.AccAddress, denom string) []byte {
	return CombineKeys(
		GetPrefixLockStatus(isUnlocking),
		KeyPrefixLockIndexAccountDenom,
		account,
		[]byte(denom),
	)
}

func GetKeyLockIndexUnlockingByTimestamp(timestamp time.Time) []byte {
	return CombineKeys(
		GetPrefixLockStatus(true),
		KeyPrefixLockIndexTimestamp,
		GetTimeKey(timestamp),
	)
}

func GetKeyLockIndexUnlockingByAccountTimestamp(account sdk.AccAddress, timestamp time.Time) []byte {
	return CombineKeys(
		GetPrefixLockStatus(true),
		KeyPrefixLockIndexAccountTimestamp,
		account,
		GetTimeKey(timestamp),
	)
}

func GetKeyLockIndexUnlockingByDenomTimestamp(denom string, timestamp time.Time) []byte {
	return CombineKeys(
		GetPrefixLockStatus(true),
		KeyPrefixLockIndexDenomTimestamp,
		[]byte(denom),
		GetTimeKey(timestamp),
	)
}

func GetKeyLockIndexUnlockingByAccountDenomTimestamp(account sdk.AccAddress, denom string, timestamp time.Time) []byte {
	return CombineKeys(
		GetPrefixLockStatus(true),
		KeyPrefixLockIndexAccountDenomTimestamp,
		account,
		[]byte(denom),
		GetTimeKey(timestamp),
	)
}

// func GetKeyLockIndexByPairTick(isUnlocking bool, pairID string, tickIndex int64) []byte {
// 	return CombineKeys(
// 		GetPrefixLockStatus(isUnlocking),
// 		KeyPrefixLockIndexPairTick,
// 		// TODO
// 		[]byte(pairID),
// 		Int64ToBytes(tickIndex),
// 	)
// }
