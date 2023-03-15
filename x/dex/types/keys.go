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

const (
	TokensKey      = "Tokens-value-"
	TokensCountKey = "Tokens-count-"

	// TokenMapKeyPrefix is the prefix to retrieve all TokenMap
	TokenMapKeyPrefix = "TokenMap/value"

	// TickKeyPrefix is the prefix to retrieve all Tick
	BaseTickKeyPrefix = "Tick/value"

	// TickLiquidityKeyPrefix is the prefix to retrieve all TickLiquidity
	TickLiquidityKeyPrefix = "TickLiquidity/value/"
)

func TickPrefix(pairId *PairId) []byte {
	return append(KeyPrefix(BaseTickKeyPrefix), KeyPrefix(pairId.Stringify())...)
}

// TokenMapKey returns the store key to retrieve a TokenMap from the index fields
func TokenMapKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickIndexToBytes(tickIndex int64, pairId *PairId, tokenIn string) []byte {
	// NOTE: We flip the sign on ticks storing token0 so that all liquidity is index left to right.
	// This allows us to iterate through liquidity consistently regardless of 0to1 vs 1to0
	if pairId.Token0 == tokenIn {
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

// Limit Order Pool Mappings and Keys
const (
	BaseLimitOrderPrefix = "LimitOrderTranche/value"

	// LimitOrderTrancheUserKeyPrefix is the prefix to retrieve all LimitOrderTrancheUser
	LimitOrderTrancheUserKeyPrefix = "LimitOrderTrancheUser/value"

	// LimitOrderTrancheKeyPrefix is the prefix to retrieve all LimitOrderTranche
	LimitOrderTrancheKeyPrefix = "LimitOrderTranche/value"

	FilledLimitOrderTrancheKeyPrefix = "FilledLimitOrderTranche/value/"

	// GoodTilRecordKeyPrefix is the prefix to retrieve all GoodTilRecord
	GoodTilRecordKeyPrefix = "GoodTilRecord/value/"
)

// LimitOrderTrancheUserKey returns the store key to retrieve a LimitOrderTrancheUser from the index fields
func LimitOrderTrancheUserKey(pairId *PairId, tickIndex int64, token string, trancheKey string, address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	pairIdBytes := []byte(pairId.Stringify())
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairId, token)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
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

// FilledLimitOrderTrancheKey returns the store key to retrieve a FilledLimitOrderTranche from the index fields
func FilledLimitOrderTrancheKey(
	pairId *PairId,
	tokenIn string,
	tickIndex int64,
	trancheKey string,
) []byte {
	var key []byte

	pairIdBytes := []byte(pairId.Stringify())
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairId, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	trancheKeyBytes := []byte(trancheKey)
	key = append(key, trancheKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}

func FilledLimitOrderTranchePrefix(
	pairId *PairId,
	tokenIn string,
	tickIndex int64,
) []byte {
	var key []byte = KeyPrefix(FilledLimitOrderTrancheKeyPrefix)

	pairIdBytes := []byte(pairId.Stringify())
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairId, tokenIn)
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
	pairId *PairId,
	tokenIn string,
	tickIndex int64,
	liquidityType string,
	liquidityIndex interface{},
) []byte {
	var key []byte

	pairIdBytes := []byte(pairId.Stringify())
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairId, tokenIn)
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
	pairId *PairId,
	tokenIn string,
	tickIndex int64,
) []byte {
	var key []byte = KeyPrefix(TickLiquidityKeyPrefix)

	pairIdBytes := []byte(pairId.Stringify())
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	tokenInBytes := []byte(tokenIn)
	key = append(key, tokenInBytes...)
	key = append(key, []byte("/")...)

	tickIndexBytes := TickIndexToBytes(tickIndex, pairId, tokenIn)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	liquidityTypeBytes := []byte(LiquidityTypeLimitOrder)
	key = append(key, liquidityTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickLiquidityPrefix(pairId *PairId, tokenIn string) []byte {
	var key []byte
	key = append(KeyPrefix(TickLiquidityKeyPrefix), KeyPrefix(pairId.Stringify())...)
	key = append(key, KeyPrefix(tokenIn)...)
	return key
}

func GoodTilRecordKey(
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
	SwapEventTokenIn  = "TokenIn"
	SwapEventTokenOut = "TokenOut"
	SwapEventAmountIn = "AmountIn"
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
	PlaceLimitOrderEventTrancheKey = "TrancheKey"
)

const (
	WithdrawFilledLimitOrderEventKey           = "NewWithdraw"
	WithdrawFilledLimitOrderEventCreator       = "Creator"
	WithdrawFilledLimitOrderEventToken0        = "Token0"
	WithdrawFilledLimitOrderEventToken1        = "Token1"
	WithdrawFilledLimitOrderEventTokenKey      = "TokenKey"
	WithdrawFilledLimitOrderEventLimitOrderKey = "LimitOrderKey"
	WithdrawFilledLimitOrderEventAmountOut     = "AmountOut"
)

const (
	GoodTilPurgeHitGasLimitEventKey = "GoodTilPurgeHitGasLimit"
	GoodTilPurgeHitGasLimitEventGas = "Gas"
)

const (
	CancelLimitOrderEventKey           = "NewWithdraw"
	CancelLimitOrderEventCreator       = "Creator"
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

const (
	// NOTE: have to add letter so that LP deposits are indexed ahead of LimitOrders
	LiquidityTypePoolReserves = "A_PoolDeposit"
	LiquidityTypeLimitOrder   = "B_LODeposit"
)

var (
	JITGoodTilTime = time.Time{}
)

const (
	// TODO: jcp figure out a good number here
	GoodTilPurgeGasBuffer = 1000
)
