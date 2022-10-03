package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LimitOrderPoolFillMapKeyPrefix is the prefix to retrieve all LimitOrderPoolFillMap
	LimitOrderPoolFillMapKeyPrefix = "LimitOrderPoolFillMap/value/"
)

// LimitOrderPoolFillMapKey returns the store key to retrieve a LimitOrderPoolFillMap from the index fields
func LimitOrderPoolFillMapKey(
	count string,
) []byte {
	var key []byte

	countBytes := []byte(count)
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	return key
}
