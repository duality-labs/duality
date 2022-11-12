package keeper_test

import (
	"errors"

	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *MsgServerTestSuite) TestMultiDepositMinFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 1}
	feeTiers := []uint64{0, 0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositMaxFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 1}
	feeTiers := []uint64{uint64(len(s.feeTiers) - 1), uint64(len(s.feeTiers) - 1)}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositInvalidFee() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 1}
	feeTiers := []uint64{uint64(len(s.feeTiers)), uint64(len(s.feeTiers))}
	expectedTxErr := types.ErrValidFeeIndexNotFound
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, expectedTxErr)
}

func (s *MsgServerTestSuite) TestMultiDepositInitTick() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 1}
	feeTiers := []uint64{0, 0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositInitFeeTier() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 1}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositExistingPair() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositBehindEnemyLines() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(5)}, []sdk.Dec{NewDec(5), NewDec(5)}
	tickIndexes := []int64{0, -3}
	feeTiers := []uint64{0, 1}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)
}

func (s *MsgServerTestSuite) TestMultiDepositInvalidOneSidedDeposit() {
	// one sided deposit, then attempt to deposit on the other side
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(5), NewDec(0)}, []sdk.Dec{NewDec(0), NewDec(5)}
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 1}
	// TODO: need to add error for this
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, errors.New(""))
}
