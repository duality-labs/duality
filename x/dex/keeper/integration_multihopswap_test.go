package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type PoolSetup struct {
	TokenA    string
	TokenB    string
	AmountA   int
	AmountB   int
	TickIndex int
	Fee       int
}

func NewPoolSetup(tokenA, tokenB string, amountA, amountB, tickIndex, fee int) PoolSetup {
	return PoolSetup{
		TokenA:    tokenA,
		TokenB:    tokenB,
		AmountA:   amountA,
		AmountB:   amountB,
		TickIndex: tickIndex,
		Fee:       fee,
	}
}

func (s *MsgServerTestSuite) SetupMultiplePools(poolSetups ...PoolSetup) {
	for _, p := range poolSetups {
		coins := sdk.NewCoins(
			sdk.NewCoin(p.TokenA, sdk.NewInt(int64(p.AmountA))),
			sdk.NewCoin(p.TokenB, sdk.NewInt(int64(p.AmountB))),
		)
		s.fundAccountBalancesExotic(s.bob, coins)
		pairID := types.PairID{Token0: p.TokenA, Token1: p.TokenB}
		s.deposits(
			s.bob,
			[]*Deposit{NewDeposit(p.AmountA, p.AmountB, p.TickIndex, p.Fee)},
			pairID,
		)
	}
}

func (s *MsgServerTestSuite) TestMultiHopSwapSingleRoute() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D,
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 100, 0, 1),
	)

	// WHEN alice multihopswaps A<>B => B<>C => C<>D,
	route := [][]string{{"TokenA", "TokenB", "TokenC", "TokenD"}}
	s.aliceMultiHopSwaps(route, 100, sdk.MustNewDecFromStr("0.9"), false)

	// THEN alice gets out 99 TokenD
	s.assertAccountBalanceExotic(s.alice, "TokenA", 0)
	s.assertAccountBalanceExotic(s.alice, "TokenD", 99)

	s.assertDexBalanceExotic("TokenA", 100)
	s.assertDexBalanceExotic("TokenB", 100)
	s.assertDexBalanceExotic("TokenC", 100)
	s.assertDexBalanceExotic("TokenD", 1)
}

func (s *MsgServerTestSuite) TestMultiHopSwapInsufficientLiquiditySingleRoute() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D with insufficient liquidity in C<>D
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 50, 0, 1),
	)

	// THEN alice multihopswap fails
	route := [][]string{{"TokenA", "TokenB", "TokenC", "TokenD"}}
	s.aliceMultiHopSwapFails(types.ErrInsufficientLiquidity, route, 100, sdk.MustNewDecFromStr("0.9"), false)
}

func (s *MsgServerTestSuite) TestMultiHopSwapLimitPriceNotMetSingleRoute() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D with insufficient liquidity in C<>D
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 100, 1200, 1),
	)

	// THEN alice multihopswap fails
	route := [][]string{{"TokenA", "TokenB", "TokenC", "TokenD"}}
	s.aliceMultiHopSwapFails(types.ErrExitLimitPriceHit, route, 50, sdk.MustNewDecFromStr("0.9"), false)
}

func (s *MsgServerTestSuite) TestMultiHopSwapMultiRouteOneGood() {
	s.fundAliceBalances(100, 0)

	// GIVEN viable liquidity in pools A<>B, B<>E, E<>X
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenX", 0, 50, 0, 1),
		NewPoolSetup("TokenC", "TokenX", 0, 50, 2200, 1),
		NewPoolSetup("TokenB", "TokenD", 0, 100, 0, 1),
		NewPoolSetup("TokenD", "TokenX", 0, 50, 0, 1),
		NewPoolSetup("TokenD", "TokenX", 0, 50, 2200, 1),
		NewPoolSetup("TokenB", "TokenE", 0, 100, 0, 1),
		NewPoolSetup("TokenE", "TokenX", 0, 100, 0, 1),
	)

	// WHEN alice multihopswaps with three routes the first two routes fail and the third works
	routes := [][]string{
		{"TokenA", "TokenB", "TokenC", "TokenX"},
		{"TokenA", "TokenB", "TokenD", "TokenX"},
		{"TokenA", "TokenB", "TokenE", "TokenX"},
	}
	s.aliceMultiHopSwaps(routes, 100, sdk.MustNewDecFromStr("0.9"), false)

	// THEN swap succeeds through route A<>B, B<>E, E<>X

	s.assertAccountBalanceExotic(s.alice, "TokenA", 0)
	s.assertAccountBalanceExotic(s.alice, "TokenX", 99)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenA", Token1: "TokenB"}, sdk.NewInt(100), sdk.NewInt(1), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenE"}, sdk.NewInt(100), sdk.NewInt(1), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenE", Token1: "TokenX"}, sdk.NewInt(100), sdk.NewInt(1), 0, 1)

	// Other pools are unaffected
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenC"}, sdk.NewInt(0), sdk.NewInt(100), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenC", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(50), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenC", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(50), 2200, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenD"}, sdk.NewInt(0), sdk.NewInt(100), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenD", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(50), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenD", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(50), 2200, 1)
}

func (s *MsgServerTestSuite) TestMultiHopSwapMultiRouteAllFail() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in sufficient liquidity but inadequate prices
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenX", 0, 50, 0, 1),
		NewPoolSetup("TokenC", "TokenX", 0, 50, 2200, 1),
		NewPoolSetup("TokenB", "TokenD", 0, 100, 0, 1),
		NewPoolSetup("TokenD", "TokenX", 0, 50, 0, 1),
		NewPoolSetup("TokenD", "TokenX", 0, 50, 2200, 1),
		NewPoolSetup("TokenB", "TokenE", 0, 50, 0, 1),
		NewPoolSetup("TokenE", "TokenX", 0, 50, 2200, 1),
	)

	// WHEN alice multihopswaps with three routes they all fail
	routes := [][]string{
		{"TokenA", "TokenB", "TokenC", "TokenX"},
		{"TokenA", "TokenB", "TokenD", "TokenX"},
		{"TokenA", "TokenB", "TokenE", "TokenX"},
	}

	// Then fails with findBestRoute
	s.aliceMultiHopSwapFails(types.ErrExitLimitPriceHit, routes, 100, sdk.MustNewDecFromStr("0.9"), true)

	// and with findFirstRoute

	s.aliceMultiHopSwapFails(types.ErrInsufficientLiquidity, routes, 100, sdk.MustNewDecFromStr("0.9"), false)
}

func (s *MsgServerTestSuite) TestMultiHopSwapMultiRouteFindBestRoute() {
	s.fundAliceBalances(100, 0)

	// GIVEN viable liquidity in pools but with a best route through E<>X
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenX", 0, 1000, -1000, 1),
		NewPoolSetup("TokenB", "TokenD", 0, 100, 0, 1),
		NewPoolSetup("TokenD", "TokenX", 0, 1000, -2000, 1),
		NewPoolSetup("TokenB", "TokenE", 0, 100, 0, 1),
		NewPoolSetup("TokenE", "TokenX", 0, 1000, -3000, 1),
	)

	// WHEN alice multihopswaps with three routes
	routes := [][]string{
		{"TokenA", "TokenB", "TokenC", "TokenX"},
		{"TokenA", "TokenB", "TokenD", "TokenX"},
		{"TokenA", "TokenB", "TokenE", "TokenX"},
	}
	s.aliceMultiHopSwaps(routes, 100, sdk.MustNewDecFromStr("0.9"), true)

	// THEN swap succeeds through route A<>B, B<>E, E<>X

	s.assertAccountBalanceExotic(s.alice, "TokenA", 0)
	s.assertAccountBalanceExotic(s.alice, "TokenX", 134)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenA", Token1: "TokenB"}, sdk.NewInt(100), sdk.NewInt(1), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenE"}, sdk.NewInt(100), sdk.NewInt(1), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenE", Token1: "TokenX"}, sdk.NewInt(100), sdk.NewInt(866), -3000, 1)

	// Other pools are unaffected
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenC"}, sdk.NewInt(0), sdk.NewInt(100), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenC", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(1000), -1000, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenB", Token1: "TokenD"}, sdk.NewInt(0), sdk.NewInt(100), 0, 1)
	s.assertLiquidityAtTickExotic(&types.PairID{Token0: "TokenD", Token1: "TokenX"}, sdk.NewInt(0), sdk.NewInt(1000), -2000, 1)
}
