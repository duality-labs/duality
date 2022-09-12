package types

import "strconv"

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"

	Separator = "/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TokensKey      = "Tokens-value-"
	TokensCountKey = "Tokens-count-"

	// TokenMapKeyPrefix is the prefix to retrieve all TokenMap
	TokenMapKeyPrefix = "TokenMap/value/"

	// TickMapKeyPrefix is the prefix to retrieve all TickMap
	BaseTickMapKeyPrefix = "TickMap/value/"

	// PairMapKeyPrefix is the prefix to retrieve all PairMap
	PairMapKeyPrefix = "PairMap/value/"

	// SharesKeyPrefix is the prefix to retrieve all Shares
	SharesKeyPrefix = "Shares/value/"
)

func TickPrefix(pairId string) []byte {
	return append(KeyPrefix(BaseTickMapKeyPrefix), KeyPrefix(pairId)...)
}

// TokenMapKey returns the store key to retrieve a TokenMap from the index fields
func TokenMapKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickMapKey(tickIndex int64) []byte {
	var key []byte

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func PairMapKey(pairId string) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

// SharesKey returns the store key to retrieve a Shares from the index fields
func SharesKey(address string, pairId string, priceIndex int64, feeIndex uint64) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	priceIndexBytes := []byte(string(priceIndex))
	key = append(key, priceIndexBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(string(feeIndex))
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// Deposit Event Attributes
const (
	DepositEventKey          = "NewDeposit"
	DepositEventCreator      = "Creator"
	DepositEventToken0       = "Token0"
	DepositEventToken1       = "Token1"
	DepositEventPrice        = "Price"
	DepositEventFee          = "Fee"
	DepositEventReceiver     = "Receiver"
	DepositEventOldReserves0 = "OldReserves0"
	DepositEventOldReserves1 = "OldReserves1"
	DepositEventNewReserves0 = "NewReserves0"
	DepositEventNewReserves1 = "NewReserves1"
	DepositEventSharesMinted = "SharesMinted"
)

// Withdraw Event Attributes
const (
	WithdrawEventKey           = "NewWithdraw"
	WithdrawEventCreator       = "Creator"
	WithdrawEventToken0        = "Token0"
	WithdrawEventToken1        = "Token1"
	WithdrawEventPrice         = "Price"
	WithdrawEventFee           = "Fee"
	WithdrawEventReceiver      = "Receiver"
	WithdrawEventOldReserve0   = "OldReserve0"
	WithdrawEventOldReserve1   = "OldReserve0"
	WithdrawEventNewReserve0   = "NewReserve0"
	WithdrawEventNewReserve1   = "NewReserve1"
	WithdrawEventSharesRemoved = "SharesRemoved"
)

const (
	FeeListKey      = "FeeList-value-"
	FeeListCountKey = "FeeList-count-"
)

const (
	EdgeRowKey      = "EdgeRow-value-"
	EdgeRowCountKey = "EdgeRow-count-"
)

const (
	AdjanceyMatrixKey      = "AdjanceyMatrix-value-"
	AdjanceyMatrixCountKey = "AdjanceyMatrix-count-"
)
