package types

import (
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

const (
	DepositSharesPrefix = "DualityLPShares"
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

	// TickKeyPrefix is the prefix to retrieve all Tick
	BaseTickKeyPrefix = "Tick/value"

	// TradingPairKeyPrefix is the prefix to retrieve all TradingPair
	TradingPairKeyPrefix = "TradingPair/value"
)

func TickPrefix(pairId string) []byte {
	return append(KeyPrefix(BaseTickKeyPrefix), KeyPrefix(pairId)...)
}

// TokenMapKey returns the store key to retrieve a TokenMap from the index fields
func TokenMapKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickKey(pairId string, tickIndex int64) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := []byte(strconv.Itoa(int(tickIndex)))
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TradingPairKey(pairId string) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

// Limit Order Pool Mappings and Keys
const (
	BaseLimitOrderPrefix = "LimitOrderTranche/value"

	// LimitOrderTrancheUserSharesWithdrawnKeyPrefix is the prefix to retrieve all LimitOrderTrancheUserSharesWithdrawn
	LimitOrderTrancheUserSharesWithdrawnKeyPrefix = "LimitOrderTrancheUserSharesWithdrawn/value"

	// LimitOrderTrancheUserKeyPrefix is the prefix to retrieve all LimitOrderTrancheUser
	LimitOrderTrancheUserKeyPrefix = "LimitOrderTrancheUser/value"

	// LimitOrderTrancheKeyPrefix is the prefix to retrieve all LimitOrderTranche
	LimitOrderTrancheKeyPrefix = "LimitOrderTranche/value"

	// LimitOrderTrancheReserveMapKeyPrefix is the prefix to retrieve all LimitOrderTrancheReserveMap
	LimitOrderTrancheReserveMapKeyPrefix = "LimitOrderTrancheReserveMap/value"

	// LimitOrderTrancheFillMapKeyPrefix is the prefix to retrieve all LimitOrderTrancheFillMap
	LimitOrderTrancheFillMapKeyPrefix = "LimitOrderTrancheFillMap/value"
)

// LimitOrderTrancheUserSharesWithdrawnKey returns the store key to retrieve a LimitOrderTrancheUserSharesWithdrawn from the index fields
func LimitOrderTrancheUserSharesWithdrawnKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
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

// LimitOrderTrancheUserKey returns the store key to retrieve a LimitOrderTrancheUser from the index fields
func LimitOrderTrancheUserKey(pairId string, tickIndex int64, token string, count uint64, address string) []byte {
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

// LimitOrderTrancheKey returns the store key to retrieve a LimitOrderTranche from the index fields
func LimitOrderTrancheKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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

func LimitOrderTrancheReserveMapKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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

func LimitOrderTrancheFillMapKey(pairId string, tickIndex int64, token string, count uint64) []byte {
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

// DepositFailed Event Attributes
const (
	DepositFailEventKey          = "NewDeposit"
	DepositFailEventCreator      = "Creator"
	DepositFailEventToken0       = "Token0"
	DepositFailEventToken1       = "Token1"
	DepositFailEventPrice        = "TickIndex"
	DepositFailEventFeeIndex     = "FeeIndex"
	DepositFailEventReceiver     = "Receiver"
	DepositFailEventOldReserves0 = "OldReserves0"
	DepositFailEventOldReserves1 = "OldReserves1"
	DepositFailAmountToDeposit0  = "AmountToDeposit0"
	DepositFailAmountToDeposit1  = "AmountToDeposit1"
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
	WithdrawEventOldReserves0  = "OldReserves0"
	WithdrawEventOldReserves1  = "OldReserves1"
	WithdrawEventNewReserves0  = "NewReserves0"
	WithdrawEventNewReserves1  = "NewReserves1"
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
	FeeTierKey      = "FeeTier-value-"
	FeeTierCountKey = "FeeTier-count-"
)
