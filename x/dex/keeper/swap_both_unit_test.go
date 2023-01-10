package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSwapNoLiqudityPairNotFound() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity

	// WHEN
	// swap 5 of tokenA
	// THEN
	// swap should fail with PairNotFound error
	err := types.ErrValidPairNotFound
	s.bobMarketSellFails(err, "TokenA", 5, 0)

}

func (s *MsgServerTestSuite) TestSwapExhaustFeeTiersAndLimitOrder() {

	// TODO: this fails due to fill and place key bug
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// place LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// Partially fill the LO, will have some token B remaining to fill
	s.bobMarketSells("TokenA", 5, 4)
	// in 5.000499950004999500, out 4.999500049995000500
	expectedAmountLeftSetup, amountOutSetup := s.calculateSingleSwapOnlyLOAToB(1, 10, 5)
	amountInSetup := sdk.NewInt(5).Sub(expectedAmountLeftSetup)
	s.assertLimitLiquidityAtTickInt("TokenB", 1, sdk.NewInt(10).Sub(amountOutSetup))

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	lpLiquiditySetup := sdk.NewInt(10)
	limitLiquiditySetup := sdk.NewInt(20).Sub(amountOutSetup)
	totalLiquiditySetup := lpLiquiditySetup.Add(limitLiquiditySetup)
	s.assertLimitLiquidityAtTick("TokenB", 1, limitLiquiditySetup.Int64())
	s.assertPoolLiquidity(0, 10, 0, 0)
	bobBalanceSetupB := sdk.NewInt(50).Sub(amountInSetup)
	s.assertBobBalancesInt(bobBalanceSetupB, amountOutSetup)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn := sdk.NewInt(30)
	s.bobMarketSells("TokenA", int(amountIn.Int64()), 0)

	// THEN
	// swap should have in out
	// swap effect on balances
	expectedTotalAmountLeft, expectedTotalAmountOut := s.calculateSingleSwapAToB(1, lpLiquiditySetup.Int64(), limitLiquiditySetup.Int64(), amountIn.Int64())
	expectedAmountIn := amountIn.Sub(expectedTotalAmountLeft)
	s.assertBobBalancesEpsilon(bobBalanceSetupB.Sub(expectedAmountIn), amountOutSetup.Add(expectedTotalAmountOut))
	s.assertDexBalancesEpsilon(expectedAmountIn.Add(amountInSetup), totalLiquiditySetup.Sub(expectedTotalAmountOut))

	// calculate amount traded against LPs (i.e. which gets swapped into fee tier)
	expectedAmountLeftAfterLP, _ := s.calculateSingleSwapNoLOAToB(1, lpLiquiditySetup.Int64(), amountIn.Int64())
	lpAmountIn := amountIn.Sub(expectedAmountLeftAfterLP)
	s.assertLiquidityAtTickInt(lpAmountIn, sdk.ZeroInt(), 0, 0)
	// limit orders exhausted
	s.assertLimitLiquidityAtTickInt("TokenB", 1, sdk.ZeroInt())
}
