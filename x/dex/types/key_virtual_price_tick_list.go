package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualPriceTickListKeyPrefix is the prefix to retrieve all VirtualPriceTickList
	VirtualPriceTickListKeyPrefix = "VirtualPriceTickList/value/"
)

// VirtualPriceTickListKey returns the store key to retrieve a VirtualPriceTickList from the index fields
func VirtualPriceTickListKey(
	vPrice string,
	direction string,
	orderType string,
) []byte {
	var key []byte

	vPriceBytes := []byte(vPrice)
	key = append(key, vPriceBytes...)
	key = append(key, []byte("/")...)

	directionBytes := []byte(direction)
	key = append(key, directionBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}
