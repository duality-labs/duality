package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	dextypes "github.com/duality-labs/duality/x/dex/types"
	. "github.com/duality-labs/duality/x/incentives/types"
)

func TestQueryCondition(t *testing.T) {
	pairID := &dextypes.PairID{
		Token0: "coin1",
		Token1: "coin2",
	}

	tests := []struct {
		name       string
		queryCond  QueryCondition
		poolParams dextypes.PoolParams
		testResult bool
	}{
		{
			name:       "Matching denom and tick range",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			poolParams: dextypes.NewPoolParams(pairID, 15, 5),
			testResult: true,
		},
		{
			name:       "Non-matching denom",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			poolParams: dextypes.NewPoolParams(&dextypes.PairID{Token0: "coin1", Token1: "coin3"}, 15, 5),
			testResult: false,
		},
		{
			name:       "Non-matching tick range",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			poolParams: dextypes.NewPoolParams(pairID, 15, 6),
			testResult: false,
		},
		{
			name:       "Non-matching tick fee range lower",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			poolParams: dextypes.NewPoolParams(pairID, 10, 5),
			testResult: false,
		},
		{
			name:       "Non-matching tick fee range upper",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			poolParams: dextypes.NewPoolParams(pairID, 20, 5),
			testResult: false,
		},
		{
			name:       "Invalid denom",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			poolParams: dextypes.NewPoolParams(&dextypes.PairID{Token0: "coinz", Token1: "coinz"}, 15, 5),
			testResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.queryCond.Test(tt.poolParams)
			assert.Equal(t, tt.testResult, result)
		})
	}
}
