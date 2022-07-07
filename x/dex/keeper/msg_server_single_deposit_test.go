package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//"fmt"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

func newACoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin("TokenA", amt)
}

func newBCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin("TokenB", amt)
}

func (suite *IntegrationTestSuite) TestHasBalance() {
	app, ctx := suite.app, suite.ctx
	addr := sdk.AccAddress([]byte("addr1_______________"))

	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	app.AccountKeeper.SetAccount(ctx, acc)

	balances := sdk.NewCoins(newACoin(100))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(99)))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, addr, balances))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(101)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(100)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(1)))
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

	balanceAlice := sdk.NewCoins(newACoin(100), newBCoin(500))
	balanceBob := sdk.NewCoins(newACoin(100), newBCoin(200))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balanceBob))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(100)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(100)))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(1000)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(500)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(200)))

	goCtx := sdk.WrapSDKContext(ctx)
	createResponse, err := suite.msgServer.SingleDeposit(goCtx, &types.MsgSingleDeposit{
		Creator: alice.String(),
		Token0:    "TokenA",
		Token1:   "TokenB",
		Price: "1.0",
		Fee: 300,
		Amounts0: 50,
		Amounts1: 100,
		Receiver: alice.String(),
	})
	suite.Require().Nil(err)
	_ = createResponse
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(100)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(100)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(400)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(200)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(50)))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(100)))
	
}

// func (suite *IntegrationTestSuite) TestSingleDeposit() {
// 	fmt.Println("test 1")
// 	suite.setupSuiteWithBalances()
// 	fmt.Println("test 2")
// 	goCtx := sdk.WrapSDKContext(suite.ctx)
// 	// string creator = 1;
// 	// string token0 = 2P;
// 	// string token1 = 3;
// 	// string price = 4;
// 	// uint64 fee = 5;
// 	// uint64 amounts0 = 6;
// 	// uint64 amounts1 = 7;
// 	// string receiver = 8;

// 	createResponse, err := suite.msgServer.SingleDeposit(goCtx, &types.MsgSingleDeposit{
// 		Creator: alice,
// 		Token0:    "A",
// 		Token1:   "B",
// 		Price: "1",
// 		Fee: 300,
// 		Amounts0: 500,
// 		Amounts1: 100,
// 		Receiver: alice,
// 	})
// 	suite.Require().Nil(err)

// 	_ = createResponse
// }
