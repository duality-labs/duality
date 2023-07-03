package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/duality-labs/duality/app/apptesting"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/keeper"
	"github.com/duality-labs/duality/x/incentives/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	QueryServer keeper.QueryServer
	MsgServer   types.MsgServer
	LPDenom0    string
	LPDenom1    string
}

// SetupTest sets incentives parameters from the suite's context
func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.QueryServer = keeper.NewQueryServer(suite.App.IncentivesKeeper)
	suite.MsgServer = keeper.NewMsgServerImpl(&suite.App.IncentivesKeeper)
	suite.LPDenom0 = dextypes.NewDepositDenom(
		&dextypes.PairID{
			Token0: "TokenA",
			Token1: "TokenB",
		},
		0,
		1,
	).String()
	suite.LPDenom1 = dextypes.NewDepositDenom(
		&dextypes.PairID{
			Token0: "TokenA",
			Token1: "TokenB",
		},
		1,
		1,
	).String()

	suite.SetEpochStartTime()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
