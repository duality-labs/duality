package keeper_test

import (
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = suite.TestingSuite(nil)

func (suite *KeeperTestSuite) TestGaugeLifecycle() {
	addr0 := suite.SetupAddr(0)

	// setup dex deposit and stake of those shares
	suite.SetupDepositAndStake(depositStakeSpec{
		depositSpec: depositSpec{
			addr:   addr0,
			token0: sdk.NewInt64Coin("TokenA", 10),
			token1: sdk.NewInt64Coin("TokenB", 10),
			tick:   0,
			fee:    1,
		},
		stakeTimeOffset: -24 * time.Hour,
	})

	// setup gauge starting 24 hours in the future
	suite.SetupGauge(gaugeSpec{
		startTime:   suite.Ctx.BlockTime().Add(24 * time.Hour),
		isPerpetual: false,
		rewards:     sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10)),
		paidOver:    2,
		startTick:   -10,
		endTick:     10,
		pricingTick: 0,
	})

	// assert that the gauge is not in effect yet by triggering an epoch end before gauge start
	suite.App.IncentivesKeeper.AfterEpochEnd(suite.Ctx, "day", 1)
	// no distribution yet
	require.Equal(suite.T(), "0foocoin", suite.App.BankKeeper.GetBalance(suite.Ctx, addr0, "foocoin").String())
	// assert that gauge state is well-managed
	require.Equal(suite.T(), len(suite.QueryServer.GetUpcomingGauges(suite.Ctx)), 1)
	require.Equal(suite.T(), len(suite.QueryServer.GetActiveGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetFinishedGauges(suite.Ctx)), 0)

	// advance time to epoch at or after the gauge starts, triggering distribution
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))
	suite.App.IncentivesKeeper.AfterEpochEnd(suite.Ctx, "day", 2)

	// assert that the gauge distributed
	require.Equal(suite.T(), "5foocoin", suite.App.BankKeeper.GetBalance(suite.Ctx, addr0, "foocoin").String())
	// assert that gauge state is well-managed
	require.Equal(suite.T(), len(suite.QueryServer.GetUpcomingGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetActiveGauges(suite.Ctx)), 1)
	require.Equal(suite.T(), len(suite.QueryServer.GetFinishedGauges(suite.Ctx)), 0)

	// advance to next epoch
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))
	suite.App.IncentivesKeeper.AfterEpochEnd(suite.Ctx, "day", 3)

	// assert new distribution
	require.Equal(suite.T(), "10foocoin", suite.App.BankKeeper.GetBalance(suite.Ctx, addr0, "foocoin").String())
	// assert that gauge state is well-managed
	require.Equal(suite.T(), len(suite.QueryServer.GetUpcomingGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetActiveGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetFinishedGauges(suite.Ctx)), 1)

	// repeat advancing to next epoch until gauge should be finished
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(24 * time.Hour))
	suite.App.IncentivesKeeper.AfterEpochEnd(suite.Ctx, "day", 4)

	// assert no additional distribution from finished gauge
	require.Equal(suite.T(), "10foocoin", suite.App.BankKeeper.GetBalance(suite.Ctx, addr0, "foocoin").String())
	// assert that gauge state is well-managed
	require.Equal(suite.T(), len(suite.QueryServer.GetUpcomingGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetActiveGauges(suite.Ctx)), 0)
	require.Equal(suite.T(), len(suite.QueryServer.GetFinishedGauges(suite.Ctx)), 1)
	// fin.
}
