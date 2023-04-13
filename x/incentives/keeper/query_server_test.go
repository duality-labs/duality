package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

var _ = suite.TestingSuite(nil)

func (suite *KeeperTestSuite) TestGetFutureRewardEstimate() {
	addr1 := suite.SetupAddr(0)
	suite.SetupDepositAndLock(depositSpec{
		addr:   addr1,
		token0: sdk.NewInt64Coin("TokenA", 10),
		token1: sdk.NewInt64Coin("TokenB", 10),
		tick:   0,
		fee:    1,
	})
	addr2 := suite.SetupAddr(1)
	suite.SetupDepositAndLock(depositSpec{
		addr:   addr2,
		token0: sdk.NewInt64Coin("TokenA", 10),
		token1: sdk.NewInt64Coin("TokenB", 10),
		tick:   0,
		fee:    1,
	})
	suite.SetupGauge(gaugeSpec{
		isPerpetual: false,
		rewards:     sdk.NewCoins(sdk.NewInt64Coin("foocoin", 1000)),
		paidOver:    100,
		startTick:   -10,
		endTick:     10,
		pricingTick: 0,
		startTime:   suite.Ctx.BlockTime(),
	})
	suite.SetupGauge(gaugeSpec{
		isPerpetual: false,
		rewards:     sdk.NewCoins(sdk.NewInt64Coin("foocoin", 1000)),
		paidOver:    100,
		startTick:   -10,
		endTick:     10,
		pricingTick: 0,
		startTime:   suite.Ctx.BlockTime().Add(315 * 24 * time.Hour),
	})
	estimate, err := suite.QueryServer.GetRewardsEstimate(suite.Ctx, addr1, nil, 365)
	suite.Require().NoError(err)
	suite.Require().Equal(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 750)), estimate)
}
