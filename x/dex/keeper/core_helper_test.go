package keeper_test

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dualityapp "github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/x/dex/keeper"
	. "github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	feeCount = 4
)

// Test Suite ///////////////////////////////////////////////////////////////
type CoreHelpersTestSuite struct {
	suite.Suite
	app         *dualityapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
	alice       sdk.AccAddress
	bob         sdk.AccAddress
	carol       sdk.AccAddress
	dan         sdk.AccAddress
	goCtx       context.Context
	feeTiers    []types.FeeTier
}

func TestCoreHelpersTestSuite(t *testing.T) {
	suite.Run(t, new(CoreHelpersTestSuite))
}

func (s *CoreHelpersTestSuite) SetupTest() {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DexKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, s.alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, s.bob)
	app.AccountKeeper.SetAccount(ctx, accBob)
	accCarol := app.AccountKeeper.NewAccountWithAddress(ctx, s.carol)
	app.AccountKeeper.SetAccount(ctx, accCarol)
	accDan := app.AccountKeeper.NewAccountWithAddress(ctx, s.dan)
	app.AccountKeeper.SetAccount(ctx, accDan)

	// add the fee tiers of 1, 3, 5, 10 ticks
	feeTiers := []types.FeeTier{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[0])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[1])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[2])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[3])

	s.app = app
	s.msgServer = keeper.NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
	s.feeTiers = feeTiers
}

func (s *CoreHelpersTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) {
	pairId := &types.PairId{"TokenA", "TokenB"}
	sharesId := CreateSharesId("TokenA", "TokenB", tickIndex, 0)
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, pairId, tickIndex, s.feeTiers[0])

	if err != nil {
		panic(err)
	}

	lowerTick, upperTick := pool.LowerTick0, pool.UpperTick1
	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))
	totalShares := pool.CalcSharesMinted(amountAInt, amountBInt)

	s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesId)

	lowerTick.Reserves = amountAInt
	upperTick.Reserves = amountBInt
	pool.Save(s.ctx, s.app.DexKeeper)
}

// TokenInit //////////////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestTokenInitNew() {

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")

	tokenMap, found := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")

	s.Assert().True(found)
	s.Assert().Equal("TokenA", tokenMap.Address)
	s.Assert().Equal(uint64(1), s.app.DexKeeper.GetTokensCount(s.ctx))
}

func (s *CoreHelpersTestSuite) TestTokenInitExisting() {

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")
	s.Require().Equal(uint64(1), s.app.DexKeeper.GetTokensCount(s.ctx))

	s.app.DexKeeper.TokenInit(s.ctx, "TokenA")
}

// GetOrInitUserShareData /////////////////////////////////////////////////////

// TODO: WRITE ME

// FindNextTick ////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestFindNextTick1To0NoLiq() {
	// GIVEN there is no ticks with token0 in the pool

	s.setLPAtFee0Pool(1, 0, 10)

	// THEN GetCurrTick1To0 doesn't find a tick

	_, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairId)
	s.Assert().False(found)

}

func (s *CoreHelpersTestSuite) TestGetCurrTick1To0WithLiq() {
	// Given multiple locations of token0
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 10, 0)

	// THEN GetCurrTick1To0 finds the tick at -1

	tickIdx, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairId)
	s.Require().True(found)
	s.Assert().Equal(int64(-1), tickIdx)

}

func (s *CoreHelpersTestSuite) TestGetCurrTick1To0WithMinLiq() {
	// GIVEN tick with token0 @ index -1
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// THEN GetCurrTick1To0 finds the tick at -2

	tickIdx, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairId)
	s.Require().True(found)
	s.Assert().Equal(int64(-2), tickIdx)

}

// GetCurrTick0To1 ///////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1NoLiq() {
	// GIVEN there are no tick with Token1 in the pool

	s.setLPAtFee0Pool(0, 10, 0)

	// THEN GetCurrTick0To1 doesn't find a tick

	_, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairId)
	s.Assert().False(found)

}

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1WithLiq() {
	// GIVEN multiple locations of token1

	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 0, 10)
	s.setLPAtFee0Pool(1, 0, 10)

	// THEN GetCurrTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairId)
	s.Require().True(found)
	s.Assert().Equal(int64(1), tickIdx)
}

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1WithMinLiq() {
	// WHEN tick with token1 @ index 1
	s.setLPAtFee0Pool(-1, 0, 0)
	s.setLPAtFee0Pool(1, 0, 10)

	// THEN GetCurrTick0To1 finds the tick at 2

	tickIdx, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairId)
	s.Require().True(found)
	s.Assert().Equal(int64(2), tickIdx)
}
