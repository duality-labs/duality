package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *IntegrationTestSuite) TestMultiLimitOrderTick0to1() {

	fmt.Println("Limit Order Swap Tests 0 to 1")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	orderResponse1, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 1,
		TokenIn:   "TokenA",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	orderResponse2, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 2,
		TokenIn:   "TokenA",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	fmt.Println(app.DexKeeper.GetAllLimitOrderPoolReserveMap(ctx))
	fmt.Println(app.DexKeeper.GetLimitOrderPoolReserveMap(ctx, "TokenA/TokenB", 1, "TokenB", 0))
	fmt.Println(app.DexKeeper.GetLimitOrderPoolReserveMap(ctx, "TokenA/TokenB", 2, "TokenB", 0))
	fmt.Println(app.DexKeeper.GetAllTickMap(ctx))

	_ = orderResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  "TokenB",
		AmountIn: newDec("40"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("160000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("139000000000000000000"))))

	_ = swapResponse

}

func (suite *IntegrationTestSuite) TestMultiTickLimitOrder1to0() {

	fmt.Println("Limit Order Swap Tests 1 to 0")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	orderResponse1, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	orderResponse2, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: -1,
		TokenIn:   "TokenB",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  "TokenA",
		AmountIn: newDec("40"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("60000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("239000000000000000000"))))

	_ = swapResponse

}

func (suite *IntegrationTestSuite) TestMultiTickLimitOrderAndDeposit1to0() {

	fmt.Println("Limit Order Swap Tests 1 to 0")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	orderResponse1, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: -1,
		TokenIn:   "TokenB",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{newDec("25")},
		TickIndexes: []int64{-1},
		FeeIndexes:  []uint64{0},
		Receiver:    alice.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  "TokenA",
		AmountIn: newDec("40"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("60000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("239000000000000000000"))))

	_ = swapResponse

}

func (suite *IntegrationTestSuite) TestMultiTickLimitOrderAndDeposit1to0_2() {

	fmt.Println("Limit Order Swap Tests 1 to 0")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	orderResponse1, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	orderResponse2, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: -1,
		TokenIn:   "TokenB",
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse2

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{newDec("25")},
		TickIndexes: []int64{-1},
		FeeIndexes:  []uint64{0},
		Receiver:    alice.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  "TokenA",
		AmountIn: newDec("60"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("40000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("259000000000000000000"))))

	_ = swapResponse

}
