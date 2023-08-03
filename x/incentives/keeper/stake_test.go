package keeper_test

import (
	"testing"
	"time"

	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = suite.TestingSuite(nil)

func (suite *KeeperTestSuite) TestStakeLifecycle() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpecs: []depositSpec{
			{
				addr:   addr0,
				token0: sdk.NewInt64Coin("TokenA", 10),
				token1: sdk.NewInt64Coin("TokenB", 10),
				tick:   0,
				fee:    1,
			},
		},
		stakeDistEpochOffset: -2,
	})

	retrievedStake, err := suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)

	// unstake the full amount
	suite.App.IncentivesKeeper.Unstake(suite.Ctx, stake, sdk.Coins{})
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom0, 20)), balances)
	_, err = suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	// should be deleted
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestMultipleStakeLifecycle() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpecs: []depositSpec{
			{
				addr:   addr0,
				token0: sdk.NewInt64Coin("TokenA", 10),
				token1: sdk.NewInt64Coin("TokenB", 10),
				tick:   0,
				fee:    1,
			},
			{
				addr:   addr0,
				token0: sdk.NewInt64Coin("TokenA", 10),
				token1: sdk.NewInt64Coin("TokenB", 10),
				tick:   1,
				fee:    1,
			},
		},
		stakeDistEpochOffset: -2,
	})

	retrievedStake, err := suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)

	// unstake the full amount
	suite.App.IncentivesKeeper.Unstake(suite.Ctx, stake, sdk.Coins{})
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().Equal(
		sdk.NewCoins(
			sdk.NewInt64Coin(suite.LPDenom0, 20),
			sdk.NewInt64Coin(suite.LPDenom1, 20),
		), balances)
	_, err = suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	// should be deleted
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestStakeUnstakePartial() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpecs: []depositSpec{
			{
				addr:   addr0,
				token0: sdk.NewInt64Coin("TokenA", 10),
				token1: sdk.NewInt64Coin("TokenB", 10),
				tick:   0,
				fee:    1,
			},
		},
		stakeDistEpochOffset: -2,
	})

	retrievedStake, err := suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)

	// unstake the full amount
	err = suite.App.IncentivesKeeper.Unstake(
		suite.Ctx,
		stake,
		sdk.Coins{sdk.NewInt64Coin(suite.LPDenom0, 9)},
	)
	suite.Require().NoError(err)
	balances := suite.App.BankKeeper.GetAllBalances(suite.Ctx, addr0)
	suite.Require().ElementsMatch(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom0, 9)), balances)
	// should still be accessible
	retrievedStake, err = suite.App.IncentivesKeeper.GetStakeByID(suite.Ctx, stake.ID)
	suite.Require().NoError(err)
	suite.Require().NotNil(retrievedStake)
	suite.Require().
		ElementsMatch(sdk.NewCoins(sdk.NewInt64Coin(suite.LPDenom0, 11)), retrievedStake.Coins)

	// fin.
}

func (suite *KeeperTestSuite) TestStakesCoinsByQueryCondition(t *testing.T) {
	addr0 := suite.SetupAddr(0)

	stake := suite.SetupDepositAndStake(depositStakeSpec{
		depositSpecs: []depositSpec{
			{
				addr:   addr0,
				token0: sdk.NewInt64Coin("TokenA", 10),
				token1: sdk.NewInt64Coin("TokenB", 10),
				tick:   0,
				fee:    1,
			},
		},
		stakeDistEpochOffset: -2,
	})

	allCoins := sdk.Coins{
		sdk.NewInt64Coin(
			dextypes.NewPoolMetadata(
				1,
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				25,
				10,
			).Denom(),
			100,
		),
		sdk.NewInt64Coin(
			dextypes.NewPoolMetadata(
				2,
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				75,
				10,
			).Denom(),
			200,
		),
		sdk.NewInt64Coin(
			dextypes.NewPoolMetadata(
				3,
				&dextypes.PairID{
					Token0: "coin1",
					Token1: "coin2",
				},
				75,
				50,
			).Denom(),
			200,
		),
	}

	stakes := types.Stakes{
		types.NewStake(1, owner, sdk.Coins{allCoins[0]}, time.Time{}, 0),
		types.NewStake(2, owner, sdk.Coins{allCoins[1]}, time.Time{}, 0),
		types.NewStake(3, owner, sdk.Coins{allCoins[2]}, time.Time{}, 0),
	}

	pairID := &dextypes.PairID{
		Token0: "coin1",
		Token1: "coin2",
	}

	tests := []struct {
		name        string
		queryCond   types.QueryCondition
		coinsByCond sdk.Coins
	}{
		{
			name:        "All coins",
			queryCond:   types.QueryCondition{PairID: pairID, StartTick: 0, EndTick: 100},
			coinsByCond: sdk.Coins{allCoins[0], allCoins[1]},
		},
		{
			name:        "Coin1 only",
			queryCond:   types.QueryCondition{PairID: pairID, StartTick: 0, EndTick: 50},
			coinsByCond: sdk.Coins{allCoins[0]},
		},
		{
			name:        "Coin2 only",
			queryCond:   types.QueryCondition{PairID: pairID, StartTick: 50, EndTick: 100},
			coinsByCond: sdk.Coins{allCoins[1]},
		},
		{
			name:        "No coins",
			queryCond:   types.QueryCondition{PairID: pairID, StartTick: 100, EndTick: 200},
			coinsByCond: sdk.Coins{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins := keeper.StakeCoinsPassingQueryCondition(ctx, stakes, tt.queryCond)
			assert.Equal(t, tt.coinsByCond, coins)
		})
	}
}

func (suite *KeeperTestSuite) TestStakesCoinsByQueryConditionMultiple(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1xv9tklw7d82sezh9haa573wufgy59vmwe6xxe5")
	require.NoError(t, err)

	allCoins := sdk.Coins{}

	coins1 := sdk.NewInt64Coin(
		dextypes.NewPoolMetadata(
			1,
			&dextypes.PairID{
				Token0: "coin1",
				Token1: "coin2",
			},
			25,
			10,
		).Denom(),
		100,
	)
	allCoins = allCoins.Add(coins1)

	coins2 := sdk.NewInt64Coin(
		dextypes.NewPoolMetadata(
			2,
			&dextypes.PairID{
				Token0: "coin1",
				Token1: "coin2",
			},
			75,
			10,
		).Denom(),
		200,
	)
	allCoins = allCoins.Add(coins2)

	coins3 := sdk.NewInt64Coin(
		dextypes.NewPoolMetadata(
			3,
			&dextypes.PairID{
				Token0: "coin1",
				Token1: "coin2",
			},
			75,
			50,
		).Denom(),
		200,
	)
	allCoins = allCoins.Add(coins3)
	assert.Equal(t, allCoins, allCoins.Sort())

	stakes := types.Stakes{
		types.NewStake(1, owner, allCoins, time.Time{}, 0),
	}

	pairID := &dextypes.PairID{
		Token0: "coin1",
		Token1: "coin2",
	}

	tests := []struct {
		name      string
		queryCond types.QueryCondition
		expected  sdk.Coins
	}{
		{
			name:      "All coins",
			queryCond: types.QueryCondition{PairID: pairID, StartTick: 0, EndTick: 100},
			expected:  sdk.Coins{coins1, coins2},
		},
		{
			name:      "Coin1 only",
			queryCond: types.QueryCondition{PairID: pairID, StartTick: 0, EndTick: 50},
			expected:  sdk.Coins{coins1},
		},
		{
			name:      "Coin2 only",
			queryCond: types.QueryCondition{PairID: pairID, StartTick: 50, EndTick: 100},
			expected:  sdk.Coins{coins2},
		},
		{
			name:      "No coins",
			queryCond: types.QueryCondition{PairID: pairID, StartTick: 100, EndTick: 200},
			expected:  sdk.Coins{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins := suite.App.DexKeeper.StakeCoinsByQueryCondition(ctx, stakes, tt.queryCond)
			assert.Equal(t, tt.expected, coins)
		})
	}
}
