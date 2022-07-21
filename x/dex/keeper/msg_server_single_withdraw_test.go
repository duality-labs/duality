package keeper_test

import (

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//"fmt"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (suite *IntegrationTestSuite) TestSingleWithdraw() {

	//fmt.Println("Withdraw Testing:")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	bob := sdk.AccAddress([]byte("bob"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accBob)

	//100 & 500
	balanceAlice := sdk.NewCoins(newACoin(convInt("100000000000000000000")), newBCoin(convInt("500000000000000000000")))
	//100 & 200
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

	createResponse2, err := suite.msgServer.SingleDeposit(goCtx, &types.MsgSingleDeposit{
		Creator:  bob.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		Amounts0: "50",
		Amounts1: "0",
		Receiver: bob.String(),
	})
	
	_ = createResponse2
	suite.Require().Nil(err)
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("48000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))

	withdrawResponse, err := suite.msgServer.SingleWithdraw(goCtx, &types.MsgSingleWithdraw{
		Creator:  bob.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		SharesRemoving: "50",
		Receiver: bob.String(),

	})
	
	_ = withdrawResponse
	
	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("75000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("75000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("73300000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("225000000000000000000"))))

	withdrawResponse2, err := suite.msgServer.SingleWithdraw(goCtx, &types.MsgSingleWithdraw{
		Creator:  bob.String(),
		Token0:   "TokenA",
		Token1:   "TokenB",
		Price:    "1.0",
		Fee:      "300",
		SharesRemoving: "50",
		Receiver: bob.String(),

	})
	suite.Require().Error(err)
	_ = withdrawResponse2




}