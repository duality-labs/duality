package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// IncentivePlanKeyPrefix is the prefix to retrieve all IncentivePlan
	IncentivePlanKeyPrefix = "IncentivePlan/value/"
)

// IncentivePlanKey returns the store key to retrieve a IncentivePlan from the index fields
func IncentivePlanKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
