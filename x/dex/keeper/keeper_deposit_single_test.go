package keeper_test

import (
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestSingleDepositMinFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

}

func (s *MsgServerTestSuite) TestSingleDepositMaxFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{uint64(len(s.feeTiers) - 1)}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

}

func (s *MsgServerTestSuite) TestSingleDepositInvalidFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{uint64(len(s.feeTiers))}
	expectedTxErr := types.ErrValidFeeIndexNotFound
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, expectedTxErr)

}

func (s *MsgServerTestSuite) TestSingleDepositInitPair() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(10)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

}

func (s *MsgServerTestSuite) TestSingleDepositInitTick() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5)}, []sdk.Dec{NewDec(5)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

	newTickIndexes := []int64{1}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, newTickIndexes, feeTiers, nil)

}

func (s *MsgServerTestSuite) TestSingleDepositInitFeeTier() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5)}, []sdk.Dec{NewDec(5)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

	newFeeTiers := []uint64{1}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, newFeeTiers, nil)

}

func (s *MsgServerTestSuite) TestSingleDepositExistingPair() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5)}, []sdk.Dec{NewDec(5)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestSingleDepositBehindEnemyLines() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5)}, []sdk.Dec{NewDec(5)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
	newTickIndexes := []int64{-3}
	newFeeTiers := []uint64{1}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, newTickIndexes, newFeeTiers, nil)

}

// table driven test. can't work with testify because setup isn't run before each s.T().Run()
/* func (s *MsgServerTestSuite) TestSingleDepositExistingLiquidity() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	tests := map[string]struct {
		// acc sdk.AccAddress
		initialBalanceA int64
		initialBalanceB int64
		denomA          string
		denomB          string
		amountsA        []sdk.Dec
		amountsB        []sdk.Dec
		tickIndexes     []int64
		feeTiers        []uint64
		expectedTxErr   []error
	}{
		"InitTick":         {10, 10, denomA, denomB, []sdk.Dec{NewDec(5), NewDec(3)}, []sdk.Dec{NewDec(5), NewDec(3)}, []int64{0, 1}, []uint64{0, 0}, []error{nil, nil}},
		"InitFeeTier":      {10, 10, denomA, denomB, []sdk.Dec{NewDec(5), NewDec(3)}, []sdk.Dec{NewDec(5), NewDec(3)}, []int64{0, 0}, []uint64{0, 1}, []error{nil, nil}},
		"ExistingPair":     {10, 10, denomA, denomB, []sdk.Dec{NewDec(5), NewDec(3)}, []sdk.Dec{NewDec(5), NewDec(3)}, []int64{0, 0}, []uint64{0, 0}, []error{nil, nil}},
		"BehindEnemyLines": {initialBalanceA: 10, initialBalanceB: 10, denomA: denomA, denomB: denomB, amountsA: []sdk.Dec{NewDec(5), NewDec(3)}, amountsB: []sdk.Dec{NewDec(5), NewDec(3)}, tickIndexes: []int64{0, -3}, feeTiers: []uint64{0, 1}, expectedTxErr: []error{nil, nil}},
	}
	for name, tc := range tests {
		s.T().Log(name)
		s.T().Run(name, func(t *testing.T) {
			s.fundAliceBalances(int(tc.initialBalanceA), int(tc.initialBalanceB))
			DepositTemplate(s, tc.denomA, tc.denomB, tc.amountsA[:1], tc.amountsB[:1], acc, tc.tickIndexes[:1], tc.feeTiers[:1], tc.expectedTxErr[0])
			DepositTemplate(s, tc.denomA, tc.denomB, tc.amountsA[1:], tc.amountsB[1:], acc, tc.tickIndexes[1:], tc.feeTiers[1:], tc.expectedTxErr[1])
		})
	}
} */
