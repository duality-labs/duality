package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/duality-labs/duality/app/apptesting"
	"github.com/duality-labs/duality/x/lockup/keeper"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	querier keeper.Querier
	cleanup func()
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.querier = keeper.NewQuerier(suite.App.LockupKeeper)
	// unbondingDuration := suite.App.StakingKeeper.GetParams(suite.Ctx).UnbondingTime
	// suite.App.IncentivesKeeper.SetLockableDurations(suite.Ctx, []time.Duration{
	// 	time.Hour * 24 * 14,
	// 	time.Hour,
	// 	time.Hour * 3,
	// 	time.Hour * 7,
	// 	unbondingDuration,
	// })
}

func (suite *KeeperTestSuite) Cleanup() {
	suite.cleanup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
