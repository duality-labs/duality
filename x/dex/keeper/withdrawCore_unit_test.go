package keeper_test

import (
	"math"
	//. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	//"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestPartialWithdrawOnlyA() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of A at tick 0, fee tier 0
	// and then withdraws only 5 shares of A

	// DATA
	// Alice should be credited 10 total shares
	// Shares = amount0 + price1to0 * amount1
	// Shares = 10 + 0 * 0 = 10
	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))

	s.assertAliceBalances(45, 50)
	s.assertDexBalances(5, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)

}

func (s *MsgServerTestSuite) TestPartialWithdrawOnlyB() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of B at tick 0, fee tier 0
	// and then withdraws only 5 shares of B

	// DATA
	// Alice should be credited 10 total shares
	// Shares = amount0 + price1to0 * amount1
	// Shares = 10 + 0 * 0 = 10
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))

	s.assertAliceBalances(50, 45)
	s.assertDexBalances(0, 5)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
}
