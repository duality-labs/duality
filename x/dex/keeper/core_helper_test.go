package keeper_test

import (
	"context"
	"math"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

func (s *CoreHelpersTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) (lowerTick types.Tick, upperTick types.Tick) {
	// pairId := "TokenA<>TokenB"
	pairId := &types.PairId{"TokenA", "TokenB"}
	// sharesId := fmt.Sprintf("%s%st%df%d", "TokenA", "TokenB", tickIndex, 1)
	sharesId := CreateSharesId("TokenA", "TokenB", tickIndex, 0)
	lowerTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex-1)
	s.Assert().NoError(err)
	upperTick, err = s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex+1)
	s.Assert().NoError(err)
	priceCenter1To0, err := keeper.CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}

	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))
	lowerTick.TickData.Reserve0[0] = amountAInt
	totalShares := keeper.CalcShares(amountAInt, amountBInt, priceCenter1To0).TruncateInt()
	s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesId)
	upperTick.TickData.Reserve1[0] = amountBInt
	s.app.DexKeeper.SetTick(s.ctx, pairId, lowerTick)
	s.app.DexKeeper.SetTick(s.ctx, pairId, upperTick)
	return lowerTick, upperTick
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

// GetOrInitPair //////////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestGetOrInitPairNew() {
	// GIVEN we initialize a new pair
	s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// THEN the component tokens are also initialized...
	_, found0 := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")
	_, found1 := s.app.DexKeeper.GetTokenMap(s.ctx, "TokenA")

	s.Assert().True(found0)
	s.Assert().True(found1)

	// AND 1 pair is initialized with the correct default values
	pairCount := len(s.app.DexKeeper.GetAllTradingPair(s.ctx))
	s.Assert().Equal(1, pairCount)

	pair, foundPair := s.app.DexKeeper.GetTradingPair(s.ctx, defaultPairId)

	s.Require().True(foundPair)

	s.Assert().Equal(pair.PairId, &types.PairId{Token0: "TokenA", Token1: "TokenB"})
	s.Assert().Equal(int64(math.MaxInt64), pair.CurrentTick0To1)
	s.Assert().Equal(int64(math.MinInt64), pair.CurrentTick1To0)
	s.Assert().Equal(int64(math.MaxInt64), pair.MinTick)
	s.Assert().Equal(int64(math.MinInt64), pair.MaxTick)
}

func (s *CoreHelpersTestSuite) TestGetOrInitPairExisting() {

	// GIVEN we initialize a pair TokenA/TokenB
	s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// WHEN we update values on that pair
	pair, _ := s.app.DexKeeper.GetTradingPair(s.ctx, defaultPairId)
	pair.MinTick = 20
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// AND try to initialize the same pair again
	newPair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")

	// THEN there is still only 1 pair and it retains the values we set
	pairCount := len(s.app.DexKeeper.GetAllTradingPair(s.ctx))
	s.Assert().Equal(1, pairCount)
	s.Assert().Equal(int64(20), newPair.MinTick)
}

// GetOrInitTick //////////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestGetOrInitTickNew() {
	// GIVEN we initialize a new tick
	s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)

	// THEN 1 tick is initialized with the correct values

	tickCount := len(s.app.DexKeeper.GetAllTick(s.ctx))
	s.Assert().Equal(1, tickCount)

	tick, found := s.app.DexKeeper.GetTick(s.ctx, defaultPairId, 0)

	s.Require().True(found)

	s.Assert().Equal(tick.PairId, defaultPairId)
	s.Assert().Equal(tick.TickIndex, int64(0))
	s.Assert().Equal(feeCount, len(tick.TickData.Reserve0))
	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve0[0],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve0[feeCount-1],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve1[0],
	)

	s.Assert().Equal(
		sdk.ZeroInt(),
		tick.TickData.Reserve1[feeCount-1],
	)

	//AND tranche fill maps are initialized
	_, fill0Found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenA", 0)
	_, fill1Found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenB", 0)

	s.Assert().True(fill0Found)
	s.Assert().True(fill1Found)
}

func (s *CoreHelpersTestSuite) TestGetOrInitTickExisting() {

	// GIVEN we initialize a tick
	s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)

	// WHEN we update values on that tick
	tick, _ := s.app.DexKeeper.GetTick(s.ctx, defaultPairId, 0)
	tick.TickData.Reserve0[0] = sdk.NewInt(10)
	s.app.DexKeeper.SetTick(s.ctx, defaultPairId, tick)

	// AND try to initialize the same tick again
	newTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)

	// THEN there is still only 1 tick and it retains the values we set
	tickCount := len(s.app.DexKeeper.GetAllTick(s.ctx))
	s.Assert().Equal(1, tickCount)
	s.Assert().Equal(sdk.NewInt(10), newTick.TickData.Reserve0[0])
}

// GetOrInitTickTrancheFillMap ////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitReserveData ///////////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitUserShareData /////////////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitOrderPoolTotalShares //////////////////////////////////////////////

// TODO: WRITE ME

// GetOrInitLimitOrderMaps ////////////////////////////////////////////////////

// TODO: WRITE ME

// FindNextTick1To0 ///////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestFindNextTick1To0NoLiq() {
	// GIVEN there is no ticks with token0 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(1, 0, 10)

	// THEN FindNextTick1To0 doesn't find a tick

	_, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Assert().False(found)

}

func (s *CoreHelpersTestSuite) TestFindNextTick1To0WithLiq() {
	// GIVEN tick with token0 @ index 0 & currentTick0To1 is 1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// tick -2: (10, 0)
	// tick -1: (10, 0)
	pair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick1To0 finds the tick at 0

	tickIdx, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(-1), tickIdx)

}

func (s *CoreHelpersTestSuite) TestFindNextTick1To0WithMinLiq() {
	// GIVEN tick with token0 @ index -1 & MinTick = -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(1, 0, 0)

	// tick -2: (10, 0)
	pair.CurrentTick1To0 = 1
	pair.MinTick = -2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick1To0 finds the tick at -1

	tickIdx, found := s.app.DexKeeper.FindNextTick1To0(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(-2), tickIdx)

}

// FindNextTick0To1 ///////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestFindNextTick0To1NoLiq() {
	// GIVEN there are no tick with Token1 in the pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	// NOTE: Actually adding/funding a tick isn't really neccesary but should make the test less
	// dependent upon current implementation details
	s.setLPAtFee0Pool(0, 10, 0)

	// THEN FindNextTick0To1 doesn't find a tick

	_, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Assert().False(found)

}

func (s *CoreHelpersTestSuite) TestFindNextTick0To1WithLiq() {
	// WHEN tick with token1 @ index 0 & currentTick0To1 is -1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(0, 0, 10)
	s.setLPAtFee0Pool(1, 0, 10)

	// tick -2: (10,  0)
	// tick -1: ( 0,  0)
	// tick  0: ( 0,  0)
	// tick  1: ( 0, 10)
	// tick  2: ( 0, 10)

	pair.CurrentTick0To1 = -1
	pair.MaxTick = 1
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(1), tickIdx)
}

func (s *CoreHelpersTestSuite) TestFindNextTick0To1WithMinLiq() {
	// WHEN tick with token1 @ index 1 & MaxTick = 1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 0, 0)
	s.setLPAtFee0Pool(1, 0, 10)

	// tick 2: (0, 10)
	pair.CurrentTick0To1 = -1
	pair.MaxTick = 2
	s.app.DexKeeper.SetTradingPair(s.ctx, pair)

	// THEN FindNextTick0To1 finds the tick at 1

	tickIdx, found := s.app.DexKeeper.FindNextTick0To1(s.goCtx, pair)
	s.Require().True(found)
	s.Assert().Equal(int64(2), tickIdx)
}

// CalcTrueAmounts ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolBothSides(){
// 	// WHEN deposit into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(1), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50), sdk.NewDec(20))

// 	// THEN both amounts are used fully

// 	s.Assert().Equal(sdk.NewDec(50), amount0)
// 	s.Assert().Equal(sdk.NewDec(20), amount1)
// 	s.Assert().Equal(sdk.NewDec(70), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolToken0(){
// 	// WHEN deposit only Token0 into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50), sdk.NewDec(0))

// 	// THEN all of Token0 is used

// 	s.Assert().Equal(sdk.NewDec(50), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(50), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmountsEmptyPoolToken1(){
// 	// WHEN deposit only Token1 into an empty pool
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(50))

// 	// THEN all of Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(50), amount1)
// 	s.Assert().Equal(sdk.NewDec(100), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesRightRatio(){
// 	// WHEN deposit into a pool with a ratio of 2:5 with the same ratio all of the tokens are used
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(2), sdk.NewDec(20), sdk.NewDec(50), sdk.NewDec(4), sdk.NewDec(10))

// 	// THEN both amounts are fully user

// 	s.Assert().Equal(sdk.NewDec(4), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(24), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio(){
// 	// WHEN deposit into a pool with a ratio of 3:2 with a ratio of 2:1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(3), sdk.NewDec(2), sdk.NewDec(20), sdk.NewDec(10))

// 	// THEN all of Token1 is used and 3/4 of token0 is used

// 	s.Assert().Equal(sdk.NewDec(15), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(45), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts2SidedPoolBothSidesWrongRatio2(){
// 	// IF deposit into a pool with a ratio of 2:3 with a ratio of 1:2
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(2), sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(20))

// 	// THEN all of Token0 is used and 3/4 of token1 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(15), amount1)
// 	s.Assert().Equal(sdk.NewDec(55), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPoolBothSides(){
// 	// WHEN deposit Token0 and Token1 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10))

// 	// THEN only Token0 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(10), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPoolBothSides2(){
// 	// WHEN deposit Token0 and Token1 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10), sdk.NewDec(10))

// 	// THEN only Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(30), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0(){
// 	// WHEN deposit Token0 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(10), sdk.NewDec(0))

// 	// THEN no amounts are used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(0), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken0B(){
// 	// WHEN deposit Token0 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(0))

// 	// THEN all of Token0 is used

// 	s.Assert().Equal(sdk.NewDec(10), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(10), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1(){
// 	// WHEN deposit Token1 into a pool with only Token0
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(3), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(1))

// 	// THEN no amounts are used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(0), amount1)
// 	s.Assert().Equal(sdk.NewDec(0), sharesMinted)
// }

// func (s *MsgServerTestSuite) TestCalcTrueAmounts1SidedPool1SidedToken1B(){
// 	// WHEN deposit Token1 into a pool with only Token1
// 	amount0, amount1, sharesMinted := keeper.CalcTrueAmounts(sdk.NewDec(4), sdk.NewDec(0), sdk.NewDec(10), sdk.NewDec(0), sdk.NewDec(10))

// 	// THEN all of Token1 is used

// 	s.Assert().Equal(sdk.NewDec(0), amount0)
// 	s.Assert().Equal(sdk.NewDec(10), amount1)
// 	s.Assert().Equal(sdk.NewDec(40), sharesMinted)
// }

// Calc_price_0to1 ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalc_price_1to0(){
// 	price := s.app.DexKeeper.Calc_price_1to0(0)
// 	expected, _ := sdk.NewDecFromStr("1.0")

// 	s.Assert().Equal(expected, price)

// 	price= s.app.DexKeeper.Calc_price_1to0(1)
// 	expected, _ = sdk.NewDecFromStr("1.0001")

// 	s.Assert().Equal(expected, price)

// 	// TODO: ADD MORE ITAMAR
// }

// Calc_price_1to0 ////////////////////////////////////////////////////////////

// func (s *MsgServerTestSuite) TestCalc_price_0to1(){
// 	price := s.app.DexKeeper.Calc_price_0to1(0)
// 	expected, _ := sdk.NewDecFromStr("1.0")

// 	s.Assert().Equal(expected, price)

// 	price = s.app.DexKeeper.Calc_price_0to1(1)
// 	expected, _ = sdk.NewDecFromStr("0.9999000099990001")

// 	s.Assert().Equal(expected, price)
// 	// TODO: ADD MORE ITAMAR
// }

// HasToken0 //////////////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestHasToken0Empty() {

	// WHEN tick only has limit orders and reserves of Token1
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve1[0] = sdk.NewInt(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *CoreHelpersTestSuite) TestHasToken0HasReserves() {

	// WHEN tick has Reserves0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve0[3] = sdk.NewInt(10)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

func (s *CoreHelpersTestSuite) TestHasToken0HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken0(s.ctx, &tick))
}

// HasToken1 //////////////////////////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestHasToken1Empty() {

	// WHEN tick only has limit orders and reserves of Token0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve0[0] = sdk.NewInt(10)

	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenA", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().False(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *CoreHelpersTestSuite) TestHasToken1HasReserves() {

	// WHEN tick has Reserves0
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.TickData.Reserve1[0] = sdk.NewInt(10)

	// THEN HasToken1() = true
	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

func (s *CoreHelpersTestSuite) TestHasToken1HasLimitOrders() {

	// WHEN there are limit orders at the tick
	tick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, defaultPairId, 0)
	s.Assert().NoError(err)
	tick.LimitOrderTranche0To1.FillTrancheIndex = 0
	tranche := s.app.DexKeeper.GetOrInitLimitOrderTranche(s.ctx, defaultPairId, 0, "TokenB", 0)
	tranche.ReservesTokenIn = sdk.NewInt(100)
	s.app.DexKeeper.SetLimitOrderTranche(s.ctx, tranche)

	s.Assert().True(s.app.DexKeeper.TickHasToken1(s.ctx, &tick))
}

// CalcTickPointersPostAddToken0 //////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken0NoToken() {
	// GIVEN current tick still has 0 Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken0NoLiq() {
	// GIVEN minTick == MaxInt64 ie. no liquidity of Token0 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(2, 10, 0)
	// THEN MinTick and cur1To0 are set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(1), pair.CurrentTick1To0)
	s.Assert().Equal(int64(1), pair.MinTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken0New1To0() {
	// GIVEN current tick provides Token0 at a higher price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 1
	lower, _ := s.setLPAtFee0Pool(3, 10, 0)
	// THEN curTick1To0 is set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), pair.MinTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken0NewMinTick() {
	// GIVEN current tick provides Token0 at a new lowest price (tick -12)
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(-11, 10, 0)
	// THEN MinTick is set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-12), pair.MinTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken0NoChange() {
	// GIVEN current tick provides Token0 below 1To0 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -10
	pair.CurrentTick1To0 = 2
	lower, _ := s.setLPAtFee0Pool(1, 10, 0)
	// THEN no changes are made to MinTick & Cur1To0
	s.app.DexKeeper.UpdateTickPointersPostAddToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(int64(2), pair.CurrentTick1To0)
	s.Assert().Equal(int64(-10), pair.MinTick)
}

// CalcTickPointersPostAddToken1 //////////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken1NoToken() {
	// GIVEN current tick still has NO Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostAddToken1(s.goCtx, &pair, &lower)
	s.Assert().Equal(pair, pair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken1NoLiq() {
	// GIVEN maxTick == MinInt64 ie. no liquidity of Token1 in pool
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = math.MinInt64
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN MinTick and cur0To1 are set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(2), pair.MaxTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken1New1To0() {
	// GIVEN current tick provides Token1 at a lower price (tick 2) ie. spread tightens
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.CurrentTick0To1 = 3
	_, upper := s.setLPAtFee0Pool(1, 0, 10)
	// THEN curTick1To0 is set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(10), pair.MaxTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken1NewMaxTick() {
	// GIVEN current tick provides Token1 at a new highest price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 10
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(11, 0, 10)
	// THEN  MaxTick is set to CurrentTick's index
	s.app.DexKeeper.UpdateTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(12), pair.MaxTick)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostAddToken1NoChange() {
	// GIVEN current tick provides Token1 above 0To1 Price
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(3, 0, 10)
	// THEN no changes are made to MinTick & Cur0To1
	s.app.DexKeeper.UpdateTickPointersPostAddToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(int64(2), pair.CurrentTick0To1)
	s.Assert().Equal(int64(5), pair.MaxTick)
}

// CalcTickPointersPostRemoveToken0 ///////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken0NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = 1
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(0, 0, 0)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken0NotDrained() {
	// GIVEN current tick still has Token0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = -4
	originalPair := pair
	lower, _ := s.setLPAtFee0Pool(-10, 10, 0)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(pair, originalPair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken0DrainLiq() {
	// GIVEN current tick removes liquidity at MinTick && MinTick == Current1To0
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -6
	pair.CurrentTick1To0 = -6
	lower, _ := s.setLPAtFee0Pool(-5, 0, 0)
	// THEN Current0to1 is set to MinInt && MaxTick tick is set to MaxInt64
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(math.MinInt64, int(pair.CurrentTick1To0))
	s.Assert().Equal(math.MaxInt64, int(pair.MinTick))
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken0MinTick() {
	// GIVEN current tick removes liquidity at MinTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MinTick = -5
	pair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(-4, 0, 0)
	// THEN Current0to1 is unchanged && MinTick tick is increased
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(0, int(pair.CurrentTick1To0))
	s.Assert().Less(-5, int(pair.MinTick))
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken0CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(-1, 10, 0)
	s.setLPAtFee0Pool(-2, 10, 0)
	pair.MinTick = -3
	pair.CurrentTick1To0 = 0
	lower, _ := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current1to0 is changed to the next lowest tick (-2)
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken0(s.goCtx, &pair, &lower)
	s.Assert().Equal(-2, int(pair.CurrentTick1To0))
	s.Assert().Equal(-3, int(pair.MinTick))
}

// CalcTickPointersPostRemoveToken1 ///////////////////////////////////////////

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken1NoChange() {
	// GIVEN current tick removes liquidity between bounds
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 1
	originalPair := pair
	_, upper := s.setLPAtFee0Pool(3, 0, 0)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(pair, originalPair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken1NotDrained() {
	// GIVEN current tick still has Token1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 4
	originalPair := pair
	_, upper := s.setLPAtFee0Pool(5, 0, 10)
	// THEN no changes are made
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(pair, originalPair)
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken1DrainLiq() {
	// GIVEN current tick removes liquidity at MaxTick && MaxTick == Current0To1
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 5
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is set to MaxInt && MaxTick tick is set to MinInt
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(math.MaxInt64, int(pair.CurrentTick0To1))
	s.Assert().Equal(math.MinInt64, int(pair.MaxTick))
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken1MaxTick() {
	// GIVEN current tick removes liquidity at MaxTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	pair.MaxTick = 5
	pair.CurrentTick0To1 = 0
	_, upper := s.setLPAtFee0Pool(4, 0, 0)
	// THEN Current0to1 is unchanged && MaxTick tick is decreased
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(0, int(pair.CurrentTick0To1))
	s.Assert().Greater(5, int(pair.MaxTick))
}

func (s *CoreHelpersTestSuite) TestCalcTickPointersPostRemoveToken1CurTick() {
	// GIVEN current tick removes liquidity at the CurrentTick
	pair := s.app.DexKeeper.GetOrInitPair(s.goCtx, "TokenA", "TokenB")
	s.setLPAtFee0Pool(1, 0, 10)
	s.setLPAtFee0Pool(2, 0, 10)
	s.setLPAtFee0Pool(3, 0, 10)
	pair.MaxTick = 4
	pair.CurrentTick0To1 = 2
	_, upper := s.setLPAtFee0Pool(1, 0, 0)

	// THEN Current0to1 is changed to the next lowest tick (3)
	s.app.DexKeeper.UpdateTickPointersPostRemoveToken1(s.goCtx, &pair, &upper)
	s.Assert().Equal(3, int(pair.CurrentTick0To1))
	s.Assert().Equal(4, int(pair.MaxTick))
}
