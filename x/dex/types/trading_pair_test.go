package types_test

import (
	"context"
	"testing"

	dualityapp "github.com/duality-labs/duality/app"
	. "github.com/duality-labs/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TradingPairTestSuite struct {
	suite.Suite
	app           *dualityapp.App
	ctx           sdk.Context
	goCtx         context.Context
	defaultPairId *PairId
}

func TestTradingPairTestSuite(t *testing.T) {
	suite.Run(t, new(TradingPairTestSuite))
}

func (s *TradingPairTestSuite) SetupTest() {
	s.app = dualityapp.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	feeTiers := []FeeTier{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[0])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[1])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[2])
	s.app.DexKeeper.AppendFeeTier(s.ctx, feeTiers[3])
	s.defaultPairId = &PairId{Token0: "TokenA", Token1: "TokenB"}
}

func (s *TradingPairTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) (lowerTick Tick, upperTick Tick) {

	lowerTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, s.defaultPairId, tickIndex-1)
	s.Assert().NoError(err)
	upperTick, err = s.app.DexKeeper.GetOrInitTick(s.goCtx, s.defaultPairId, tickIndex+1)
	s.Assert().NoError(err)
	// priceCenter1To0, err := keeper.CalcPrice0To1(tickIndex)
	// if err != nil {
	// 	panic(err)
	// }

	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))
	lowerTick.TickData.Reserve0[0] = amountAInt
	// totalShares := keeper.CalcShares(amountAInt, amountBInt, priceCenter1To0).TruncateInt()
	// s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesId)
	upperTick.TickData.Reserve1[0] = amountBInt
	s.app.DexKeeper.SetTick(s.ctx, s.defaultPairId, lowerTick)
	s.app.DexKeeper.SetTick(s.ctx, s.defaultPairId, upperTick)
	return lowerTick, upperTick
}

// PairIdToTokens ///////////////////////////////////////////////////////////////

func (s *TradingPairTestSuite) TestPairToTokens() {

	token0, token1 := PairIdToTokens(&PairId{Token0: "TokenA", Token1: "TokenB"})

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("TokenB", token1)

}

func (s *TradingPairTestSuite) TestPairToTokensIBCis0() {

	token0, token1 := PairIdToTokens(&PairId{Token0: "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", Token1: "TokenB"})

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("TokenB", token1)
}

func (s *TradingPairTestSuite) TestPairToTokensIBCis1() {

	token0, token1 := PairIdToTokens(&PairId{Token0: "TokenA", Token1: "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"})

	s.Assert().Equal("TokenA", token0)
	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token1)

}
func (s *TradingPairTestSuite) TestPairToTokensIBCisBoth() {

	token0, token1 :=
		PairIdToTokens(&PairId{
			Token0: "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
			Token1: "ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223",
		})

	s.Assert().Equal("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", token0)
	s.Assert().Equal("ibc/94644FB092D9ACDA56123C74F36E4234926001AA44A9CA97EA622B25F41E5223", token1)
}
