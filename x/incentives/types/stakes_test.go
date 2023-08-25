package types_test

// JCP TODO: fix? delete?
// import (
// 	"testing"
// 	time "time"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	dextypes "github.com/duality-labs/duality/x/dex/types"
// 	. "github.com/duality-labs/duality/x/incentives/types"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestStakesCoinsByQueryCondition(t *testing.T) {
// 	owner, err := sdk.AccAddressFromBech32("dual1lyaz7emmzreenas4fpz49a49958kye7wxuvsdr")
// 	require.NoError(t, err)

// 	allCoins := sdk.Coins{
// 		sdk.NewInt64Coin(
// 			dextypes.NewPoolDenom(
// 				&dextypes.PairID{
// 					Token0: "coin1",
// 					Token1: "coin2",
// 				},
// 				25,
// 				10,
// 			).String(),
// 			100,
// 		),
// 		sdk.NewInt64Coin(
// 			dextypes.NewPoolDenom(
// 				&dextypes.PairID{
// 					Token0: "coin1",
// 					Token1: "coin2",
// 				},
// 				75,
// 				10,
// 			).String(),
// 			200,
// 		),
// 		sdk.NewInt64Coin(
// 			dextypes.NewPoolDenom(
// 				&dextypes.PairID{
// 					Token0: "coin1",
// 					Token1: "coin2",
// 				},
// 				75,
// 				50,
// 			).String(),
// 			200,
// 		),
// 	}

// 	stakes := Stakes{
// 		NewStake(1, owner, sdk.Coins{allCoins[0]}, time.Time{}, 0),
// 		NewStake(2, owner, sdk.Coins{allCoins[1]}, time.Time{}, 0),
// 		NewStake(3, owner, sdk.Coins{allCoins[2]}, time.Time{}, 0),
// 	}

// 	pairID := &dextypes.PairID{
// 		Token0: "coin1",
// 		Token1: "coin2",
// 	}

// 	tests := []struct {
// 		name        string
// 		queryCond   QueryCondition
// 		coinsByCond sdk.Coins
// 	}{
// 		{
// 			name:        "All coins",
// 			queryCond:   QueryCondition{PairID: pairID, StartTick: 0, EndTick: 100},
// 			coinsByCond: sdk.Coins{allCoins[0], allCoins[1]},
// 		},
// 		{
// 			name:        "Coin1 only",
// 			queryCond:   QueryCondition{PairID: pairID, StartTick: 0, EndTick: 50},
// 			coinsByCond: sdk.Coins{allCoins[0]},
// 		},
// 		{
// 			name:        "Coin2 only",
// 			queryCond:   QueryCondition{PairID: pairID, StartTick: 50, EndTick: 100},
// 			coinsByCond: sdk.Coins{allCoins[1]},
// 		},
// 		{
// 			name:        "No coins",
// 			queryCond:   QueryCondition{PairID: pairID, StartTick: 100, EndTick: 200},
// 			coinsByCond: sdk.Coins{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			coins := stakes.CoinsByQueryCondition(tt.queryCond)
// 			assert.Equal(t, tt.coinsByCond, coins)
// 		})
// 	}
// }

// func TestStakesCoinsByQueryConditionMultiple(t *testing.T) {
// 	owner, err := sdk.AccAddressFromBech32("dual1lyaz7emmzreenas4fpz49a49958kye7wxuvsdr")
// 	require.NoError(t, err)

// 	allCoins := sdk.Coins{}

// 	coins1 := sdk.NewInt64Coin(
// 		dextypes.NewPoolDenom(
// 			&dextypes.PairID{
// 				Token0: "coin1",
// 				Token1: "coin2",
// 			},
// 			25,
// 			10,
// 		).String(),
// 		100,
// 	)
// 	allCoins = allCoins.Add(coins1)

// 	coins2 := sdk.NewInt64Coin(
// 		dextypes.NewPoolDenom(
// 			&dextypes.PairID{
// 				Token0: "coin1",
// 				Token1: "coin2",
// 			},
// 			75,
// 			10,
// 		).String(),
// 		200,
// 	)
// 	allCoins = allCoins.Add(coins2)

// 	coins3 := sdk.NewInt64Coin(
// 		dextypes.NewPoolDenom(
// 			&dextypes.PairID{
// 				Token0: "coin1",
// 				Token1: "coin2",
// 			},
// 			75,
// 			50,
// 		).String(),
// 		200,
// 	)
// 	allCoins = allCoins.Add(coins3)
// 	assert.Equal(t, allCoins, allCoins.Sort())

// 	stakes := Stakes{
// 		NewStake(1, owner, allCoins, time.Time{}, 0),
// 	}

// 	pairID := &dextypes.PairID{
// 		Token0: "coin1",
// 		Token1: "coin2",
// 	}

// 	tests := []struct {
// 		name      string
// 		queryCond QueryCondition
// 		expected  sdk.Coins
// 	}{
// 		{
// 			name:      "All coins",
// 			queryCond: QueryCondition{PairID: pairID, StartTick: 0, EndTick: 100},
// 			expected:  sdk.Coins{coins1, coins2},
// 		},
// 		{
// 			name:      "Coin1 only",
// 			queryCond: QueryCondition{PairID: pairID, StartTick: 0, EndTick: 50},
// 			expected:  sdk.Coins{coins1},
// 		},
// 		{
// 			name:      "Coin2 only",
// 			queryCond: QueryCondition{PairID: pairID, StartTick: 50, EndTick: 100},
// 			expected:  sdk.Coins{coins2},
// 		},
// 		{
// 			name:      "No coins",
// 			queryCond: QueryCondition{PairID: pairID, StartTick: 100, EndTick: 200},
// 			expected:  sdk.Coins{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			coins := stakes.CoinsByQueryCondition(tt.queryCond)
// 			assert.Equal(t, tt.expected, coins)
// 		})
// 	}
// }
