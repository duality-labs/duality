package keeper_test

import (
	"errors"

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
func (s *MsgServerTestSuite) TestSingleDepositInvalidExistingLiquidityOppositeSide() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(0)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
	DepositTemplate(s, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers, nil)

	amountsA2, amountsB2 := []sdk.Dec{NewDec(0)}, []sdk.Dec{NewDec(10)}
	// TODO: need to add error case here
	DepositTemplate(s, denomA, denomB, amountsA2, amountsB2, acc, tickIndexes, feeTiers, errors.New(""))
}

func (s *MsgServerTestSuite) TestSingleDepositOneSided() {
	acc := s.alice
	denomA, denomB := "TokenA", "TokenB"
	s.fundAliceBalances(10, 10)
	amountsA, amountsB := []sdk.Dec{NewDec(10)}, []sdk.Dec{NewDec(0)}
	tickIndexes := []int64{0}
	feeTiers := []uint64{0}
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
