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