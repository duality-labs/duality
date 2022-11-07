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
	key := []byte(p)
	key = append(key, []byte("/")...)
	return key
}

const (
	TokensKey      = "Tokens-value-"
	TokensCountKey = "Tokens-count-"

	// TokenObjectKeyPrefix is the prefix to retrieve all TokenObject
	TokenObjectKeyPrefix = "TokenObject/value"

	// TickObjectKeyPrefix is the prefix to retrieve all TickObject
	BaseTickObjectKeyPrefix = "TickObject/value"

	// PairObjectKeyPrefix is the prefix to retrieve all PairObject
	PairObjectKeyPrefix = "PairObject/value"

	// SharesKeyPrefix is the prefix to retrieve all Shares
	SharesKeyPrefix = "Shares/value"
)

func TickPrefix(pairId string) []byte {
	return append(KeyPrefix(BaseTickObjectKeyPrefix), KeyPrefix(pairId)...)
}

// TokenObjectKey returns the store key to retrieve a TokenObject from the index fields
func TokenObjectKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickObjectKey(pairId string, tickIndex int64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func PairObjectKey(pairId string) []byte {
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

	tickIndexBytes := []byte(string(tickIndex))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(string(feeIndex))
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// Limit Order Pool Mappings and Keys
const (
	BaseLimitOrderPrefix = "LimitOrderPool/value"

	// LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix is the prefix to retrieve all LimitOrderPoolUserSharesWithdrawnObject
	LimitOrderPoolUserSharesWithdrawnObjectKeyPrefix = "LimitOrderPoolUserSharesWithdrawnObject/value"

	// LimitOrderPoolUserShareObjectKeyPrefix is the prefix to retrieve all LimitOrderPoolUserShareObject
	LimitOrderPoolUserShareObjectKeyPrefix = "LimitOrderPoolUserShareObject/value"

	// LimitOrderPoolTotalSharesObjectKeyPrefix is the prefix to retrieve all LimitOrderPoolTotalSharesObject
	LimitOrderPoolTotalSharesObjectKeyPrefix = "LimitOrderPoolTotalSharesObject/value"

	// LimitOrderPoolReserveObjectKeyPrefix is the prefix to retrieve all LimitOrderPoolReserveObject
	LimitOrderPoolReserveObjectKeyPrefix = "LimitOrderPoolReserveObject/value"

	// LimitOrderPoolFillObjectKeyPrefix is the prefix to retrieve all LimitOrderPoolFillObject
	LimitOrderPoolFillObjectKeyPrefix = "LimitOrderPoolFillObject/value"
)

// LimitOrderPoolUserSharesWithdrawnObjectKey returns the store key to retrieve a LimitOrderPoolUserSharesWithdrawnObject from the index fields
func LimitOrderPoolUserSharesWithdrawnObjectKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
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

// LimitOrderPoolUserShareObjectKey returns the store key to retrieve a LimitOrderPoolUserShareObject from the index fields
func LimitOrderPoolUserShareObjectKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
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

// LimitOrderPoolTotalSharesObjectKey returns the store key to retrieve a LimitOrderPoolTotalSharesObject from the index fields
func LimitOrderPoolTotalSharesObjectKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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

func LimitOrderPoolReserveObjectKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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

func LimitOrderPoolFillObjectKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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
