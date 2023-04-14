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

	// KeyLastStakeID defines key to store stake ID used by last.
	KeyLastStakeID = []byte{0x06}

	// KeyPrefixStake defines prefix to store period stake by ID.
	KeyPrefixStake = []byte{0x07}

	// KeyPrefixStakeIndexAccount defines prefix for the iteration of stake IDs by account.
	KeyPrefixStakeIndex = []byte{0x08}

	// KeyPrefixStakeIndexAccount defines prefix for the iteration of stake IDs by account.
	KeyPrefixStakeIndexAccount = []byte{0x09}

	// KeyPrefixStakeIndexDenom defines prefix for the iteration of stake IDs by denom.
	KeyPrefixStakeIndexDenom = []byte{0x0c}

	// KeyPrefixStakeIndexPairTick defines prefix for the iteration of stake IDs by pairId and tick index.
	KeyPrefixStakeIndexPairTick = []byte{0x0d}

	// KeyPrefixStakeIndexAccountDenom defines prefix for the iteration of stake IDs by account, denomination.
	KeyPrefixStakeIndexAccountDenom = []byte{0x0e}

	// KeyPrefixStakeIndexTimestamp defines prefix for the iteration of stake IDs by timestamp.
	KeyPrefixStakeIndexPairTimestamp = []byte{0x0f}

	// // KeyPrefixStakeIndexAccountTimestamp defines prefix for the iteration of stake IDs by account and timestamp.
	// KeyPrefixStakeIndexAccountTimestamp = []byte{0x10}

	// // KeyPrefixStakeIndexDenomTimestamp defines prefix for the iteration of stake IDs by denom and timestamp.
	// KeyPrefixStakeIndexDenomTimestamp = []byte{0x11}

	// // KeyPrefixStakeIndexAccountDenomTimestamp defines prefix for the iteration of stake IDs by account, denomination and timestamp.
	// KeyPrefixStakeIndexAccountDenomTimestamp = []byte{0x12}

	// KeyndexSeparator defines separator between keys when combine, it should be one that is not used in denom expression.
	KeyIndexSeparator = []byte{0xFF}
)

// stakeStoreKey returns action store key from ID.
func GetStakeStoreKey(ID uint64) []byte {
	return CombineKeys(KeyPrefixStake, sdk.Uint64ToBigEndian(ID))
}

// combineKeys combine bytes array into a single bytes.
func CombineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, KeyIndexSeparator)
}

// getTimeKey returns the key used for getting a set of period stakes
// where unstakeTime is after a specific time.
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

func CombineStakes(pl1 Stakes, pl2 Stakes) Stakes {
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

func GetKeyStakeIndexByAccount(account sdk.AccAddress) []byte {
	return CombineKeys(
		KeyPrefixStakeIndexAccount,
		account,
	)
}

func GetKeyStakeIndexByDenom(denom string) []byte {
	return CombineKeys(
		KeyPrefixStakeIndexDenom,
		[]byte(denom),
	)
}

func GetKeyStakeIndexByAccountDenom(account sdk.AccAddress, denom string) []byte {
	return CombineKeys(
		KeyPrefixStakeIndexAccountDenom,
		account,
		[]byte(denom),
	)
}

func GetKeyStakeIndexByTimestamp(pairID string, timestamp time.Time) []byte {
	return CombineKeys(
		KeyPrefixStakeIndexPairTimestamp,
		[]byte(pairID),
		GetTimeKey(timestamp),
	)
}

func GetKeyStakeIndexByPairTick(pairID string, tickIndex int64) []byte {
	return CombineKeys(
		KeyPrefixStakeIndexPairTick,
		[]byte(pairID),
		TickIndexToBytes(tickIndex),
	)
}

func TickIndexToBytes(tickIndex int64) []byte {
	key := make([]byte, 9)
	if tickIndex < 0 {
		copy(key[1:], sdk.Uint64ToBigEndian(uint64(tickIndex)))
	} else {
		copy(key[:1], []byte{0x01})
		copy(key[1:], sdk.Uint64ToBigEndian(uint64(tickIndex)))
	}
	return key
}
