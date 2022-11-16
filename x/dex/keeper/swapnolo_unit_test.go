package keeper_test

import (
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSwapNoLONoLiqudityPairNotFound() {
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

func (s *MsgServerTestSuite) TestSwapNoLONoLiqudity() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	// GIVEN
	// no liqudity of token A
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	// WHEN
	// swap 5 of tokenB

	// THEN
	// swap should fail with Error Not enough coins
	err := types.ErrNotEnoughCoins
	s.bobMarketSellFails(err, "TokenB", 5, 0)

}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledMaxReachedSlippageToleranceReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)

	// WHEN
	// swap 20 of tokenA with minOut 15

	// THEN
	// swap should fail with ErrNotEnoughCoins error
	err := types.ErrNotEnoughCoins
	s.bobMarketSellFails(err, "TokenA", 20, 15)

}

func (s *MsgServerTestSuite) TestSwapNoLOPartiallyFilledMaxReachedSlippageToleranceNotReached() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)
	// GIVEN
	// deposit 10 of tokenB
	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.assertAliceBalances(50, 40)
	s.assertDexBalances(0, 10)

	// WHEN
	// swap 20 of tokenA at
	s.bobMarketSells("TokenA", 20, 5)

	// THEN
	// swap should fail with ErrNotEnoughCoins error
	expectedAmountIn, expectedAmountOut := s.calculateSingleSwapNoLOAToB(1, NewDec(10), NewDec(20))

	s.assertBobBalancesDec(expectedAmountOut)
	s.assertDexBalancesDec(expectedAmountIn, NewDec(10).Sub(expectedAmountOut))

}
