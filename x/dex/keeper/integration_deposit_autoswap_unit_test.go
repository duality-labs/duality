package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestAutoswapperWithdraws() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	bobDep0 := 10
	bobDep1 := 10
	tickIndex := 200
	feeIndex := 2

	bobSharesMinted := s.calcSharesMinted(int64(tickIndex), uint64(feeIndex), int64(bobDep0), int64(bobDep1))

	s.bobDeposits(NewDeposit(bobDep0, bobDep1, tickIndex, feeIndex))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(12, 5, tickIndex, feeIndex, DepositOptions{true}))
	s.assertAliceBalances(38, 45)
	s.assertDexBalances(22, 15)

	// Calculated expected amounts out
	autoswapSharesMinted := s.calcAutoswapSharesMinted(int64(tickIndex), uint64(feeIndex), 7, 0, 5, 5, bobSharesMinted.Int64(), bobSharesMinted.Int64())
	//totalShares := autoswapSharesMinted.Add(sdk.NewInt(20))

	aliceExpectedBalance0, aliceExpectedBalance1, dexExpectedBalance0, dexExpectedBalance1 := s.calcExpectedBalancesAfterWithdrawOnePool(autoswapSharesMinted, s.alice, int64(tickIndex), uint64(feeIndex))

	s.aliceWithdraws(NewWithdrawlInt(autoswapSharesMinted, int64(tickIndex), uint64(feeIndex)))

	s.assertAliceBalances(aliceExpectedBalance0.Int64(), aliceExpectedBalance1.Int64())
	s.assertDexBalances(dexExpectedBalance0.Int64(), dexExpectedBalance1.Int64())
}

func (s *MsgServerTestSuite) TestAutoswapOtherDepositorWithdraws() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	bobDep0 := 10
	bobDep1 := 10
	tickIndex := 150
	feeIndex := 3

	bobSharesMinted := s.calcSharesMinted(int64(tickIndex), uint64(feeIndex), int64(bobDep0), int64(bobDep1))

	s.bobDeposits(NewDeposit(bobDep0, bobDep1, tickIndex, feeIndex))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(10, 7, tickIndex, feeIndex, DepositOptions{true}))
	s.assertAliceBalances(40, 43)
	s.assertDexBalances(20, 17)

	// Calculated expected amounts out

	bobExpectedBalance0, bobExpectedBalance1, dexExpectedBalance0, dexExpectedBalance1 := s.calcExpectedBalancesAfterWithdrawOnePool(bobSharesMinted, s.bob, int64(tickIndex), uint64(feeIndex))

	s.bobWithdraws(NewWithdrawl(bobSharesMinted.Int64(), int64(tickIndex), uint64(feeIndex)))

	s.assertBobBalances(bobExpectedBalance0.Int64(), bobExpectedBalance1.Int64())
	s.assertDexBalances(dexExpectedBalance0.Int64(), dexExpectedBalance1.Int64())

}

func (s *MsgServerTestSuite) TestAutoswapBothWithdraws() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	bobDep0 := 10
	bobDep1 := 10
	tickIndex := 10000
	feeIndex := 3

	bobSharesMinted := s.calcSharesMinted(int64(tickIndex), uint64(feeIndex), int64(bobDep0), int64(bobDep1))

	s.bobDeposits(NewDeposit(bobDep0, bobDep1, tickIndex, feeIndex))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(10, 5, tickIndex, feeIndex, DepositOptions{true}))
	s.assertAliceBalances(40, 45)
	s.assertDexBalances(20, 15)

	// Calculated expected amounts out
	autoswapSharesMinted := s.calcAutoswapSharesMinted(int64(tickIndex), uint64(feeIndex), 5, 0, 5, 5, bobSharesMinted.Int64(), bobSharesMinted.Int64())
	//totalShares := autoswapSharesMinted.Add(sdk.NewInt(20))

	bobExpectedBalance0, bobExpectedBalance1, dexExpectedBalance0, dexExpectedBalance1 := s.calcExpectedBalancesAfterWithdrawOnePool(bobSharesMinted, s.bob, int64(tickIndex), uint64(feeIndex))

	s.bobWithdraws(NewWithdrawl(bobSharesMinted.Int64(), int64(tickIndex), uint64(feeIndex)))

	s.assertBobBalances(bobExpectedBalance0.Int64(), bobExpectedBalance1.Int64())
	s.assertDexBalances(dexExpectedBalance0.Int64(), dexExpectedBalance1.Int64())

	aliceExpectedBalance0 := sdk.NewInt(0)
	aliceExpectedBalance1 := sdk.NewInt(0)
	aliceExpectedBalance0, aliceExpectedBalance1, dexExpectedBalance0, dexExpectedBalance1 = s.calcExpectedBalancesAfterWithdrawOnePool(autoswapSharesMinted, s.alice, int64(tickIndex), uint64(feeIndex))

	s.aliceWithdraws(NewWithdrawlInt(autoswapSharesMinted, int64(tickIndex), uint64(feeIndex)))

	s.assertAliceBalances(aliceExpectedBalance0.Int64(), aliceExpectedBalance1.Int64())
	s.assertDexBalances(dexExpectedBalance0.Int64(), dexExpectedBalance1.Int64())
}
