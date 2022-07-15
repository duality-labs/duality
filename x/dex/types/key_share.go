package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ShareKeyPrefix is the prefix to retrieve all Share
	ShareKeyPrefix = "Share/value/"
)

// ShareKey returns the store key to retrieve a Share from the index fields
func ShareKey(
	owner string,
	token0 string,
	token1 string,
	price string,
	fee string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	token0Bytes := []byte(token0)
	key = append(key, token0Bytes...)
	key = append(key, []byte("/")...)

	token1Bytes := []byte(token1)
	key = append(key, token1Bytes...)
	key = append(key, []byte("/")...)

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}
