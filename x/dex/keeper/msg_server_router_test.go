package keeper_test

import (
	"fmt"
	"math/big"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewCoin(amt *big.Int, tokenName string) sdk.Coin {
	return sdk.NewCoin(tokenName, sdk.NewIntFromBigInt(amt))
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
	balanceAlice := sdk.NewCoins(sdk.NewCoin("JUNO", sdk.NewInt(100000)))
	fmt.Println(balanceAlice.AmountOf("JUNO"))
	balanceLP := sdk.NewCoins(sdk.NewCoin("STARS", sdk.NewInt(100000)), sdk.NewCoin("DUAL", sdk.NewInt(100000)), sdk.NewCoin("USDC", sdk.NewInt(100000)))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, lp, balanceLP))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, sdk.NewCoin("JUNO", sdk.NewInt(100000))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("STARS", sdk.NewInt(100000))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("DUAL", sdk.NewInt(100000))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("USDC", sdk.NewInt(100000))))
	goCtx := sdk.WrapSDKContext(ctx)

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	//fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	_ = feeList
	// fmt.Println(feeList)

	tenThousandDec := sdk.NewDecFromInt(sdk.NewInt(10000))
	fmt.Println(tenThousandDec)
	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     lp.String(),
		TokenA:      "JUNO",
		TokenB:      "USDC",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{tenThousandDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
		Receiver:    lp.String(),
	})

	suite.Require().Nil(err)

	// Confirm LP Balance post Deposit
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("JUNO", sdk.NewInt(90000))))

	// Confirm Pool Balance
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), sdk.NewCoin("JUNO", sdk.NewInt(10000))))

	_ = createResponse

	pairId := app.DexKeeper.CreatePairId("JUNO", "STARS")

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     lp.String(),
		TokenA:      "USDC",
		TokenB:      "STARS",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{tenThousandDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
		Receiver:    lp.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse2

	//fmt.Println(app.DexKeeper.GetAllPairMap(ctx))

	swapResponse, err := suite.msgServer.Route(goCtx, &types.MsgRoute{
		Creator:  alice.String(),
		TokenIn:  "JUNO",
		TokenOut: "STARS",
		AmountIn: "10",
		MinOut:   "0",
		Receiver: alice.String(),
	})

	_ = swapResponse
	print(app.BankKeeper.GetBalance(ctx, alice, "STARS").Amount.String())
	suite.Require().Nil(err)

	_ = pairId
	_ = goCtx

}

func (suite *IntegrationTestSuite) TestMultiHopRoute() {

}

func (suite *IntegrationTestSuite) TestSwapRoute2() {

}
