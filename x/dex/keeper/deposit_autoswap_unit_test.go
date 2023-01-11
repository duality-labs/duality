package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestAutoswapperWithdraws() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.bobDeposits(NewDeposit(10, 10, 0, 0))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(10, 5, 0, 0, DepositOptions{true}))
	s.assertAliceBalances(40, 45)
	s.assertDexBalances(20, 15)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// Calculated expected amounts out
	autoswapSharesMinted := s.calcAutoswapSharesMinted(0, 1, 5, 0, 5, 5, 20, 20)
	totalShares := autoswapSharesMinted.Add(sdk.NewInt(20))

	dexTotalAmount0 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenA").Amount
	dexTotalAmount1 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenB").Amount
	amountOut0 := dexTotalAmount0.Mul(autoswapSharesMinted).Quo(totalShares)
	amountOut1 := dexTotalAmount1.Mul(autoswapSharesMinted).Quo(totalShares)
	aliceExpectedBalance0 := sdk.NewInt(40).Add(amountOut0)
	aliceExpectedBalance1 := sdk.NewInt(45).Add(amountOut1)

	dexExpectedBalance0 := sdk.NewInt(20).Sub(amountOut0)
	dexExpectedBalance1 := sdk.NewInt(15).Sub(amountOut1)

	s.aliceWithdraws(NewWithdrawlInt(autoswapSharesMinted, 0, 0))

	aliceActualBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenA").Amount
	aliceActualBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenB").Amount

	fmt.Println("Alice Expected Balance 0: ", aliceExpectedBalance0)
	fmt.Println("Alice Expected Balance 1: ", aliceExpectedBalance1)
	fmt.Println("Alice Actual Balance 0: ", aliceActualBalance0)
	fmt.Println("Alice Actual Balance 1: ", aliceActualBalance1)

	s.assertAliceBalancesEpsilon(aliceExpectedBalance0, aliceExpectedBalance1)
	s.assertDexBalancesEpsilon(dexExpectedBalance0, dexExpectedBalance1)
}

func (s *MsgServerTestSuite) TestAutoswapOtherDepositorWithdraws() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)

	// GIVEN
	// create spread around -1, 1
	s.bobDeposits(NewDeposit(10, 10, 0, 0))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(10, 5, 0, 0, DepositOptions{true}))
	s.assertAliceBalances(40, 45)
	s.assertDexBalances(20, 15)
	s.assertCurr1To0(-1)
	s.assertCurr0To1(1)
	s.assertMinTick(-1)
	s.assertMaxTick(1)

	// Calculated expected amounts out
	autoswapSharesMinted := s.calcAutoswapSharesMinted(0, 1, 5, 0, 5, 5, 20, 20)
	totalShares := autoswapSharesMinted.Add(sdk.NewInt(20))

	dexTotalAmount0 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenA").Amount
	dexTotalAmount1 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenB").Amount
	bobAmountOut0 := dexTotalAmount0.Mul(sdk.NewInt(20)).Quo(totalShares)
	bobAmountOut1 := dexTotalAmount1.Mul(sdk.NewInt(20)).Quo(totalShares)
	bobExpectedBalance0 := sdk.NewInt(40).Add(bobAmountOut0)
	bobExpectedBalance1 := sdk.NewInt(40).Add(bobAmountOut1)

	dexExpectedBalance0 := sdk.NewInt(20).Sub(bobAmountOut0)
	dexExpectedBalance1 := sdk.NewInt(15).Sub(bobAmountOut1)

	s.bobWithdraws(NewWithdrawl(20, 0, 0))

	bobActualBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenA").Amount
	bobActualBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenB").Amount

	fmt.Println("Bob Expected Balance 0: ", bobExpectedBalance0)
	fmt.Println("Bob Expected Balance 1: ", bobExpectedBalance1)
	fmt.Println("Bob Actual Balance 0: ", bobActualBalance0)
	fmt.Println("Bob Actual Balance 1: ", bobActualBalance1)

	s.assertBobBalancesEpsilon(bobExpectedBalance0, bobExpectedBalance1)
	s.assertDexBalancesEpsilon(dexExpectedBalance0, dexExpectedBalance1)

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

	bobActualBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenA").Amount
	bobActualBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenB").Amount

	fmt.Println("Bob Expected Balance 0: ", bobExpectedBalance0)
	fmt.Println("Bob Expected Balance 1: ", bobExpectedBalance1)
	fmt.Println("Bob Actual Balance 0: ", bobActualBalance0)
	fmt.Println("Bob Actual Balance 1: ", bobActualBalance1)

	s.assertBobBalancesEpsilon(bobExpectedBalance0, bobExpectedBalance1)
	s.assertDexBalancesEpsilon(dexExpectedBalance0, dexExpectedBalance1)

	aliceExpectedBalance0 := sdk.NewInt(0)
	aliceExpectedBalance1 := sdk.NewInt(0)
	aliceExpectedBalance0, aliceExpectedBalance1, dexExpectedBalance0, dexExpectedBalance1 = s.calcExpectedBalancesAfterWithdrawOnePool(autoswapSharesMinted, s.alice, int64(tickIndex), uint64(feeIndex))

	s.aliceWithdraws(NewWithdrawlInt(autoswapSharesMinted, int64(tickIndex), uint64(feeIndex)))

	aliceActualBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenA").Amount
	aliceActualBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenB").Amount

	fmt.Println("Alice Expected Balance 0: ", aliceExpectedBalance0)
	fmt.Println("Alice Expected Balance 1: ", aliceExpectedBalance1)
	fmt.Println("Alice Actual Balance 0: ", aliceActualBalance0)
	fmt.Println("Alice Actual Balance 1: ", aliceActualBalance1)

	s.assertAliceBalancesEpsilon(aliceExpectedBalance0, aliceExpectedBalance1)
	s.assertDexBalancesEpsilon(dexExpectedBalance0, dexExpectedBalance1)
}
