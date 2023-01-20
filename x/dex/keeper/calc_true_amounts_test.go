package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/keeper"
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
		sdk.NewInt(10),
		sdk.NewInt(40), // value 20
		sdk.NewInt(100),
		sdk.NewInt(100), // effectively (25, 100), value 50
		sdk.NewInt(50),
	)
	s.Assert().Equal(sdk.NewInt(25), trueAmount0)
	s.Assert().Equal(sdk.NewInt(100), trueAmount1)
	s.Assert().Equal(sdk.NewInt(125), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestBothReservesZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(100),
		sdk.NewInt(100),
		sdk.NewInt(50),
	)
	s.Assert().Equal(sdk.NewInt(100), trueAmount0)
	s.Assert().Equal(sdk.NewInt(100), trueAmount1)
	s.Assert().Equal(sdk.NewInt(125), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestWrongCoinDeposited() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		sdk.NewInt(100),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(100),
		sdk.NewInt(50),
	)
	s.Assert().Equal(sdk.NewInt(0), trueAmount0)
	s.Assert().Equal(sdk.NewInt(0), trueAmount1)
	s.Assert().Equal(sdk.NewInt(0), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		sdk.NewInt(0),
		sdk.NewInt(100),
		sdk.NewInt(100),
		sdk.NewInt(0),
		sdk.NewInt(50),
	)
	s.Assert().Equal(sdk.NewInt(0), trueAmount0)
	s.Assert().Equal(sdk.NewInt(0), trueAmount1)
	s.Assert().Equal(sdk.NewInt(0), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestOneReserveZero() {
	trueAmount0, trueAmount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		sdk.NewInt(100),
		sdk.NewInt(0), // value 100
		sdk.NewInt(100),
		sdk.NewInt(100), // effective (100, 0), value 100
		sdk.NewInt(50),  // value went down
	)
	s.Assert().Equal(sdk.NewInt(100), trueAmount0)
	s.Assert().Equal(sdk.NewInt(0), trueAmount1)
	s.Assert().Equal(sdk.NewInt(50), sharesMinted)

	trueAmount0, trueAmount1, sharesMinted = keeper.CalcTrueAmounts(
		sdk.MustNewDecFromStr("0.25"),
		sdk.NewInt(0),
		sdk.NewInt(100), // value 25
		sdk.NewInt(100),
		sdk.NewInt(100), // effective (0, 100), value 25
		sdk.NewInt(50),  // value went up
	)
	s.Assert().Equal(sdk.NewInt(0), trueAmount0)
	s.Assert().Equal(sdk.NewInt(100), trueAmount1)
	s.Assert().Equal(sdk.NewInt(50), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmountsEmptyPoolBothSides() {
	// WHEN deposit into an empty pool
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(1),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(50),
		sdk.NewInt(20),
		sdk.NewInt(0),
	)

	// THEN both amounts are used fully

	s.Assert().Equal(sdk.NewInt(50), amount0)
	s.Assert().Equal(sdk.NewInt(20), amount1)
	s.Assert().Equal(sdk.NewInt(70), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmountsEmptyPoolToken0() {
	// WHEN deposit only Token0 into an empty pool
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(2),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(50),
		sdk.NewInt(0),
		sdk.NewInt(0),
	)

	// THEN all of Token0 is used

	s.Assert().Equal(sdk.NewInt(50), amount0)
	s.Assert().Equal(sdk.NewInt(0), amount1)
	s.Assert().Equal(sdk.NewInt(50), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmountsEmptyPoolToken1() {
	// WHEN deposit only Token1 into an empty pool
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(2),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(50),
		sdk.NewInt(0),
	)

	// THEN all of Token1 is used

	s.Assert().Equal(sdk.NewInt(0), amount0)
	s.Assert().Equal(sdk.NewInt(50), amount1)
	s.Assert().Equal(sdk.NewInt(100), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesRightRatio() {
	// WHEN deposit into a pool with a ratio of 2:5 with the same ratio all of the tokens are used
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(2),
		sdk.NewInt(20),
		sdk.NewInt(50),
		sdk.NewInt(4),
		sdk.NewInt(10),
		sdk.NewInt(120),
	)

	// THEN both amounts are fully user

	s.Assert().Equal(sdk.NewInt(4), amount0)
	s.Assert().Equal(sdk.NewInt(10), amount1)
	s.Assert().Equal(sdk.NewInt(24), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio() {
	// WHEN deposit into a pool with a ratio of 3:2 with a ratio of 2:1
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(3),
		sdk.NewInt(2),
		sdk.NewInt(20),
		sdk.NewInt(10),
		sdk.NewInt(9),
	)

	// THEN all of Token1 is used and 3/4 of token0 is used

	s.Assert().Equal(sdk.NewInt(15), amount0)
	s.Assert().Equal(sdk.NewInt(10), amount1)
	s.Assert().Equal(sdk.NewInt(45), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio2() {
	// IF deposit into a pool with a ratio of 2:3 with a ratio of 1:2
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(2),
		sdk.NewInt(3),
		sdk.NewInt(10),
		sdk.NewInt(20),
		sdk.NewInt(11),
	)

	// THEN all of Token0 is used and 3/4 of token1 is used

	s.Assert().Equal(sdk.NewInt(10), amount0)
	s.Assert().Equal(sdk.NewInt(15), amount1)
	s.Assert().Equal(sdk.NewInt(55), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPoolBothSides() {
	// WHEN deposit Token0 and Token1 into a pool with only Token0
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(10),
		sdk.NewInt(10),
	)

	// THEN only Token0 is used

	s.Assert().Equal(sdk.NewInt(10), amount0)
	s.Assert().Equal(sdk.NewInt(0), amount1)
	s.Assert().Equal(sdk.NewInt(10), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPoolBothSides2() {
	// WHEN deposit Token0 and Token1 into a pool with only Token1
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(10),
		sdk.NewInt(10),
		sdk.NewInt(30),
	)

	// THEN only Token1 is used

	s.Assert().Equal(sdk.NewInt(0), amount0)
	s.Assert().Equal(sdk.NewInt(10), amount1)
	s.Assert().Equal(sdk.NewInt(30), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0() {
	// WHEN deposit Token0 into a pool with only Token1
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(10),
	)

	// THEN no amounts are used

	s.Assert().Equal(sdk.NewInt(0), amount0)
	s.Assert().Equal(sdk.NewInt(0), amount1)
	s.Assert().Equal(sdk.NewInt(0), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0B() {
	// WHEN deposit Token0 into a pool with only Token0
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(10),
	)

	// THEN all of Token0 is used

	s.Assert().Equal(sdk.NewInt(10), amount0)
	s.Assert().Equal(sdk.NewInt(0), amount1)
	s.Assert().Equal(sdk.NewInt(10), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1() {
	// WHEN deposit Token1 into a pool with only Token0
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(3),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(0),
		sdk.NewInt(1),
		sdk.NewInt(10),
	)

	// THEN no amounts are used

	s.Assert().Equal(sdk.NewInt(0), amount0)
	s.Assert().Equal(sdk.NewInt(0), amount1)
	s.Assert().Equal(sdk.NewInt(0), sharesMinted)
}

func (s *CalcTrueAmountsTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1B() {
	// WHEN deposit Token1 into a pool with only Token1
	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(
		sdk.NewDec(4),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(0),
		sdk.NewInt(10),
		sdk.NewInt(40),
	)

	// THEN all of Token1 is used

	s.Assert().Equal(sdk.NewInt(0), amount0)
	s.Assert().Equal(sdk.NewInt(10), amount1)
	s.Assert().Equal(sdk.NewInt(40), sharesMinted)
}

// Calc_price_0to1 ////////////////////////////////////////////////////////////

func (s *CalcTrueAmountsTestSuite) TestCalc_price_1to0() {
	price := keeper.MustCalcPrice1To0(0)
	expected, _ := sdk.NewDecFromStr("1.0")

	s.Assert().Equal(expected, price)

	price = keeper.MustCalcPrice1To0(1)
	expected, _ = sdk.NewDecFromStr("1.0001")

	s.Assert().Equal(expected, price)
}

// Calc_price_1to0 ////////////////////////////////////////////////////////////

func (s *CalcTrueAmountsTestSuite) TestCalc_price_0to1() {
	price := keeper.MustCalcPrice0To1(0)
	expected, _ := sdk.NewDecFromStr("1.0")

	s.Assert().Equal(expected, price)

	price = keeper.MustCalcPrice0To1(1)
	expected, _ = sdk.NewDecFromStr("0.9999000099990001")

	s.Assert().Equal(expected, price)
}
