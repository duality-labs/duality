package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PairMapKeyPrefix is the prefix to retrieve all PairMap
	PairMapKeyPrefix = "PairMap/value/"
)

// PairMapKey returns the store key to retrieve a PairMap from the index fields
func PairMapKey(
	pairId string,
) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
