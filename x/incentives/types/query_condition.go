package types

import (
	"strings"

	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
)

func (qc QueryCondition) Test(denom string) bool {
	denomPrefix := dexkeeper.DepositDenomPairIDPrefix(qc.PairID.Token0, qc.PairID.Token1)
	if !strings.Contains(denom, denomPrefix) {
		return false
	}

	depositDenom, err := dexkeeper.NewDepositDenomFromString(denom)
	if err != nil {
		return false
	}

	lowerTick := depositDenom.Tick - int64(depositDenom.Fee)
	upperTick := depositDenom.Tick + int64(depositDenom.Fee)
	lowerTickQualifies := qc.StartTick <= lowerTick && lowerTick <= qc.EndTick
	upperTickQualifies := qc.StartTick <= upperTick && upperTick <= qc.EndTick
	return lowerTickQualifies && upperTickQualifies
}
