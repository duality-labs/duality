package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestDepositMultiPartialFailure() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// no existing liquidity

	// WHEN
	// alice deposits 0 A, 5 B at tick 0 fee 0 and 5 A, 0 B at tick 0 fee 0
	s.aliceDeposits(
		NewDeposit(5, 0, 0, 0),
		NewDeposit(0, 5, 0, 0),
	)

	// THEN
	// only the first deposit should go through
	s.assertAliceBalances(45, 50)
	s.assertLiquidityAtTickInt(sdk.NewInt(5), sdk.NewInt(0), 0, 0)
	s.assertDexBalances(5, 0)
}

func (s *MsgServerTestSuite) TestDepositMultiCompleteFailure() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// alice deposits 5 A, 5 B at tick 0 fee 0
	s.aliceDeposits(NewDeposit(5, 5, 0, 0))

	// WHEN
	// alice deposits 0 A, 5 B at tick 0 fee 0 and 5 A, 0 B at tick 0 fee 0
	// THEN
	// deposit should fail with ErrAllDepositsFailed

	err := types.ErrAllDepositsFailed
	s.assertAliceDepositFails(err,
		NewDeposit(5, 0, 0, 0),
		NewDeposit(0, 5, 0, 0),
	)
}

func (s *MsgServerTestSuite) TestDepositMultiSuccess() {
	s.fundAliceBalances(50, 50)

	// GIVEN
	// no existing liquidity

	// WHEN
	// alice deposits 5 A, 5 B at tick 0 fee 0 and then 10 A, 10 B at tick 5 fee 0
	s.aliceDeposits(
		NewDeposit(5, 5, 0, 0),
		NewDeposit(10, 10, 0, 0),
	)

	// THEN
	// both deposits should go through
	s.assertAliceBalances(35, 35)
	s.assertLiquidityAtTickInt(sdk.NewInt(15), sdk.NewInt(15), 0, 0)
	s.assertDexBalances(15, 15)
}
