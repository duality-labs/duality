package types

import (
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

func (qc QueryCondition) Test(poolParams dextypes.PoolParams) bool {
	lowerTick := poolParams.Tick - int64(poolParams.Fee)
	upperTick := poolParams.Tick + int64(poolParams.Fee)
	lowerTickQualifies := qc.StartTick <= lowerTick && lowerTick <= qc.EndTick
	upperTickQualifies := qc.StartTick <= upperTick && upperTick <= qc.EndTick
	return lowerTickQualifies && upperTickQualifies
}
