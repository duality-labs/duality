package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TickMapKeyPrefix is the prefix to retrieve all TickMap
	TickMapKeyPrefix = "TickMap/value/"
)

// TickMapKey returns the store key to retrieve a TickMap from the index fields
func TickMapKey(
	tickIndex string,
) []byte {
	var key []byte

	tickIndexBytes := []byte(tickIndex)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}
