package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TicksKeyPrefix is the prefix to retrieve all Ticks
	TicksKeyPrefix = "Ticks/value/"
)

// TicksKey returns the store key to retrieve a Ticks from the index fields
func TicksKey(
	price string,
	fee string,
	direction string,
	orderType string,
) []byte {
	var key []byte

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	directionBytes := []byte(direction)
	key = append(key, directionBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}
