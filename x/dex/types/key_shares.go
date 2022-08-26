package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SharesKeyPrefix is the prefix to retrieve all Shares
	SharesKeyPrefix = "Shares/value/"
)

// SharesKey returns the store key to retrieve a Shares from the index fields
func SharesKey(
	address string,
	price string,
	fee string,
	orderType string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}
