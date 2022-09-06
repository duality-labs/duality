package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenMapKeyPrefix is the prefix to retrieve all TokenMap
	TokenMapKeyPrefix = "TokenMap/value/"
)

// TokenMapKey returns the store key to retrieve a TokenMap from the index fields
func TokenMapKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
