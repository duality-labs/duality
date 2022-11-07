package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *IntegrationTestSuite) TestCancelLimitOrderSingle() {

	fmt.Println("Limit Order Cancel Test 1")
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
		TokenIn:   "TokenB",
		TokenOut:  "TokenA",
		TickIndex: 0,
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	cancelResponse1, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("25"),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))

	_ = cancelResponse1

}

func (suite *IntegrationTestSuite) TestCancelLimitOrder_2() {

	fmt.Println("Limit Order Cancel Tests 2")
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
		TokenIn:   "TokenB",
		TokenOut:  "TokenA",
		TickIndex: 0,
		AmountIn:  newDec("40"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("460000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("40000000000000000000"))))

	_ = orderResponse1

	cancelResponse1, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("15"),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("475000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("25000000000000000000"))))
	_ = cancelResponse1

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenIn:  "TokenA",
		TokenOut: "TokenB",
		AmountIn: newDec("40"),
		MinOut:   newDec("14"),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))

	_ = swapResponse

	cancelResponse2, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("15"),
	})

	suite.Require().Error(err)

	_ = cancelResponse2

}

func (suite *IntegrationTestSuite) TestCancelLimitOrder_3() {

	fmt.Println("Limit Order Cancel Tests 2")
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
		TokenIn:   "TokenB",
		TokenOut:  "TokenA",
		TickIndex: 0,
		AmountIn:  newDec("40"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("460000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("40000000000000000000"))))

	_ = orderResponse1

	cancelResponse1, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("15"),
	})
	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("475000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("25000000000000000000"))))
	_ = cancelResponse1

	cancelResponse2, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("15"),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("490000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("10000000000000000000"))))
	_ = cancelResponse2

	cancelResponse3, err := suite.msgServer.CancelLimitOrder(goCtx, &types.MsgCancelLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
		SharesOut: newDec("15"),
	})

	suite.Require().Error(err)

	_ = cancelResponse3
}
