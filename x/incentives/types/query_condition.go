package types

import (
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

func (qc QueryCondition) Test(metadata *dextypes.PoolMetadata) bool {
	if *qc.PairID != *metadata.PairId {
		return false
	}

	lowerTick := metadata.NormalizedCenterTickIndex - int64(metadata.Fee)
	upperTick := metadata.NormalizedCenterTickIndex + int64(metadata.Fee)
	lowerTickQualifies := qc.StartTick <= lowerTick && lowerTick <= qc.EndTick
	upperTickQualifies := qc.StartTick <= upperTick && upperTick <= qc.EndTick
	return lowerTickQualifies && upperTickQualifies
}
