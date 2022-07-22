package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//"fmt"
	"github.com/NicholasDotSol/duality/x/dex/types"
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

func (suite *IntegrationTestSuite) TestHasBalance() {
	app, ctx := suite.app, suite.ctx
	addr := sdk.AccAddress([]byte("addr1_______________"))

	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	app.AccountKeeper.SetAccount(ctx, acc)

	balances := sdk.NewCoins(newACoin(sdk.NewInt(100)))

	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(99))))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, addr, balances))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(101))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(100))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(1))))
}

func (suite *IntegrationTestSuite) TestSingleDeposit() {
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
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

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("1000000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	goCtx := sdk.WrapSDKContext(ctx)
	createResponse, err := suite.msgServer.SingleDeposit(goCtx, &types.MsgSingleDeposit{
		Creator:  alice.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "50",
		Amounts1: "100",
		Receiver: alice.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))

}
