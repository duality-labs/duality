package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
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
		denom      string
		testResult bool
	}{
		{
			name:       "Matching denom and tick range",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			denom:      dexkeeper.NewDepositDenom(&dextypes.PairID{Token0: "coin1", Token1: "coin2"}, 15, 5).String(),
			testResult: true,
		},
		{
			name:       "Non-matching denom",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			denom:      dexkeeper.NewDepositDenom(&dextypes.PairID{Token0: "coin1", Token1: "coin3"}, 15, 5).String(),
			testResult: false,
		},
		{
			name:       "Non-matching tick range",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			denom:      dexkeeper.NewDepositDenom(&dextypes.PairID{Token0: "coin1", Token1: "coin3"}, 15, 6).String(),
			testResult: false,
		},
		{
			name:       "Non-matching tick fee range lower",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			denom:      dexkeeper.NewDepositDenom(&dextypes.PairID{Token0: "coin1", Token1: "coin3"}, 10, 5).String(),
			testResult: false,
		},
		{
			name:       "Non-matching tick fee range upper",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 30, EndTick: 40},
			denom:      dexkeeper.NewDepositDenom(&dextypes.PairID{Token0: "coin1", Token1: "coin3"}, 20, 5).String(),
			testResult: false,
		},
		{
			name:       "Invalid denom",
			queryCond:  QueryCondition{PairID: pairID, StartTick: 10, EndTick: 20},
			denom:      "invalid-denom",
			testResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.queryCond.Test(tt.denom)
			assert.Equal(t, tt.testResult, result)
		})
	}
}
