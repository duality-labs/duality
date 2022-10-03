package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewCoin(amt float64, tokenName string) sdk.Coin {
	var newAmt int64 = int64(amt)
	return sdk.NewCoin(tokenName, sdk.NewInt(newAmt))
}

// Swapping from JUNO to STARS through USDC
func (suite *IntegrationTestSuite) TestBasicMultiHopRoute() {
	fmt.Println("Route Tests")
	app, ctx := suite.app, suite.ctx
	//holderAcc := authtypes.NewEmptyModuleAccount("holder")
	alice := sdk.AccAddress([]byte("alice"))
	lp := sdk.AccAddress([]byte("lp"))

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accLP := app.AccountKeeper.NewAccountWithAddress(ctx, lp)
	app.AccountKeeper.SetAccount(ctx, accLP)

	// Base 18 decimals
	balanceAlice := sdk.NewCoins(NewCoin(1e23, "JUNO"))
	balanceLP := sdk.NewCoins(NewCoin(1e23, "STARS"), NewCoin(1e23, "DUAL"), NewCoin(1e23, "USDC"))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, lp, balanceLP))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, NewCoin(1e23, "JUNO")))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, NewCoin(1e23, "STARS")))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, NewCoin(1e23, "DUAL")))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, NewCoin(1e23, "USDC")))
	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:   lp.String(),
		TokenA:    "JUNO",
		TokenB:    "USDC",
		AmountA:   "0",
		AmountB:   "100",
		TickIndex: 0,
		FeeIndex:  0,
		Receiver:  lp.String(),
	})

	suite.Require().Nil(err)

	// Confirm LP Balance post Deposit
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, NewCoin(9.9e22, "STARS")))

	// Confirm Pool Balances
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), NewCoin(0, "STARS")))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), NewCoin(100e18, "JUNO")))

	_ = createResponse

	pairId := app.DexKeeper.CreatePairId("JUNO", "STARS")

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:   lp.String(),
		TokenA:    "USDC",
		TokenB:    "STARS",
		AmountA:   "0",
		AmountB:   "100",
		TickIndex: 0,
		FeeIndex:  1,
		Receiver:  lp.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse2

	//fmt.Println(app.DexKeeper.GetAllPairMap(ctx))

	swapResponse, err := suite.msgServer.SwapRoute(goCtx, &types.MsgSwap{
		Creator:  alice.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: "20",
		TokenIn:  "TokenB",
		MinOut:   "10",
		Receiver: alice.String(),
	})

	_ = swapResponse

	suite.Require().Nil(err)

	_ = createResponse4
	_ = pairId
	_ = goCtx

}

func (suite *IntegrationTestSuite) TestMultiHopRoute() {
	fmt.Println("Swap Tests")
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

	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountA:   "50",
		AmountB:   "0",
		TickIndex: 0,
		FeeIndex:  0,
		Receiver:  alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("99950000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse
	_ = goCtx

	swapRepsone, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: "49.99500",
		TokenIn:  "TokenB",
		MinOut:   "10",
		Receiver: bob.String(),
	})

	_ = swapRepsone

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("49995000000000000000"))))

}

func (suite *IntegrationTestSuite) TestSwapRoute2() {
	fmt.Println("Swap Tests")
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

	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountA:   "50",
		AmountB:   "0",
		TickIndex: 0,
		FeeIndex:  0,
		Receiver:  alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("99950000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		AmountA:   "50",
		AmountB:   "0",
		TickIndex: 100000,
		FeeIndex:  0,
		Receiver:  alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("99900000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse2

	swapRepsone, err := suite.msgServer.Swap(goCtx, &types.MsgSwap{
		Creator:  bob.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		AmountIn: "60",
		TokenIn:  "TokenB",
		MinOut:   "10",
		Receiver: bob.String(),
	})

	_ = swapRepsone

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("60000000000000000000"))))

	_ = goCtx

}
