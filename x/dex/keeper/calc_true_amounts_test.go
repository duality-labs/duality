package keeper_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		sdk.MustNewDecFromStr("0.25"),
		NewDec(10),
		NewDec(40), // value 20
		NewDec(100),
		NewDec(100), // effectively (25, 100), value 50
		NewDec(50),
	)
	s.Assert().Equal(NewDec(25), trueAmount0)
	s.Assert().Equal(NewDec(100), trueAmount1)
	s.Assert().Equal(NewDec(125), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestBothReservesZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		NewDec(0),
		NewDec(0),
		NewDec(100),
		NewDec(100),
		NewDec(50),
	)
	s.Assert().Equal(NewDec(100), trueAmount0)
	s.Assert().Equal(NewDec(100), trueAmount1)
	s.Assert().Equal(NewDec(125), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestWrongCoinDeposited() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		NewDec(100),
		NewDec(0),
		NewDec(0),
		NewDec(100),
		NewDec(50),
	)
	s.Assert().Equal(NewDec(0), trueAmount0)
	s.Assert().Equal(NewDec(0), trueAmount1)
	s.Assert().Equal(NewDec(0), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		NewDec(0),
		NewDec(100),
		NewDec(100),
		NewDec(0),
		NewDec(50),
	)
	s.Assert().Equal(NewDec(0), trueAmount0)
	s.Assert().Equal(NewDec(0), trueAmount1)
	s.Assert().Equal(NewDec(0), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestOneReserveZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		NewDec(100),
		NewDec(0), // value 100
		NewDec(100),
		NewDec(100), // effective (100, 0), value 100
		NewDec(50),  // value went down
	)
	s.Assert().Equal(NewDec(100), trueAmount0)
	s.Assert().Equal(NewDec(0), trueAmount1)
	s.Assert().Equal(NewDec(50), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		NewDec(0),
		NewDec(100), // value 25
		NewDec(100),
		NewDec(100), // effective (0, 100), value 25
		NewDec(50),  // value went up
	)
	s.Assert().Equal(NewDec(0), trueAmount0)
	s.Assert().Equal(NewDec(100), trueAmount1)
	s.Assert().Equal(NewDec(50), sharesMinted)
}
