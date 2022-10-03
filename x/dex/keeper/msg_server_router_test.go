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

	fiftyDec, _ := sdk.NewDecFromStr("50")
	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     lp.String(),
		TokenA:      "JUNO",
		TokenB:      "USDC",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{fiftyDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
		Receiver:    lp.String(),
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
		Creator:     lp.String(),
		TokenA:      "USDC",
		TokenB:      "STARS",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{fiftyDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
		Receiver:    lp.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse2

	//fmt.Println(app.DexKeeper.GetAllPairMap(ctx))

	swapResponse, err := suite.msgServer.Route(goCtx, &types.MsgRoute{
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

	_ = pairId
	_ = goCtx

}

func (suite *IntegrationTestSuite) TestMultiHopRoute() {

}

func (suite *IntegrationTestSuite) TestSwapRoute2() {

}
