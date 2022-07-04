package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TickKeyPrefix is the prefix to retrieve all Tick
	TickKeyPrefix = "Tick/value/"
)

// TickKey returns the store key to retrieve a Tick from the index fields
func TickKey(
	token0 string,
	token1 string,
	price string,
	fee uint64,
) []byte {
	var key []byte

	token0Bytes := []byte(token0)
	key = append(key, token0Bytes...)
	key = append(key, []byte("/")...)

	token1Bytes := []byte(token1)
	key = append(key, token1Bytes...)
	key = append(key, []byte("/")...)

	priceBytes := []byte(price)
	key = append(key, priceBytes...)
	key = append(key, []byte("/")...)

	feeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(feeBytes, fee)
	key = append(key, feeBytes...)
	key = append(key, []byte("/")...)

	return key
}
