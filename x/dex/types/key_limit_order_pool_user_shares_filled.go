package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LimitOrderPoolUserSharesFilledKeyPrefix is the prefix to retrieve all LimitOrderPoolUserSharesFilled
	LimitOrderPoolUserSharesFilledKeyPrefix = "LimitOrderPoolUserSharesFilled/value/"
)

// LimitOrderPoolUserSharesFilledKey returns the store key to retrieve a LimitOrderPoolUserSharesFilled from the index fields
func LimitOrderPoolUserSharesFilledKey(
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
