package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualPriceQueueKeyPrefix is the prefix to retrieve all VirtualPriceQueue
	VirtualPriceQueueKeyPrefix = "VirtualPriceQueue/value/"
)

// VirtualPriceQueueKey returns the store key to retrieve a VirtualPriceQueue from the index fields
func VirtualPriceQueueKey(
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
