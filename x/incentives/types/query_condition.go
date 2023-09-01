package types

import (
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

func (qc QueryCondition) Test(poolMetadata dextypes.PoolMetadata) bool {
	if !poolMetadata.PairID.Equal(qc.PairID) {
		return false
	}

	lowerTick := poolMetadata.Tick - int64(poolMetadata.Fee)
	upperTick := poolMetadata.Tick + int64(poolMetadata.Fee)
	lowerTickQualifies := qc.StartTick <= lowerTick && lowerTick <= qc.EndTick
	upperTickQualifies := qc.StartTick <= upperTick && upperTick <= qc.EndTick

	return lowerTickQualifies && upperTickQualifies
}
