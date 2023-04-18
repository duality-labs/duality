package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/utils"
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
	DepositSharesPrefix = "DualityPoolShares"
)

func KeyPrefix(p string) []byte {
	key := []byte(p)
	key = append(key, []byte("/")...)
	return key
}

func TickIndexToBytes(tickIndex int64, pairID *PairID, tokenIn string) []byte {
	// NOTE: We flip the sign on ticks storing token0 so that all liquidity is index left to right.
	// This allows us to iterate through liquidity consistently regardless of 0to1 vs 1to0
	if pairID.Token0 == tokenIn {
		tickIndex *= -1
	}
	key := make([]byte, 9)
	if tickIndex < 0 {
		copy(key[1:], sdk.Uint64ToBigEndian(uint64(tickIndex)))
	} else {
		copy(key[:1], []byte{0x01})
		copy(key[1:], sdk.Uint64ToBigEndian(uint64(tickIndex)))
	}

	return key
}

const (
	// TickLiquidityKeyPrefix is the prefix to retrieve all TickLiquidity
	TickLiquidityKeyPrefix = "TickLiquidity/value/"

	// LimitOrderTrancheUserKeyPrefix is the prefix to retrieve all LimitOrderTrancheUser
	LimitOrderTrancheUserKeyPrefix = "LimitOrderTrancheUser/value"

	// LimitOrderTrancheKeyPrefix is the prefix to retrieve all LimitOrderTranche
	LimitOrderTrancheKeyPrefix = "LimitOrderTranche/value"

	// InactiveLimitOrderTrancheKeyPrefix is the prefix to retrieve all InactiveLimitOrderTranche
	InactiveLimitOrderTrancheKeyPrefix = "InactiveLimitOrderTranche/value/"

	// LimitOrderExpirationKeyPrefix is the prefix to retrieve all LimitOrderExpiration
	LimitOrderExpirationKeyPrefix = "LimitOrderExpiration/value/"
)

// LimitOrderTrancheUserKey returns the store key to retrieve a LimitOrderTrancheUser from the index fields
func LimitOrderTrancheUserKey(address, trancheKey string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	trancheKeyBytes := []byte(trancheKey)
	key = append(key, trancheKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}

func LimitOrderTrancheUserAddressPrefix(address string) []byte {
	key := KeyPrefix(LimitOrderTrancheUserKeyPrefix)
	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// InactiveLimitOrderTrancheKey returns the store key to retrieve a InactiveLimitOrderTranche from the index fields
func InactiveLimitOrderTrancheKey(
	pairID *PairID,
	tokenIn string,
	tickIndex int64,
	trancheKey string,
) []byte {
	var key []byte

	pairIDBytes := []byte(pairID.Stringify())
	key = append(key, pairIDBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairID, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	trancheKeyBytes := []byte(trancheKey)
	key = append(key, trancheKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}

func InactiveLimitOrderTranchePrefix(
	pairID *PairID,
	tokenIn string,
	tickIndex int64,
) []byte {
	key := KeyPrefix(InactiveLimitOrderTrancheKeyPrefix)

	pairIDBytes := []byte(pairID.Stringify())
	key = append(key, pairIDBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairID, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func LiquidityIndexBytes(liquidityIndex interface{}) []byte {
	switch index := liquidityIndex.(type) {
	case uint64:
		liquidityIndexBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(liquidityIndexBytes, index)
		return liquidityIndexBytes
	case string:
		return []byte(index)
	default:
		panic("LiquidityIndex is not a valid type")
	}
}

func TimeBytes(timestamp time.Time) []byte {
	unixMs := uint64(timestamp.UnixMilli())
	str := utils.Uint64ToSortableString(unixMs)
	return []byte(str)
}

func TickLiquidityKey(
	pairID *PairID,
	tokenIn string,
	tickIndex int64,
	liquidityType string,
	liquidityIndex interface{},
) []byte {
	var key []byte

	pairIDBytes := []byte(pairID.Stringify())
	key = append(key, pairIDBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairID, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	liquidityTypeBytes := []byte(liquidityType)
	key = append(key, liquidityTypeBytes...)
	key = append(key, []byte("/")...)

	key = append(key, LiquidityIndexBytes(liquidityIndex)...)
	key = append(key, []byte("/")...)

	return key
}

func TickLiquidityLimitOrderPrefix(
	pairID *PairID,
	tokenIn string,
	tickIndex int64,
) []byte {
	key := KeyPrefix(TickLiquidityKeyPrefix)

	pairIDBytes := []byte(pairID.Stringify())
	key = append(key, pairIDBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairID, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	liquidityTypeBytes := []byte(LiquidityTypeLimitOrder)
	key = append(key, liquidityTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickLiquidityPrefix(pairID *PairID, tokenIn string) []byte {
	var key []byte
	key = append(KeyPrefix(TickLiquidityKeyPrefix), KeyPrefix(pairID.Stringify())...)
	key = append(key, KeyPrefix(tokenIn)...)

	return key
}

func LimitOrderExpirationKey(
	goodTilDate time.Time,
	trancheRef []byte,
) []byte {
	var key []byte

	goodTilDateBytes := TimeBytes(goodTilDate)
	key = append(key, goodTilDateBytes...)
	key = append(key, []byte("/")...)

	key = append(key, trancheRef...)
	key = append(key, []byte("/")...)

	return key
}

// Deposit Event Attributes
const (
	DepositEventKey                = "Deposit"
	DepositEventCreator            = "Creator"
	DepositEventToken0             = "Token0"
	DepositEventToken1             = "Token1"
	DepositEventPrice              = "TickIndex"
	DepositEventFee                = "Fee"
	DepositEventReceiver           = "Receiver"
	DepositEventReserves0Deposited = "Reserves0Deposited"
	DepositEventReserves1Deposited = "Reserves1Deposited"
	DepositEventSharesMinted       = "SharesMinted"
)

// Withdraw Event Attributes
const (
	WithdrawEventKey                = "Withdraw"
	WithdrawEventCreator            = "Creator"
	WithdrawEventToken0             = "Token0"
	WithdrawEventToken1             = "Token1"
	WithdrawEventPrice              = "TickIndex"
	WithdrawEventFee                = "Fee"
	WithdrawEventReceiver           = "Receiver"
	WithdrawEventReserves0Withdrawn = "Reserves0Withdrawn"
	WithdrawEventReserves1Withdrawn = "Reserves1Withdrawn"
	WithdrawEventSharesRemoved      = "SharesRemoved"
)

// Swap Event Attributes
const (
	SwapEventKey      = "Swap"
	SwapEventCreator  = "Creator"
	SwapEventReceiver = "Receiver"
	SwapEventToken0   = "Token0"
	SwapEventToken1   = "Token1"
	SwapEventTokenIn  = "TokenIn"
	SwapEventTokenOut = "TokenOut"
	SwapEventAmountIn = "AmountIn"
	SwapEventAmoutOut = "AmountOut"
)

// Multihop-Swap Event Attributes
const (
	MultihopSwapEventKey       = "MultihopSwap"
	MultihopSwapEventCreator   = "Creator"
	MultihopSwapEventReceiver  = "Receiver"
	MultihopSwapEventTokenIn   = "TokenIn"
	MultihopSwapEventTokenOut  = "TokenOut"
	MultihopSwapEventAmountIn  = "AmountIn"
	MultihopSwapEventAmountOut = "AmountOut"
	MultihopSwapEventRoute     = "Route"
)

// Place LimitOrder Event Attributes
const (
	PlaceLimitOrderEventKey        = "PlaceLimitOrder"
	PlaceLimitOrderEventCreator    = "Creator"
	PlaceLimitOrderEventReceiver   = "Receiver"
	PlaceLimitOrderEventToken0     = "Token0"
	PlaceLimitOrderEventToken1     = "Token1"
	PlaceLimitOrderEventTokenIn    = "TokenIn"
	PlaceLimitOrderEventTokenOut   = "TokenOut"
	PlaceLimitOrderEventAmountIn   = "AmountIn"
	PlaceLimitOrderEventLimitPrice = "LimitPrice"
	PlaceLimitOrderEventOrderType  = "OrderType"
	PlaceLimitOrderEventShares     = "Shares"
	PlaceLimitOrderEventTrancheKey = "TrancheKey"
)

// Withdraw LimitOrder Event Attributes
const (
	WithdrawFilledLimitOrderEventKey        = "Withdraw"
	WithdrawFilledLimitOrderEventCreator    = "Creator"
	WithdrawFilledLimitOrderEventToken0     = "Token0"
	WithdrawFilledLimitOrderEventToken1     = "Token1"
	WithdrawFilledLimitOrderEventTokenIn    = "TokenIn"
	WithdrawFilledLimitOrderEventTokenOut   = "TokenOut"
	WithdrawFilledLimitOrderEventTrancheKey = "TrancheKey"
	WithdrawFilledLimitOrderEventAmountOut  = "AmountOut"
)

// Cancel LimitOrder Event Attributes
const (
	CancelLimitOrderEventKey        = "Withdraw"
	CancelLimitOrderEventCreator    = "Creator"
	CancelLimitOrderEventToken0     = "Token0"
	CancelLimitOrderEventToken1     = "Token1"
	CancelLimitOrderEventTokenIn    = "TokenIn"
	CancelLimitOrderEventTokenOut   = "TokenOut"
	CancelLimitOrderEventTrancheKey = "TrancheKey"
	CancelLimitOrderEventAmountOut  = "AmountOut"
)

// Tick Update Event Attributes
const (
	EventTypeTickUpdate       = "TickUpdate"
	TickUpdateEventKey        = "TickUpdate"
	TickUpdateEventToken0     = "Token0"
	TickUpdateEventToken1     = "Token1"
	TickUpdateEventTokenIn    = "TokenIn"
	TickUpdateEventTickIndex  = "TickIndex"
	TickUpdateEventFee        = "Fee"
	TickUpdateEventTrancheKey = "TrancheKey"
	TickUpdateEventReserves   = "Reserves"
)

const (
	GoodTilPurgeHitGasLimitEventKey = "GoodTilPurgeHitGasLimit"
	GoodTilPurgeHitGasLimitEventGas = "Gas"
)

const (
	// NOTE: have to add letter so that LP deposits are indexed ahead of LimitOrders
	LiquidityTypePoolReserves = "A_PoolDeposit"
	LiquidityTypeLimitOrder   = "B_LODeposit"
)

func JITGoodTilTime() time.Time {
	return time.Time{}
}

const (
	// NOTE: This number is based current cost of all operations in EndBlock,
	// if that changes this value must be updated to ensure there is enough
	// remaining gas (weak proxy for timeoutPrepareProposal) to complete endBlock
	GoodTilPurgeGasBuffer = 50_000
	ExpiringLimitOrderGas = 10_000
)
