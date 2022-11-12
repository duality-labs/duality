package types

import (
	fmt "fmt"
	"strconv"
)

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
	key := []byte(p)
	key = append(key, []byte("/")...)
	return key
}

const (
	TokensKey      = "Tokens-value-"
	TokensCountKey = "Tokens-count-"

	// TokenMapKeyPrefix is the prefix to retrieve all TokenMap
	TokenMapKeyPrefix = "TokenMap/value"

	// TickMapKeyPrefix is the prefix to retrieve all TickMap
	BaseTickMapKeyPrefix = "TickMap/value"

	// PairMapKeyPrefix is the prefix to retrieve all PairMap
	PairMapKeyPrefix = "PairMap/value"

	// SharesKeyPrefix is the prefix to retrieve all Shares
	SharesKeyPrefix = "Shares/value"
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

func TickMapKey(pairId string, tickIndex int64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

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
func SharesKey(address string, pairId string, tickIndex int64, feeIndex uint64) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(fmt.Sprint(tickIndex))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fmt.Sprint(feeIndex))
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// Limit Order Pool Mappings and Keys
const (
	BaseLimitOrderPrefix = "LimitOrderPool/value"

	// LimitOrderPoolUserSharesWithdrawnKeyPrefix is the prefix to retrieve all LimitOrderPoolUserSharesWithdrawn
	LimitOrderPoolUserSharesWithdrawnKeyPrefix = "LimitOrderPoolUserSharesWithdrawn/value"

	// LimitOrderPoolUserShareMapKeyPrefix is the prefix to retrieve all LimitOrderPoolUserShareMap
	LimitOrderPoolUserShareMapKeyPrefix = "LimitOrderPoolUserShareMap/value"

	// LimitOrderPoolTotalSharesMapKeyPrefix is the prefix to retrieve all LimitOrderPoolTotalSharesMap
	LimitOrderPoolTotalSharesMapKeyPrefix = "LimitOrderPoolTotalSharesMap/value"

	// LimitOrderPoolReserveMapKeyPrefix is the prefix to retrieve all LimitOrderPoolReserveMap
	LimitOrderPoolReserveMapKeyPrefix = "LimitOrderPoolReserveMap/value"

	// LimitOrderPoolFillMapKeyPrefix is the prefix to retrieve all LimitOrderPoolFillMap
	LimitOrderPoolFillMapKeyPrefix = "LimitOrderPoolFillMap/value"
)

// LimitOrderPoolUserSharesWithdrawnKey returns the store key to retrieve a LimitOrderPoolUserSharesWithdrawn from the index fields
func LimitOrderPoolUserSharesWithdrawnKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	countBytes := []byte(strconv.Itoa(int(count)))
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// LimitOrderPoolUserShareMapKey returns the store key to retrieve a LimitOrderPoolUserShareMap from the index fields
func LimitOrderPoolUserShareMapKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	countBytes := []byte(strconv.Itoa(int(count)))
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// LimitOrderPoolTotalSharesMapKey returns the store key to retrieve a LimitOrderPoolTotalSharesMap from the index fields
func LimitOrderPoolTotalSharesMapKey(pairId string, tickIndex int64, token string, count uint64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	countBytes := []byte(strconv.Itoa(int(count)))
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}

func LimitOrderPoolReserveMapKey(pairId string, tickIndex int64, token string, count uint64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	countBytes := []byte(strconv.Itoa(int(count)))
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}

func LimitOrderPoolFillMapKey(pairId string, tickIndex int64, token string, count uint64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	countBytes := []byte(strconv.Itoa(int(count)))
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}

// Deposit Event Attributes
const (
	DepositEventKey          = "NewDeposit"
	DepositEventCreator      = "Creator"
	DepositEventToken0       = "Token0"
	DepositEventToken1       = "Token1"
	DepositEventPrice        = "TickIndex"
	DepositEventFeeIndex     = "FeeIndex"
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
	WithdrawEventPrice         = "TickIndex"
	WithdrawEventFee           = "FeeIndex"
	WithdrawEventReceiver      = "Receiver"
	WithdrawEventOldReserve0   = "OldReserve0"
	WithdrawEventOldReserve1   = "OldReserve0"
	WithdrawEventNewReserve0   = "NewReserve0"
	WithdrawEventNewReserve1   = "NewReserve1"
	WithdrawEventSharesRemoved = "SharesRemoved"
)

const (
	SwapEventKey      = "NewSwap"
	SwapEventCreator  = "Creator"
	SwapEventReceiver = "Receiver"
	SwapEventToken0   = "Token0"
	SwapEventToken1   = "Token1"
	SwapEventTokenIn  = "TokenIn"
	SwapEventAmountIn = "AmountIn"
	SwapEventMinOut   = "MinOut"
	SwapEventAmoutOut = "AmountOut"
)

const (
	PlaceLimitOrderEventKey        = "NewPlaceLimitOrder"
	PlaceLimitOrderEventCreator    = "Creator"
	PlaceLimitOrderEventReceiver   = "Receiver"
	PlaceLimitOrderEventToken0     = "Token0"
	PlaceLimitOrderEventToken1     = "Token1"
	PlaceLimitOrderEventTokenIn    = "TokenIn"
	PlaceLimitOrderEventAmountIn   = "AmountIn"
	PlaceLimitOrderEventShares     = "Shares"
	PlaceLimitOrderEventCurrentKey = "CurrentLimitOrderKey"
)

const (
	WithdrawFilledLimitOrderEventKey           = "NewWithdraw"
	WithdrawFilledLimitOrderEventCreator       = "Creator"
	WithdrawFilledLimitOrderEventReceiver      = "Receiver"
	WithdrawFilledLimitOrderEventToken0        = "Token0"
	WithdrawFilledLimitOrderEventToken1        = "Token1"
	WithdrawFilledLimitOrderEventTokenKey      = "TokenKey"
	WithdrawFilledLimitOrderEventLimitOrderKey = "LimitOrderKey"
	WithdrawFilledLimitOrderEventAmountOut     = "AmountOut"
)

const (
	CancelLimitOrderEventKey           = "NewWithdraw"
	CancelLimitOrderEventCreator       = "Creator"
	CancelLimitOrderEventReceiver      = "Receiver"
	CancelLimitOrderEventToken0        = "Token0"
	CancelLimitOrderEventToken1        = "Token1"
	CancelLimitOrderEventTokenKey      = "TokenKey"
	CancelLimitOrderEventLimitOrderKey = "LimitOrderKey"
	CancelLimitOrderEventAmountOut     = "AmountOut"
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
