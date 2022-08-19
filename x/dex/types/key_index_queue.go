package types

import (
	"encoding/binary"
	"strconv"
)

var _ binary.ByteOrder

const (
	// IndexQueueKeyPrefix is the prefix to retrieve all IndexQueue
	IndexQueueKeyPrefix = "IndexQueue/value/"
)

// IndexQueueKey returns the store key to retrieve a IndexQueue from the index fields
func IndexQueueKey(
	token0 string,
	token1 string,
	Index int32,

) []byte {
	var key []byte

	token0Bytes := []byte(token0)
	key = append(key, token0Bytes...)
	key = append(key, []byte("/")...)

	token1Bytes := []byte(token1)
	key = append(key, token1Bytes...)
	key = append(key, []byte("/")...)

	IndexBytes := []byte(strconv.Itoa(int(Index)))
	key = append(key, IndexBytes...)
	key = append(key, []byte("/")...)

	return key
}
