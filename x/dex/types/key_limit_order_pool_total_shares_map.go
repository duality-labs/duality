package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LimitOrderPoolTotalSharesMapKeyPrefix is the prefix to retrieve all LimitOrderPoolTotalSharesMap
	LimitOrderPoolTotalSharesMapKeyPrefix = "LimitOrderPoolTotalSharesMap/value/"
)

// LimitOrderPoolTotalSharesMapKey returns the store key to retrieve a LimitOrderPoolTotalSharesMap from the index fields
func LimitOrderPoolTotalSharesMapKey(
	count string,
) []byte {
	var key []byte

	countBytes := []byte(count)
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}
