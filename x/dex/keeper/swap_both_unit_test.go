package keeper_test

import (
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
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
	expectedAmountLeftSetup, amountOutSetup := s.calculateSingleSwapOnlyLOAToB(1, NewDec(10), NewDec(5))
	amountInSetup := sdk.NewDec(5).Sub(expectedAmountLeftSetup)
	s.assertLimitLiquidityAtTickDec("TokenB", sdk.NewDec(10).Sub(amountOutSetup), 1)

	// place another LO selling 10 of token B at tick 1
	s.aliceLimitSells("TokenB", 1, 10)

	// deposit 10 of token B at tick 0 fee 1
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	lpLiquiditySetup := sdk.NewDec(10)
	limitLiquiditySetup := sdk.NewDec(20).Sub(amountOutSetup)
	totalLiquiditySetup := lpLiquiditySetup.Add(limitLiquiditySetup)
	s.assertLimitLiquidityAtTickDec("TokenB", limitLiquiditySetup, 1)
	s.assertLiquidityAtTick(0, 10, 0, 0)
	bobBalanceSetupB := sdk.NewDec(50).Sub(amountInSetup)
	s.assertBobBalancesDec(bobBalanceSetupB, amountOutSetup)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn, amountInDec := 30, sdk.NewDec(30)
	s.bobMarketSells("TokenA", amountIn, 0)

	// THEN
	// swap should have in out
	// swap effect on balances
	expectedTotalAmountLeft, expectedTotalAmountOut := s.calculateSingleSwapAToB(1, lpLiquiditySetup, limitLiquiditySetup, amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedTotalAmountLeft)
	s.assertBobBalancesDec(bobBalanceSetupB.Sub(expectedAmountIn), amountOutSetup.Add(expectedTotalAmountOut))
	s.assertDexBalancesDec(expectedAmountIn.Add(amountInSetup), totalLiquiditySetup.Sub(expectedTotalAmountOut))

	// calculate amount traded against LPs (i.e. which gets swapped into fee tier)
	expectedAmountLeftAfterLP, _ := s.calculateSingleSwapNoLOAToB(1, lpLiquiditySetup, amountInDec)
	lpAmountIn := amountInDec.Sub(expectedAmountLeftAfterLP)
	s.assertLiquidityAtTickDec(lpAmountIn, sdk.ZeroDec(), 0, 0)
	// limit orders exhausted
	s.assertLimitLiquidityAtTickDec("TokenB", sdk.ZeroDec(), 1)
}
