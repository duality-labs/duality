package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *IntegrationTestSuite) TestMultiTickLimitOrder1to0_Withdraw() {

	fmt.Println("Limit Order Withdraw Tests 1 to 0")
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

	orderResponse2, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenIn:   "TokenB",
		TokenOut:  "TokenA",
		TickIndex: -1,
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenIn:  "TokenA",
		TokenOut: "TokenB",
		AmountIn: newDec("40"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)

	suite.Require().False(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("60000000000000000001"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("60000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("239000000000000000000"))))

	_ = swapResponse

	withdrawResponse, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})

	suite.Require().Nil(err)

	_ = withdrawResponse

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100025000000000000000000"))))

	withdrawResponse2, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: -1,
		KeyToken:  "TokenB",
		Key:       0,
	})

	suite.Require().Nil(err)

	_ = withdrawResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100033900000000000000000"))))

}

func (suite *IntegrationTestSuite) TestMultiTickLimitOrder0to1_Withdraw() {

	fmt.Println("Limit Order Withdraw Tests 1 to 0")
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
		TokenIn:   "TokenA",
		TokenOut:  "TokenB",
		TickIndex: 0,
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse1

	orderResponse2, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenIn:   "TokenA",
		TokenOut:  "TokenB",
		TickIndex: 1,
		AmountIn:  newDec("25"),
	})

	suite.Require().Nil(err)

	_ = orderResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	swapResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		Receiver: bob.String(),
		TokenOut: "TokenA",
		TokenIn:  "TokenB",
		AmountIn: newDec("40"),
		MinOut:   newDec("30"),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("139000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("160000000000000000000"))))

	_ = swapResponse

	withdrawResponse, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenA",
		Key:       0,
	})

	suite.Require().Nil(err)

	_ = withdrawResponse

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("525000000000000000000"))))

	withdrawResponse2, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 1,
		KeyToken:  "TokenA",
		Key:       0,
	})

	suite.Require().Nil(err)

	_ = withdrawResponse2

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("533900000000000000000"))))

}

func (suite *IntegrationTestSuite) ErrorCasesWithdrawLimitOrder() {

	fmt.Println("Limit Order Withdraw Tests 1 to 0")
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

	// errors because order at this tick does not exists
	withdrawResponse, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})

	suite.Require().Error(err)

	_ = withdrawResponse

	orderResponse1, err := suite.msgServer.PlaceLimitOrder(goCtx, &types.MsgPlaceLimitOrder{
		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenIn:   "TokenA",
		TokenOut:  "TokenB",
		TickIndex: 0,
		AmountIn:  newDec("25"),
	})

	_ = orderResponse1

	// errors because called by not the owner
	withdrawResponse2, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   bob.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})

	suite.Require().Error(err)

	_ = withdrawResponse2

	// Errors because of wrong KeyToken
	withdrawResponse3, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenA",
		Key:       0,
	})

	suite.Require().Error(err)

	_ = withdrawResponse3

	// errors because of wrong key
	withdrawResponse4, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       1,
	})

	suite.Require().Error(err)

	_ = withdrawResponse4

	withdrawResponse5, err := suite.msgServer.WithdrawFilledLimitOrder(goCtx, &types.MsgWithdrawFilledLimitOrder{

		Creator:   alice.String(),
		Receiver:  bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       1,
	})

	suite.Require().Nil(err)

	_ = withdrawResponse5

	_ = ctx
}
