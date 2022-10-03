package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LimitOrderPoolUserShareMapKeyPrefix is the prefix to retrieve all LimitOrderPoolUserShareMap
	LimitOrderPoolUserShareMapKeyPrefix = "LimitOrderPoolUserShareMap/value/"
)

// LimitOrderPoolUserShareMapKey returns the store key to retrieve a LimitOrderPoolUserShareMap from the index fields
func LimitOrderPoolUserShareMapKey(
	count string,
	address string,
) []byte {
	var key []byte

	countBytes := []byte(count)
	key = append(key, countBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
