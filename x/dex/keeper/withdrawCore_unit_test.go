package keeper_test

import (
	"fmt"
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
	s.assertCurr0To1(math.MinInt64)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))
	fmt.Println("here")

	s.assertAliceBalances(45, 50)
	s.assertDexBalances(5, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MinInt64)

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
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))

	s.assertAliceBalances(50, 45)
	s.assertDexBalances(0, 5)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)
}

func (s *MsgServerTestSuite) TestFullWithdrawOnlyB() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of B at tick 0, fee tier 0
	// and then withdraws 10 shares of B

	// DATA
	// Alice should be credited 10 total shares
	// Shares = amount0 + price1to0 * amount1
	// Shares = 10 + 0 * 0 = 10
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(1)

	s.aliceWithdraws(NewWithdrawl(10, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MaxInt64)
	s.assertCurr0To1(math.MinInt64)
}

func (s *MsgServerTestSuite) TestCurrentTickUpdatesAfterOneSidedDeposit() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of A and B with a spread (fee) of +- 3 ticks
	// Alice then deposits 10 A with a spread (fee) of -1 ticks
	// Since there is no B in Alice's second deposit, the current tick shouldn't update

	s.aliceDeposits(NewDeposit(10, 10, 0, 1))

	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-3)
	s.assertCurr0To1(3)

	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(3)
}