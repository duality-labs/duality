package cli_test

import (
	gocontext "context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/duality-labs/duality/app/apptesting"
	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/keeper"
	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type QueryTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

// LockTokens funds an account, locks tokens and returns a lockID.
func (s *QueryTestSuite) SetupLock(addr sdk.AccAddress, coins sdk.Coins, duration time.Duration) (lockID uint64) {
	msgServer := keeper.NewMsgServerImpl(&s.App.IncentivesKeeper)
	s.FundAcc(addr, coins)

	msgResponse, err := msgServer.LockTokens(sdk.WrapSDKContext(s.Ctx), types.NewMsgSetupLock(addr, duration, coins))
	s.Require().NoError(err)

	return msgResponse.ID
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	denom := dexkeeper.NewDepositDenom(
		&dextypes.PairID{
			Token0: "TokenA",
			Token1: "TokenB",
		},
		0,
		1,
	).String()

	// set up lock with id = 1
	addr := apptesting.SetupAddr(0)
	s.SetupLock(addr, sdk.Coins{sdk.NewCoin(denom, sdk.NewInt(1000000))}, time.Hour*24)

	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		// {
		// 	"Query active gauges",
		// 	"/duality.incentives.Query/ActiveGauges",
		// 	&types.ActiveGetGaugesActiveUpcomingRequest{},
		// 	&types.ActiveGetGaugesActiveUpcomingResponse{},
		// },
		// {
		// 	"Query active gauges per denom",
		// 	"/duality.incentives.Query/ActiveGaugesPerDenom",
		// 	&types.ActiveGaugesPerDenomRequest{Denom: "stake"},
		// 	&types.ActiveGaugesPerDenomResponse{},
		// },
		// {
		// 	"Query gauge by id",
		// 	"/duality.incentives.Query/GetGaugeByID",
		// 	&types.GetGaugeByIDRequest{Id: 1},
		// 	&types.GetGaugeByIDResponse{},
		// },
		// {
		// 	"Query all gauges",
		// 	"/duality.incentives.Query/Gauges",
		// 	&types.GetGaugesActiveUpcomingRequest{},
		// 	&types.GetGaugesActiveUpcomingResponse{},
		// },
		// {
		// 	"Query lockable durations",
		// 	"/duality.incentives.Query/LockableDurations",
		// 	&types.QueryLockableDurationsRequest{},
		// 	&types.QueryLockableDurationsResponse{},
		// },
		// {
		// 	"Query module to distibute coins",
		// 	"/duality.incentives.Query/GetModuleCoinsToBeDistributed",
		// 	&types.GetModuleCoinsToBeDistributedRequest{},
		// 	&types.GetModuleCoinsToBeDistributedResponse{},
		// },
		// {
		// 	"Query reward estimate",
		// 	"/duality.incentives.Query/RewardsEst",
		// 	&types.RewardsEstRequest{Owner: s.TestAccs[0].String()},
		// 	&types.RewardsEstResponse{},
		// },
		// {
		// 	"Query upcoming gauges",
		// 	"/duality.incentives.Query/UpcomingGauges",
		// 	&types.UpcomingGetGaugesActiveUpcomingRequest{},
		// 	&types.UpcomingGetGaugesActiveUpcomingResponse{},
		// },
		// {
		// 	"Query upcoming gauges",
		// 	"/duality.incentives.Query/UpcomingGaugesPerDenom",
		// 	&types.UpcomingGaugesPerDenomRequest{Denom: "stake"},
		// 	&types.UpcomingGaugesPerDenomResponse{},
		// },
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
