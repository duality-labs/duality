package types

import (
	"encoding/binary"
	"strconv"
)

var _ binary.ByteOrder

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

	// We don't want pair ABC<>DEF to have the same key as AB<>CDEF
	Separator = "/"
)

//Pair
// PairsPrefix Pairs/Value
// Key (token0, token1)

// Index Queue
// IndexQueuePrefix(token0, token1)
// => Pairs/Value/Token0|Token1/IndexQueue/value/
// Key (id int32)

// Ticks
// TicksPrefix(token0, token1)
// => Pairs/Value/Token0|Token1/Ticks/value/

//Key (price, fee, orderType)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	// NodesKeyPrefix is the prefix to retrieve all Nodes
	NodesKeyPrefix = "Nodes/value/"

	// TicksKeyPrefix is the prefix to retrieve all Ticks
	BaseTicksKeyPrefix = "Ticks/value/"
	// IndexQueueKeyPrefix is the prefix to retrieve all IndexQueue
	BaseIndexQueueKeyPrefix = "IndexQueue/value/"

	// PairsKeyPrefix is the prefix to retrieve all Pairs
	BasePairsKeyPrefix = "Pairs/value/"

	// SharesKeyPrefix is the prefix to retrieve all Shares
	BaseSharesKeyPrefix = "Shares/value/"
)

func PairsPrefix() []byte {
	return KeyPrefix(BasePairsKeyPrefix)
}

func PairPrefixHelper(token0, token1 string) []byte {
	return append(KeyPrefix(token0), append(KeyPrefix(Separator), append(KeyPrefix(token1), KeyPrefix(Separator)...)...)...)
}

func IndexQueuePrefix(token0 string, token1 string) []byte {
	return append(KeyPrefix(BaseIndexQueueKeyPrefix), PairPrefixHelper(token0, token1)...)
}

func TicksPrefix(token0 string, token1 string) []byte {
	return append(KeyPrefix(BaseTicksKeyPrefix), PairPrefixHelper(token0, token1)...)
}

func SharesPrefix(token0 string, token1 string) []byte {
	return append(KeyPrefix(BaseSharesKeyPrefix), PairPrefixHelper(token0, token1)...)
}

// SharesKey returns the store key to retrieve a Shares from the index fields
func SharesKey(address string, price string, fee string, orderType string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// NodesKey returns the store key to retrieve a Nodes from the index fields
func NodesKey(node string) []byte {
	var key []byte

	nodeBytes := []byte(node)
	key = append(key, nodeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// TicksKey returns the store key to retrieve a Ticks from the index fields
func TicksKey(price string, fee string, orderType string) []byte {
	var key []byte

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// IndexQueueKey returns the store key to retrieve a IndexQueue from the index fields
func IndexQueueKey(Index int32) []byte {
	var key []byte

	IndexBytes := []byte(strconv.Itoa(int(Index)))
	key = append(key, IndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PairsKey returns the store key to retrieve a Pairs from the index fields
func PairsKey(token0 string, token1 string) []byte {
	var key []byte

	token0Bytes := []byte(token0)
	key = append(key, token0Bytes...)
	key = append(key, []byte("/")...)

	token1Bytes := []byte(token1)
	key = append(key, token1Bytes...)
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
	DepositTokenDirection    = "TokenDirection"
	DepositEventOldReserves  = "OldReserves"
	DepositEventNewReserves  = "NewReserves"
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
