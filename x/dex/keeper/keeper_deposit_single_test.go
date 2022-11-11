package keeper_test

import (
	// stdlib

	// cosmos SDK
	"testing"

	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSingleDeposit() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	tests := map[string]struct {
		// acc sdk.AccAddress
		denomA        string
		denomB        string
		amountsA      []sdk.Dec
		amountsB      []sdk.Dec
		tickIndexes   []int64
		feeTiers      []uint64
		expectedTxErr error
	}{
		"MinFee":     {denomA, denomB, []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}, []int64{0}, []uint64{0}, nil},
		"MaxFee":     {denomA, denomB, []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}, []int64{0}, []uint64{uint64(len(s.feeTiers)) - 1}, nil},
		"InvalidFee": {denomA, denomB, []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}, []int64{0}, []uint64{uint64(len(s.feeTiers))}, types.ErrValidFeeIndexNotFound},
	}
	for name, tc := range tests {
		success := s.T().Run(name, func(t *testing.T) {
			s.fundAliceBalances(10, 10)
			DepositTemplate(s, tc.denomA, tc.denomB, tc.amountsA, tc.amountsB, acc, tc.tickIndexes, tc.feeTiers, tc.expectedTxErr)
		})
		s.Require().True(success == (tc.expectedTxErr != nil))
	}
}

// func (s *MsgServerTestSuite) TestSingleMinFeeTier() {
// 	// prep deposit args
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit with min fee tier: 0
// 	tickIndex := []int64{0}
// 	minFeeTier := []uint64{uint64(s.feeTiers[0].Fee)}

// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, minFeeTier, nil)
// }

// func (s *MsgServerTestSuite) TestSingleMaxFeeTier() {
// 	// prep deposit args
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit with max fee tier
// 	tickIndex := []int64{0}
// 	maxFeeTier := []uint64{uint64(len(s.feeTiers) - 1)}

// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, maxFeeTier, nil)
// 	// validity assertions are done inside env.TestDeposit
// }

// func (s *MsgServerTestSuite) TestSingleInvalidFeeTier() {
// 	// prep deposit args
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
// 	tickIndex := []int64{0}
// 	invalidFeeTier := []uint64{uint64(len(s.feeTiers))}

// 	// (in)validity assertions are done inside env.TestDeposit
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, invalidFeeTier, types.ErrValidFeeIndexNotFound)
// }

// func (s *MsgServerTestSuite) TestSingleInitPair() {
// 	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitPair")

// 	// prep deposit args
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit at tick 0 in fee tier 0
// 	tickIndex := []int64{0}
// 	feeTier := []uint64{0}

// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier, nil)
// 	// validity assertions are done inside env.TestDeposit
// }

// func (s *MsgServerTestSuite) TestSingleInitTick() {
// 	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitTick")

// 	// prep deposit args
// 	acc := env.addrs[0]
// 	// fifth of acc's balance of each coin
// 	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
// 	denomA, denomB := coinA.Denom, coinB.Denom
// 	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit at tick 0 in fee tier 0
// 	tickIndex := []int64{0}
// 	feeTier := []uint64{0}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

// 	newTickIndex := []int64{1}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, newTickIndex, feeTier)
// 	// validity assertions are done inside env.TestDeposit
// }

// func (s *MsgServerTestSuite) TestSingleInitFeeTier() {
// 	// fifth of acc's balance of each coin
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}

// 	// deposit at tick 0 in fee tier 0
// 	tickIndex := []int64{0}
// 	feeTier := []uint64{0}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)
// 	// validity assertions are done inside env.TestDeposit

// 	newFeeTier := []uint64{1}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, newFeeTier)
// 	// validity assertions are done inside env.TestDeposit

// }

// func (s *MsgServerTestSuite) TestSingleExistingPair() {
// 	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/ExistingPair")

// 	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA)}, []sdk.Dec{s.getBalance(acc, denomB)}
// 	acc := env.addrs[0]
// 	// fifth of acc's balance of each coin
// 	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
// 	denomA, denomB := coinA.Denom, coinB.Denom
// 	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

// 	// deposit at tick 0 in fee tier 0
// 	tickIndex := []int64{0}
// 	feeTier := []uint64{0}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

// 	// WHEN deposit in the same pair, tick and fee tier again
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

// 	// THEN the transaction should execute successfully
// }

// func (s *MsgServerTestSuite) TestSingleBehindEnemyLines() {
// 	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/BehindEnemyLines")

// 	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
// 	acc := s.alice
// 	s.fundAccountBalances(acc, 10, 10)
// 	denomA, denomB := "TokenA", "TokenB"
// 	// fifth of acc's balance of each coin
// 	amountsA, amountsB := []sdk.Dec{s.getBalance(acc, denomA).Quo(5)}, []sdk.Dec{s.getBalance(acc, denomB).Quo(5)}

// 	// deposit at tick 0 in fee tier 0
// 	tickIndex := []int64{0}
// 	feeTier := []uint64{0}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

// 	// WHEN alice deposits at tick 0 in fee tier 1
// 	newTickIndex := []int64{-3}
// 	newFeeTier := []uint64{1}
// 	s.Deposit(denomA, denomB, amountsA, amountsB, acc, newTickIndex, newFeeTier)

// 	// THEN the transaction should execute successfully
// 	// validity assertions are done inside env.TestDeposit
// }
