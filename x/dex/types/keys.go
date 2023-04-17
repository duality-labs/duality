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
	DepositEventKey          = "NewDeposit"
	DepositEventCreator      = "Creator"
	DepositEventToken0       = "Token0"
	DepositEventToken1       = "Token1"
	DepositEventPrice        = "TickIndex"
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
	WithdrawEventPrice         = "TickIndex"
	WithdrawEventFee           = "Fee"
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
	SwapEventTokenIn  = "TokenIn"
	SwapEventTokenOut = "TokenOut"
	SwapEventAmountIn = "AmountIn"
	SwapEventAmoutOut = "AmountOut"
)

const (
	MultihopSwapEventKey      = "NewMultihopSwap"
	MultihopSwapEventCreator  = "Creator"
	MultihopSwapEventReceiver = "Receiver"
	MultihopSwapEventCoinIn   = "CoinIn"
	MultihopSwapEventCoinOut  = "CoinOut"
	MultihopSwapEventRoute    = "Route"
)

const (
	PlaceLimitOrderEventKey        = "NewPlaceLimitOrder"
	PlaceLimitOrderEventCreator    = "Creator"
	PlaceLimitOrderEventReceiver   = "Receiver"
	PlaceLimitOrderEventTokenIn    = "TokenIn"
	PlaceLimitOrderEventTokenOut   = "TokenOut"
	PlaceLimitOrderEventAmountIn   = "AmountIn"
	PlaceLimitOrderEventShares     = "Shares"
	PlaceLimitOrderEventTrancheKey = "TrancheKey"
)

const (
	WithdrawFilledLimitOrderEventKey        = "NewWithdraw"
	WithdrawFilledLimitOrderEventCreator    = "Creator"
	WithdrawFilledLimitOrderEventTokenIn    = "TokenIn"
	WithdrawFilledLimitOrderEventTokenOut   = "TokenOut"
	WithdrawFilledLimitOrderEventTrancheKey = "TrancheKey"
	WithdrawFilledLimitOrderEventAmountOut  = "AmountOut"
)

const (
	GoodTilPurgeHitGasLimitEventKey = "GoodTilPurgeHitGasLimit"
	GoodTilPurgeHitGasLimitEventGas = "Gas"
)

const (
	CancelLimitOrderEventKey        = "NewWithdraw"
	CancelLimitOrderEventCreator    = "Creator"
	CancelLimitOrderEventTokenIn    = "TokenIn"
	CancelLimitOrderEventTokenOut   = "TokenOut"
	CancelLimitOrderEventTrancheKey = "TrancheKey"
	CancelLimitOrderEventAmountOut  = "AmountOut"
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
