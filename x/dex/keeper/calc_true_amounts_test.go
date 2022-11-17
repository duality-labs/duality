package keeper_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/stretchr/testify/suite"
)

type CalcTrueAmountsTestSuite struct {
	suite.Suite
}

func TestCalcTrueAmountsTestSuite(t *testing.T) {
	suite.Run(t, new(CalcTrueAmountsTestSuite))
}

func (s *CalcTrueAmountsTestSuite) TestBothReservesNonZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("10"),
		newDec("40"), // value 20
		newDec("100"),
		newDec("100"), // effectively (25, 100), value 50
		newDec("50"),
	)
	s.Assert().Equal(newDec("25"), trueAmount0)
	s.Assert().Equal(newDec("100"), trueAmount1)
	s.Assert().Equal(newDec("125"), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestBothReservesZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("0"),
		newDec("0"),
		newDec("100"),
		newDec("100"),
		newDec("50"),
	)
	s.Assert().Equal(newDec("100"), trueAmount0)
	s.Assert().Equal(newDec("100"), trueAmount1)
	s.Assert().Equal(newDec("125"), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestWrongCoinDeposited() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("100"),
		newDec("0"),
		newDec("0"),
		newDec("100"),
		newDec("50"),
	)
	s.Assert().Equal(newDec("0"), trueAmount0)
	s.Assert().Equal(newDec("0"), trueAmount1)
	s.Assert().Equal(newDec("0"), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("0"),
		newDec("100"),
		newDec("100"),
		newDec("0"),
		newDec("50"),
	)
	s.Assert().Equal(newDec("0"), trueAmount0)
	s.Assert().Equal(newDec("0"), trueAmount1)
	s.Assert().Equal(newDec("0"), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestOneReserveZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("100"),
		newDec("0"), // value 100
		newDec("100"),
		newDec("100"), // effective (100, 0), value 100
		newDec("50"),  // value went down
	)
	s.Assert().Equal(newDec("100"), trueAmount0)
	s.Assert().Equal(newDec("0"), trueAmount1)
	s.Assert().Equal(newDec("50"), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		newDec("0.25"),
		newDec("0"),
		newDec("100"), // value 25
		newDec("100"),
		newDec("100"), // effective (0, 100), value 25
		newDec("50"),  // value went up
	)
	s.Assert().Equal(newDec("0"), trueAmount0)
	s.Assert().Equal(newDec("100"), trueAmount1)
	s.Assert().Equal(newDec("50"), sharesMinted)
}
