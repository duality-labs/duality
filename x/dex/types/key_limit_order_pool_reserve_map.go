package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LimitOrderPoolReserveMapKeyPrefix is the prefix to retrieve all LimitOrderPoolReserveMap
	LimitOrderPoolReserveMapKeyPrefix = "LimitOrderPoolReserveMap/value/"
)

// LimitOrderPoolReserveMapKey returns the store key to retrieve a LimitOrderPoolReserveMap from the index fields
func LimitOrderPoolReserveMapKey(
	count string,
) []byte {
	var key []byte

	countBytes := []byte(count)
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}
