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

func (s *MsgServerTestSuite) TestMultiHopSwap() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D,
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 100, 0, 1),
	)

	// WHEN alice multihopswaps A<>B => B<>C => C<>D,
	hops := []string{"TokenA", "TokenB", "TokenC", "TokenD"}
	s.aliceMultiHopSwaps(hops, 100, sdk.MustNewDecFromStr("0.9"))

	// THEN alice gets out 99 TokenD
	s.assertAccountBalanceExotic(s.alice, "TokenA", 0)
	s.assertAccountBalanceExotic(s.alice, "TokenD", 99)

	s.assertDexBalanceExotic("TokenA", 100)
	s.assertDexBalanceExotic("TokenB", 100)
	s.assertDexBalanceExotic("TokenC", 100)
	s.assertDexBalanceExotic("TokenD", 1)
}

func (s *MsgServerTestSuite) TestMultiHopSwapInsufficientLiquidity() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D with insufficient liquidity in C<>D
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 50, 0, 1),
	)

	// THEN alice multihopswap fails
	hops := []string{"TokenA", "TokenB", "TokenC", "TokenD"}
	s.aliceMultiHopSwapFails(types.ErrInsufficientLiquidity, hops, 100, sdk.MustNewDecFromStr("0.9"))
}

func (s *MsgServerTestSuite) TestMultiHopSwapLimitPriceNotMet() {
	s.fundAliceBalances(100, 0)

	// GIVEN liquidity in pools A<>B, B<>C, C<>D with insufficient liquidity in C<>D
	s.SetupMultiplePools(
		NewPoolSetup("TokenA", "TokenB", 0, 100, 0, 1),
		NewPoolSetup("TokenB", "TokenC", 0, 100, 0, 1),
		NewPoolSetup("TokenC", "TokenD", 0, 100, 1200, 1),
	)

	// THEN alice multihopswap fails
	hops := []string{"TokenA", "TokenB", "TokenC", "TokenD"}
	s.aliceMultiHopSwapFails(types.ErrExitLimitPriceHit, hops, 50, sdk.MustNewDecFromStr("0.9"))
}
