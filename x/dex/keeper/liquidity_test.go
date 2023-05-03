package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dualityapp "github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type LiquidityTestSuite struct {
	suite.Suite
	app *dualityapp.App
	ctx sdk.Context
}

func (s *LiquidityTestSuite) SetupTest() {
	s.app = dualityapp.Setup(false)
	ctx := s.app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())
	s.ctx = ctx

	// app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	// app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	// queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	// types.RegisterQueryServer(queryHelper, app.DexKeeper)
	// queryClient := types.NewQueryClient(queryHelper)

	// accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, s.alice)
	// app.AccountKeeper.SetAccount(ctx, accAlice)
	// accBob := app.AccountKeeper.NewAccountWithAddress(ctx, s.bob)
	// app.AccountKeeper.SetAccount(ctx, accBob)
	// accCarol := app.AccountKeeper.NewAccountWithAddress(ctx, s.carol)
	// app.AccountKeeper.SetAccount(ctx, accCarol)
	// accDan := app.AccountKeeper.NewAccountWithAddress(ctx, s.dan)
	// app.AccountKeeper.SetAccount(ctx, accDan)
}

func TestLiquidityTestSuite(t *testing.T) {
	suite.Run(t, new(LiquidityTestSuite))
}

func (s *LiquidityTestSuite) TestSwapNoLiquidity() {
	// GIVEN
	// no liqudity of token A (deposit only token B and LO of token B)
	s.AddDeposit(NewDeposit(0, 10, 0, 1))
	s.AddGTCLimitOrder("TokenB", 1000, 10)

	// WHEN
	// swap 10 of tokenB
	// THEN
	// swap should do 0 swap
	tokenIn, tokenOut := s.Swap("TokenB", "TokenA", 10, 0, nil)
	s.AssertSwapOutput(tokenIn, 0, tokenOut, 0)
}

func (s *LiquidityTestSuite) TestSwapAToBPartialFillLP() {
	// GIVEN
	// 10 tokenB LP @ tick 0 fee 1
	s.AddDeposit(NewDeposit(0, 10, 0, 1))

	// WHEN
	// swap 20 of tokenA
	// THEN
	// swap should return 11 TokenA in and 10 TokenB out
	tokenIn, tokenOut := s.Swap("TokenA", "TokenB", 20, 0, nil)

	s.Assert().Equal("TokenA", tokenIn.Denom)
	s.Assert().Equal("TokenB", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 11, tokenOut, 10)
}

func (s *LiquidityTestSuite) TestSwapBToAPartialFillLP() {
	// GIVEN
	// 10 tokenA LP @ tick 0 fee 1
	s.AddDeposit(NewDeposit(10, 0, 0, 1))

	// WHEN
	// swap 20 of tokenB
	// THEN
	// swap should return 11 TokenB in and 10 TokenA out
	tokenIn, tokenOut := s.Swap("TokenB", "TokenA", 20, 0, nil)

	s.Assert().Equal("TokenB", tokenIn.Denom)
	s.Assert().Equal("TokenA", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 11, tokenOut, 10)
}

func (s *LiquidityTestSuite) TestSwapAToBFillLP() {
	// GIVEN
	// 100 tokenB LP @ tick 200 fee 5
	s.AddDeposit(NewDeposit(0, 100, 200, 5))

	// WHEN
	// swap 100 of tokenA
	// THEN
	// swap should return 100 TokenA in and 97 TokenB out
	tokenIn, tokenOut := s.Swap("TokenA", "TokenB", 100, 0, nil)

	s.Assert().Equal("TokenA", tokenIn.Denom)
	s.Assert().Equal("TokenB", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 100, tokenOut, 97)
}

func (s *LiquidityTestSuite) TestSwapBToAFillLP() {
	// GIVEN
	// 100 tokenA LP @ tick -20000 fee 1
	s.AddDeposit(NewDeposit(100, 0, -20000, 1))

	// WHEN
	// swap 100 of tokenB
	// THEN
	// swap should return 97 TokenA in and 13 TokenB out
	tokenIn, tokenOut := s.Swap("TokenB", "TokenA", 100, 0, nil)

	s.Assert().Equal("TokenB", tokenIn.Denom)
	s.Assert().Equal("TokenA", tokenOut.Denom)
	// NOTE: Given rounding for amountOut, amountIn does not use the full maxAmount
	s.AssertSwapOutput(tokenIn, 97, tokenOut, 13)
}

func (s *LiquidityTestSuite) TestSwapAToBFillLPHighFee() {
	// GIVEN
	// 100 tokenB LP @ tick 20,000 fee 1000
	s.AddDeposit(NewDeposit(0, 100, 20000, 1000))

	// WHEN
	// swap 100 of tokenA
	// THEN
	// swap should return 98 TokenA in and 12 TokenB out
	tokenIn, tokenOut := s.Swap("TokenA", "TokenB", 100, 0, nil)

	s.Assert().Equal("TokenA", tokenIn.Denom)
	s.Assert().Equal("TokenB", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 98, tokenOut, 12)
}

func (s *LiquidityTestSuite) TestSwapBToAFillLPHighFee() {
	// GIVEN
	// 1000 tokenA LP @ tick 20000 fee 1000
	s.AddDeposit(NewDeposit(1000, 0, 20000, 1000))

	// WHEN
	// swap 100 of tokenB
	// THEN
	// swap should return 100 TokenB in and 668 TokenA out
	tokenIn, tokenOut := s.Swap("TokenB", "TokenA", 100, 0, nil)

	s.Assert().Equal("TokenB", tokenIn.Denom)
	s.Assert().Equal("TokenA", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 100, tokenOut, 668)
}

// swap multi LP

func (s *LiquidityTestSuite) TestSwapAToBPartialFillLO() {
	// GIVEN
	// 10 tokenB LO @ tick 1000
	s.AddGTCLimitOrder("TokenB", 10, 1000)

	// WHEN
	// swap 20 of tokenA
	// THEN
	// swap should return 11 TokenA in and 9 TokenB out
	tokenIn, tokenOut := s.Swap("TokenA", "TokenB", 20, 0, nil)

	s.Assert().Equal("TokenA", tokenIn.Denom)
	s.Assert().Equal("TokenB", tokenOut.Denom)
	s.AssertSwapOutput(tokenIn, 11, tokenOut, 9)
}

// Test helpers ///////////////////////////////////////////////////////////////

func (s *LiquidityTestSuite) AddDeposit(deposit *Deposit) {
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, defaultPairID, deposit.TickIndex, deposit.Fee)
	s.Assert().NoError(err)
	pool.LowerTick0.Reserves = pool.LowerTick0.Reserves.Add(deposit.AmountA)
	pool.UpperTick1.Reserves = pool.UpperTick1.Reserves.Add(deposit.AmountB)
	s.app.DexKeeper.SavePool(s.ctx, pool)
}

func (s *LiquidityTestSuite) AddDeposits(deposits ...*Deposit) {
	for _, deposit := range deposits {
		s.AddDeposit(deposit)
	}
}

func (s *LiquidityTestSuite) AddGTCLimitOrder(tokenIn string, amountIn int64, tickIndex int64) {
	tranche, err := s.app.DexKeeper.GetOrInitPlaceTranche(s.ctx, defaultPairID, tokenIn, tickIndex, nil, types.LimitOrderType_GOOD_TIL_CANCELLED)
	s.Assert().NoError(err)
	tranche.PlaceMakerLimitOrder(sdk.NewInt(amountIn))
	s.app.DexKeeper.SaveTranche(s.ctx, tranche)
}

func (s *LiquidityTestSuite) Swap(tokenIn string, tokenOut string, maxAmountIn int64, maxAmountOut int64, limitPrice *sdk.Dec) (coinIn, coinOut sdk.Coin) {
	coinIn, coinOut, err := s.app.DexKeeper.Swap(s.ctx, defaultPairID, tokenIn, tokenOut, sdk.NewInt(maxAmountIn), sdk.NewInt(maxAmountOut), limitPrice)
	s.Assert().NoError(err)
	return coinIn, coinOut
}

func (s *LiquidityTestSuite) AssertSwapOutput(actualIn sdk.Coin, expectedIn int64, actualOut sdk.Coin, expectedOut int64) {
	amtIn := actualIn.Amount
	amtOut := actualOut.Amount

	s.Assert().True(amtIn.Equal(sdk.NewInt(expectedIn)), "Expected amountIn %d != %s", expectedIn, amtIn)
	s.Assert().True(amtOut.Equal(sdk.NewInt(expectedOut)), "Expected amountOut %d != %s", expectedOut, amtOut)
}

// func (s *LiquidityTestSuite) SwapSuccess(tokenIn string, tokenOut string, maxAmountIn int64, maxAmountOut int64, limitPrice *sdk.Dec) (totalInCoin, totalOutCoin sdk.Coin) {
// 	tokenIn, tokenOut, err := s.Swap(tokenIn, tokenOut, maxAmountIn, maxAmountOut, limitPrice)
// 	s.Assert().NoError(err)
// }

// func (s *LiquidityTestSuite) SwapFails(expectedErr error, tokenIn string, tokenOut string, maxAmountIn int64, maxAmountOut int64, limitPrice *sdk.Dec) {
// 	_, _, err := s.Swap(tokenIn, tokenOut, maxAmountIn, maxAmountOut, limitPrice)
// s.Assert().E (err)
// }
