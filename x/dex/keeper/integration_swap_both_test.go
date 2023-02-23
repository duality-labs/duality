package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

func (s *MsgServerTestSuite) TestSwapNoLiqudityPairNotFound() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 50)
	err := types.ErrSlippageLimitReached
	s.bobMarketSellFails(err, "TokenA", 5, 0)
}

func (s *MsgServerTestSuite) TestSwapExhaustFeeTiersAndLimitOrder() {
	s.fundAliceBalances(50, 50)
	s.fundBobBalances(50, 0)

	s.aliceLimitSells("TokenB", 0, 10)

	s.bobMarketSells("TokenA", 5, 5)

	s.assertLimitLiquidityAtTick("TokenB", 0, 5)

	s.aliceLimitSells("TokenB", 0, 10)

	s.aliceDeposits(NewDeposit(0, 10, 0, 0))

	s.assertBobBalances(45, 5)

	s.bobMarketSells("TokenA", 30, 0)

	s.assertPoolLiquidity(11, 0, 0, 0)

	s.assertLimitLiquidityAtTickInt("TokenB", 0, sdk.ZeroInt())
}
