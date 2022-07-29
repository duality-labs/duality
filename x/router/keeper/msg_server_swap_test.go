package keeper_test

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//"github.com/NicholasDotSol/duality/x/router/keeper"
	//"github.com/NicholasDotSol/duality/x/router/types"
	dextypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/NicholasDotSol/duality/x/router/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//"fmt"
)


func newACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func newBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

func convInt(amt string) sdk.Int {
	IntAmt, err := sdk.NewIntFromString(amt)

	_ = err
	return IntAmt
}

func (suite *IntegrationTestSuite) TestSwap() {
	app, ctx := suite.app, suite.ctx
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	goCtx := sdk.WrapSDKContext(ctx)
	
	

	createResponse, err :=  suite.msgServerDex.SingleDeposit(goCtx, &dextypes.MsgSingleDeposit{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "50",
		Amounts1: "0",
		Receiver: alice.String(),
	})
	_ = createResponse
	suite.Require().Nil(err)

	fmt.Println("Zero to One", app.DexKeeper.GetAllTicks(ctx)[0].PoolsZeroToOne)
	fmt.Println("One To Zero", app.DexKeeper.GetAllTicks(ctx)[0].PoolsOneToZero)
	fmt.Println(app.DexKeeper.GetAllTicks(ctx))
	createResponse2, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenB",
		TokenOut: "TokenA",
		AmountIn: "26.75",
		MinOut: "5",
	})

	_ = createResponse2
	suite.Require().Nil(err)

	fmt.Println(app.DexKeeper.GetAllTicks(ctx))


	createResponse3, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenB",
		TokenOut: "TokenA",
		AmountIn: "25",
		MinOut: "5",
	})
	_ = createResponse3

	suite.Require().Nil(err)

	fmt.Println(app.DexKeeper.GetAllTicks(ctx))


}


func (suite *IntegrationTestSuite) TestSwapNoAvailablePools() {
	app, ctx := suite.app, suite.ctx
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	goCtx := sdk.WrapSDKContext(ctx)
	fmt.Println("Swap Failing Tests: ")
	fmt.Println(app.DexKeeper.GetAllTicks(ctx))
	createResponse, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenA",
		TokenOut: "TokenB",
		AmountIn: "25",
		MinOut: "5",
	})
	suite.Require().Error(err)

	_ = createResponse
	createResponse2, err :=  suite.msgServerDex.SingleDeposit(goCtx, &dextypes.MsgSingleDeposit{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "50",
		Amounts1: "0",
		Receiver: alice.String(),
	})
	_ = createResponse2
	suite.Require().Nil(err)

	
	fmt.Println(app.DexKeeper.GetAllTicks(ctx))
	createResponse3, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenA",
		TokenOut: "TokenB",
		AmountIn: "25",
		MinOut: "5",
	})

	_ = createResponse3
	suite.Require().Error(err)

	createResponse4, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenB",
		TokenOut: "TokenA",
		AmountIn: "25",
		MinOut: "30",
	})

	_ = createResponse4
	suite.Require().Error(err)

}

func (suite *IntegrationTestSuite) TestSwapThenWithdraw() {

	app, ctx := suite.app, suite.ctx
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("500000000000000000000")))
	balanceBob := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("200000000000000000000")))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	goCtx := sdk.WrapSDKContext(ctx)


	createResponse, err :=  suite.msgServerDex.SingleDeposit(goCtx, &dextypes.MsgSingleDeposit{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "50",
		Amounts1: "0",
		Receiver: alice.String(),
	})

	fmt.Println(createResponse)
	_ = createResponse

	
	suite.Require().Nil(err)

	

	createResponse2, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenB",
		TokenOut: "TokenA",
		AmountIn: "25",
		MinOut: "5",
	})
	suite.Require().Nil(err)
	_ = createResponse2

	createResponse3, err :=  suite.msgServerDex.SingleDeposit(goCtx, &dextypes.MsgSingleDeposit{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "0",
		Amounts1: "100",
		Receiver: alice.String(),
	})
	_ = createResponse3
	fmt.Println(createResponse3)
	suite.Require().Nil(err)
	

	createResponse4, err :=  suite.msgServerDex.SingleWithdraw(goCtx, &dextypes.MsgSingleWithdraw{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		SharesRemoving: "148.522167487684729064",
		Receiver: alice.String(),
	})
	
	_ = createResponse4
	
	suite.Require().Nil(err)

	//Empty but initialized token pair
	createResponse5, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenB",
		TokenOut: "TokenA",
		AmountIn: "25",
		MinOut: "5",
	})
	suite.Require().Error(err)
	_ = createResponse5

	createResponse6, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator: alice.String(),
		TokenIn: "TokenA",
		TokenOut: "TokenG",
		AmountIn: "25",
		MinOut: "5",
	})
	suite.Require().Error(err)
	_ = createResponse6
}

