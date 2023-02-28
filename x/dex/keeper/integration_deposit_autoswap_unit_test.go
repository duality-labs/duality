package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/utils"
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

	amount0, amount1 := sdk.NewInt(int64(bobDep0)), sdk.NewInt(int64(bobDep1))
	centerPrice, _ := utils.CalcPrice1To0(int64(tickIndex))
	bobSharesMinted := amount0.ToDec().Add(centerPrice.Mul(amount1.ToDec())).TruncateInt()

	s.bobDeposits(NewDeposit(bobDep0, bobDep1, tickIndex, feeIndex))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(12, 5, tickIndex, feeIndex, DepositOptions{true}))
	s.assertAliceBalances(38, 45)
	s.assertDexBalances(22, 15)

	residual0, residual1, balanced0, balanced1, totalShares, valuePool := sdk.NewInt(7), sdk.NewInt(0), sdk.NewInt(5), sdk.NewInt(5), sdk.NewInt(bobSharesMinted.Int64()), sdk.NewInt(bobSharesMinted.Int64())

	// residualValue = 1.0001^-f * residualAmount0 + 1.0001^{i-f} * residualAmount1
	// balancedValue = balancedAmount0 + 1.0001^{i} * balancedAmount1
	// value = residualValue + balancedValue
	// shares minted = value * totalShares / valuePool
	fee := int64(s.feeTiers[uint64(feeIndex)].Fee)

	centerPrice, _ = utils.CalcPrice1To0(int64(tickIndex))
	leftPrice, _ := utils.CalcPrice1To0(int64(tickIndex) - fee)
	discountPrice, _ := utils.CalcPrice1To0(- fee)

	balancedValue := balanced0.ToDec().Add(centerPrice.Mul(balanced1.ToDec())).TruncateInt()
	residualValue := residual0.ToDec().Mul(discountPrice).Add(leftPrice.Mul(residual1.ToDec())).TruncateInt()
	valueMint := balancedValue.Add(residualValue)

	// Calculated expected amounts out
	autoswapSharesMinted := valueMint.Mul(totalShares).Quo(valuePool)

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

	amount0, amount1 := sdk.NewInt(int64(bobDep0)), sdk.NewInt(int64(bobDep1))
	centerPrice, _ := utils.CalcPrice1To0(int64(tickIndex))
	bobSharesMinted := amount0.ToDec().Add(centerPrice.Mul(amount1.ToDec())).TruncateInt()

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

	amount0, amount1 := sdk.NewInt(int64(bobDep0)), sdk.NewInt(int64(bobDep1))
	centerPrice, _ := utils.CalcPrice1To0(int64(tickIndex))
	bobSharesMinted := amount0.ToDec().Add(centerPrice.Mul(amount1.ToDec())).TruncateInt()

	s.bobDeposits(NewDeposit(bobDep0, bobDep1, tickIndex, feeIndex))
	s.assertBobBalances(40, 40)
	s.assertDexBalances(10, 10)

	// Alice deposits at a different balance ratio
	s.aliceDepositsWithOptions(NewDepositWithOptions(10, 5, tickIndex, feeIndex, DepositOptions{true}))
	s.assertAliceBalances(40, 45)
	s.assertDexBalances(20, 15)

	residual0, residual1, balanced0, balanced1, totalShares, valuePool := sdk.NewInt(5), sdk.NewInt(0), sdk.NewInt(5), sdk.NewInt(5), sdk.NewInt(bobSharesMinted.Int64()), sdk.NewInt(bobSharesMinted.Int64())

	// residualValue = 1.0001^-f * residualAmount0 + 1.0001^{i-f} * residualAmount1
	// balancedValue = balancedAmount0 + 1.0001^{i} * balancedAmount1
	// value = residualValue + balancedValue
	// shares minted = value * totalShares / valuePool
	fee := int64(s.feeTiers[uint64(feeIndex)].Fee)

	centerPrice, _ = utils.CalcPrice1To0(int64(tickIndex))
	leftPrice, _ := utils.CalcPrice1To0(int64(tickIndex) - fee)
	discountPrice, _ := utils.CalcPrice1To0(- fee)

	balancedValue := balanced0.ToDec().Add(centerPrice.Mul(balanced1.ToDec())).TruncateInt()
	residualValue := residual0.ToDec().Mul(discountPrice).Add(leftPrice.Mul(residual1.ToDec())).TruncateInt()
	valueMint := balancedValue.Add(residualValue)

	// Calculated expected amounts out
	autoswapSharesMinted := valueMint.Mul(totalShares).Quo(valuePool)

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
