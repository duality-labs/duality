package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/duality-labs/duality/app/apptesting"
	"github.com/duality-labs/duality/x/incentives/keeper"
	"github.com/duality-labs/duality/x/incentives/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	QueryServer keeper.QueryServer
	MsgServer   types.MsgServer
}

// SetupTest sets incentives parameters from the suite's context
func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.QueryServer = keeper.NewQueryServer(suite.App.IncentivesKeeper)
	suite.MsgServer = keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
