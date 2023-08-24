package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolParams struct {
	PairID *PairID
	Tick   int64
	Fee    uint64
}

func ParsePoolRefToParams(poolRef []byte) (PoolParams, error) {
	parts := bytes.Split(poolRef, []byte("/"))
	if len(parts) != 4 {
		return PoolParams{}, ErrInvalidDepositDenom
	}

	pairID, err := NewPairIDFromCanonicalString(string(parts[0]))
	if err != nil {
		return PoolParams{}, err
	}

	tick, err := BytesToTickIndex(parts[1])
	if err != nil {
		return PoolParams{}, err
	}

	fee := sdk.BigEndianToUint64(parts[2])

	return PoolParams{PairID: pairID, Tick: tick, Fee: fee}, nil
}

func NewPoolParams(pairID *PairID, tick int64, fee uint64) PoolParams {
	return PoolParams{PairID: pairID, Tick: tick, Fee: fee}
}
