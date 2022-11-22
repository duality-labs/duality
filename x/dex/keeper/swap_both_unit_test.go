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

	// TODO: uncomment
	// s.assertLimitLiquidityAtTickDec("TokenB", 1, sdk.NewDec(20).Sub(amountOutSetup))
	s.assertLiquidityAtTick(0, 10, 0, 0)
	bobBalanceSetupB := sdk.NewDec(50).Sub(amountInSetup)
	s.assertBobBalancesDec(bobBalanceSetupB, amountOutSetup)

	// WHEN
	// swap 5 of token A for B with minOut 4
	amountIn, amountInDec := 20, sdk.NewDec(20)
	s.bobMarketSells("TokenA", amountIn, 0)

	// THEN
	// swap should have in out
	expectedAmountLeft, expectedAmountOut := s.calculateSingleSwapAToB(1, sdk.NewDec(10), sdk.NewDec(10), amountInDec)
	expectedAmountIn := amountInDec.Sub(expectedAmountLeft)
	s.assertBobBalancesDec(bobBalanceSetupB.Sub(expectedAmountIn), amountOutSetup.Add(expectedAmountOut))
	s.assertDexBalancesDec(expectedAmountIn.Add(amountInSetup), sdk.NewDec(20).Sub(amountOutSetup).Sub(expectedAmountOut))
	// TODO: uncomment
	// s.assertLimitLiquidityAtTickDec("TokenB", 1, sdk.NewDec(20).Sub(amountOutSetup).Sub(expectedAmountOut))
}
