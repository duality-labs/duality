package keeper_test

import (
	"math"
	//"time"
	//. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	//"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/NicholasDotSol/duality/testing_scripts"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))

	s.assertAliceBalances(45, 50)
	s.assertDexBalances(5, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)

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
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)

	s.aliceWithdraws(NewWithdrawl(5, 0, 0))

	s.assertAliceBalances(50, 45)
	s.assertDexBalances(0, 5)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)
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
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)

	s.aliceWithdraws(NewWithdrawl(10, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestCurrentTickUpdatesAfterDoubleSidedThenSingleSidedDepositAndPartialWithdrawal() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of A and B with a spread (fee) of +- 3 ticks
	// Alice then deposits 10 A with a spread (fee) of -1 ticks
	// Finally Alice withdraws from the first pool they deposited to

	s.aliceDeposits(NewDeposit(10, 10, 0, 1))

	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-3)
	s.assertCurr0To1(3)
	s.assertMinTick(-3)
	s.assertMaxTick(3)

	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(3)
	s.assertMinTick(-3)
	s.assertMaxTick(3)

	//DEBUG
	s.aliceWithdraws(NewWithdrawl(10, 0, 1))

	s.assertAliceBalances(35, 45)
	s.assertDexBalances(15, 5)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(3)
	s.assertMinTick(-3)
	s.assertMaxTick(3)
}

func (s *MsgServerTestSuite) TestCurrentTickUpdatesAfterDoubleSidedThenSingleSidedDepositAndFulllWithdrawal() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of A and B with a spread (fee) of +- 3 ticks
	// Alice then deposits 10 A with a spread (fee) of -1 ticks
	// Finally Alice withdraws from the first pool they deposited to

	s.aliceDeposits(NewDeposit(10, 10, 0, 1))

	s.assertAliceBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-3)
	s.assertCurr0To1(3)
	s.assertMinTick(-3)
	s.assertMaxTick(3)

	s.aliceDeposits(NewDeposit(10, 0, 0, 0))

	s.assertAliceBalances(30, 40)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(3)
	s.assertMinTick(-3)
	s.assertMaxTick(3)

	s.aliceWithdraws(NewWithdrawl(20, 0, 1))

	s.assertAliceBalances(40, 50)
	s.assertDexBalances(10, 0)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMinTick(-1)
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestTwoFullDoubleSidedRebalancedAtooMuchTick0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice deposits 10 of B and 5 of Aat tick 0, fee tier 0
	// Bob tries to deposit 10 of A and 10 of B
	// Thus Bob should only end up depositing 5 of A and 10 of B
	// Alice then withdraws
	// David then withdraws

	s.aliceDeposits(NewDeposit(5, 10, 0, 0))

	s.assertAliceBalances(45, 40)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(5, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.bobDeposits(NewDeposit(10, 10, 0, 0))

	s.assertAliceBalances(45, 40)
	s.assertBobBalances(45, 40)
	s.assertDexBalances(10, 20)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.aliceWithdraws(NewWithdrawl(15, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(45, 40)
	s.assertDexBalances(5, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.bobWithdraws(NewWithdrawl(15, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestTwoFullDoubleSidedRebalancedBtooMuchTick0() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// CASE
	// Alice deposits 10 of B and 5 of Aat tick 0, fee tier 0
	// Bob tries to deposit 10 of A and 10 of B
	// Thus Bob should only end up depositing 5 of A and 10 of B
	// Alice then withdraws
	// David then withdraws

	s.aliceDeposits(NewDeposit(10, 5, 0, 0))

	s.assertAliceBalances(40, 45)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(10, 5)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.bobDeposits(NewDeposit(10, 10, 0, 0))

	s.assertAliceBalances(40, 45)
	s.assertBobBalances(40, 45)
	s.assertDexBalances(20, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.aliceWithdraws(NewWithdrawl(15, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(40, 45)
	s.assertDexBalances(10, 5)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	s.bobWithdraws(NewWithdrawl(15, 0, 0))

	s.assertAliceBalances(50, 50)
	s.assertBobBalances(50, 50)
	s.assertDexBalances(0, 0)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(math.MaxInt64)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(math.MinInt64)
}

func (s *MsgServerTestSuite) TestFullWithdrawalFindNewMaxTickDoS() {
	s.fundAliceBalances(50, 50)
	// CASE
	// Alice deposits 10 of B at tick 0, fee tier 0
	// Alice then deposits 10 of B at tick 100000 (really large tick)
	// Alice then removes all of her liquidity from tick 100000

	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)

	s.aliceDeposits(NewDeposit(0, 10, 100000, 0))

	s.assertAliceBalances(50, 30)
	s.assertDexBalances(0, 20)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(100001)

	sharesToWithdraw := testing_scripts.SharesOnDeposit(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.NewDec(10), 100000)
	s.aliceWithdraws(NewWithdrawlDec(sharesToWithdraw, 100000, 0))

	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)
	s.assertCurr1To0(math.MinInt64)
	s.assertCurr0To1(1)
	s.assertMinTick(math.MaxInt64)
	s.assertMaxTick(1)
}
