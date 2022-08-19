package types

import (
	"encoding/binary"
	"strconv"
)

var _ binary.ByteOrder

const (
	// VirtualPriceQueueKeyPrefix is the prefix to retrieve all IndexQueue
	IndexQueueKeyPrefix = "VirtualPriceQueue/value/"
)

// VirtualPriceQueueKey returns the store key to retrieve a IndexQueue from the index fields
func IndexQueueKey(
	Index int32,
) []byte {
	var key []byte

	IndexBytes := []byte(strconv.Itoa(int(Index)))
	key = append(key, IndexBytes...)
	key = append(key, []byte("/")...)

	return key
}
