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
	pairId string,
	priceIndex string,
	fee string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	priceIndexBytes := []byte(priceIndex)
	key = append(key, priceIndexBytes...)
	key = append(key, []byte("/")...)

	feeBytes := []byte(fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}
