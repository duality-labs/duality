package types_test

import (
	"testing"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	. "github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocksCoinsByQueryCondition(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1xv9tklw7d82sezh9haa573wufgy59vmwe6xxe5")
	require.NoError(t, err)

	allCoins := sdk.Coins{
		sdk.NewInt64Coin(
			dexkeeper.NewDepositDenom(
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				25,
				10,
			).String(),
			100,
		),
		sdk.NewInt64Coin(
			dexkeeper.NewDepositDenom(
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				75,
				10,
			).String(),
			200,
		),
		sdk.NewInt64Coin(
			dexkeeper.NewDepositDenom(
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				75,
				50,
			).String(),
			200,
		),
	}

	duration := time.Duration(1 * time.Hour)
	endTime := time.Time{}

	locks := Locks{
		NewLock(1, owner, duration, endTime, sdk.Coins{allCoins[0]}),
		NewLock(2, owner, duration, endTime, sdk.Coins{allCoins[1]}),
		NewLock(3, owner, duration, endTime, sdk.Coins{allCoins[2]}),
	}

	pairID := &dextypes.PairID{
		Token0: "coin1",
		Token1: "coin2",
	}

	tests := []struct {
		name        string
		queryCond   QueryCondition
		coinsByCond sdk.Coins
	}{
		{
			name:        "All coins",
			queryCond:   QueryCondition{PairID: pairID, StartTick: 0, EndTick: 100},
			coinsByCond: sdk.Coins{allCoins[0], allCoins[1]},
		},
		{
			name:        "Coin1 only",
			queryCond:   QueryCondition{PairID: pairID, StartTick: 0, EndTick: 50},
			coinsByCond: sdk.Coins{allCoins[0]},
		},
		{
			name:        "Coin2 only",
			queryCond:   QueryCondition{PairID: pairID, StartTick: 50, EndTick: 100},
			coinsByCond: sdk.Coins{allCoins[1]},
		},
		{
			name:        "No coins",
			queryCond:   QueryCondition{PairID: pairID, StartTick: 100, EndTick: 200},
			coinsByCond: sdk.Coins{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins := locks.CoinsByQueryCondition(tt.queryCond)
			assert.Equal(t, tt.coinsByCond, coins)
		})
	}
}
