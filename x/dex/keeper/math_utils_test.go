package keeper_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type MathUtilsTestSuite struct {
	suite.Suite
}

func TestMathUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(MathUtilsTestSuite))
}

func (s *MathUtilsTestSuite) TestPow() {
	result, err := keeper.Pow(sdk.NewDec(2), 2)
	s.Assert().Nil(err)
	s.Assert().Equal(sdk.NewDec(4), result)

	result, err = keeper.Pow(keeper.BasePrice(), 2)
	s.Assert().Nil(err)
	s.Assert().Equal(sdk.MustNewDecFromStr("1.00020001"), result)

	result, err = keeper.Pow(keeper.BasePrice(), 2)
	s.Assert().Nil(err)
	s.Assert().Equal(sdk.MustNewDecFromStr("0.999800029996000500"), sdk.OneDec().Quo(result))
}

func (s *MathUtilsTestSuite) TestPowFailWithHighTick() {
	_, err := keeper.Pow(keeper.BasePrice(), keeper.MaxTickExp+1)
	s.Assert().ErrorIs(err, types.ErrTickAbsValTooHigh)

	_, err = keeper.Pow(keeper.BasePrice(), keeper.MaxTickExp)
	s.Assert().Nil(err)
}
