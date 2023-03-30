package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dualityapp "github.com/duality-labs/duality/app"
	. "github.com/duality-labs/duality/x/dex/keeper"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
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

	s.app = app
	s.msgServer = NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
}

func (s *CoreHelpersTestSuite) setLPAtFee1Pool(tickIndex int64, amountA, amountB int) {
	pairID := &types.PairID{Token0: "TokenA", Token1: "TokenB"}
	sharesID := CreateSharesID("TokenA", "TokenB", tickIndex, 1)
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, pairID, tickIndex, 1)
	if err != nil {
		panic(err)
	}
	lowerTick, upperTick := pool.LowerTick0, pool.UpperTick1
	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))

	existingShares := s.app.BankKeeper.GetSupply(s.ctx, sharesID).Amount

	totalShares := pool.CalcSharesMinted(amountAInt, amountBInt, existingShares)

	s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesID)

	lowerTick.Reserves = amountAInt
	upperTick.Reserves = amountBInt
	s.app.DexKeeper.SavePool(s.ctx, pool)
}

// GetOrInitUserShareData /////////////////////////////////////////////////////

// TODO: WRITE ME

// FindNextTick ////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestFindNextTick1To0NoLiq() {
	// GIVEN there is no ticks with token0 in the pool

	s.setLPAtFee1Pool(1, 0, 10)

	// THEN GetCurrTick1To0 doesn't find a tick

	_, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairID)
	s.Assert().False(found)
}

func (s *CoreHelpersTestSuite) TestGetCurrTick1To0WithLiq() {
	// Given multiple locations of token0
	s.setLPAtFee1Pool(-1, 10, 0)
	s.setLPAtFee1Pool(0, 10, 0)

	// THEN GetCurrTick1To0 finds the tick at -1

	tickIdx, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairID)
	s.Require().True(found)
	s.Assert().Equal(int64(-1), tickIdx)
}

func (s *CoreHelpersTestSuite) TestGetCurrTick1To0WithMinLiq() {
	// GIVEN tick with token0 @ index -1
	s.setLPAtFee1Pool(-1, 10, 0)
	s.setLPAtFee1Pool(1, 0, 0)

	// THEN GetCurrTick1To0 finds the tick at -2

	tickIdx, found := s.app.DexKeeper.GetCurrTick1To0(s.ctx, defaultPairID)
	s.Require().True(found)
	s.Assert().Equal(int64(-2), tickIdx)
}

// GetCurrTick0To1 ///////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1NoLiq() {
	// GIVEN there are no tick with Token1 in the pool

	s.setLPAtFee1Pool(0, 10, 0)

	// THEN GetCurrTick0To1 doesn't find a tick

	_, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairID)
	s.Assert().False(found)
}

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1WithLiq() {
	// GIVEN multiple locations of token1

	s.setLPAtFee1Pool(-1, 10, 0)
	s.setLPAtFee1Pool(0, 0, 10)
	s.setLPAtFee1Pool(1, 0, 10)

	// THEN GetCurrTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairID)
	s.Require().True(found)
	s.Assert().Equal(int64(1), tickIdx)
}

func (s *CoreHelpersTestSuite) TestGetCurrTick0To1WithMinLiq() {
	// WHEN tick with token1 @ index 1
	s.setLPAtFee1Pool(-1, 0, 0)
	s.setLPAtFee1Pool(1, 0, 10)

	// THEN GetCurrTick0To1 finds the tick at 2

	tickIdx, found := s.app.DexKeeper.GetCurrTick0To1(s.ctx, defaultPairID)
	s.Require().True(found)
	s.Assert().Equal(int64(2), tickIdx)
}
