package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewBankBalance(tokenName string, amt int64) sdk.Coin {
	return sdk.NewCoin(tokenName, sdk.NewIntFromBigInt(sdk.NewDec(amt).BigInt()))
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
	balanceAlice := sdk.NewCoins(sdk.NewCoin("JUNO", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt())))
	fmt.Println("balance alice JUNO", balanceAlice.AmountOf("JUNO"))
	balanceLP := sdk.NewCoins(sdk.NewCoin("STARS", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt())), sdk.NewCoin("DUAL", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt())), sdk.NewCoin("USDC", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt())))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balanceAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, lp, balanceLP))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, sdk.NewCoin("JUNO", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt()))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("STARS", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt()))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("DUAL", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt()))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, sdk.NewCoin("USDC", sdk.NewIntFromBigInt(sdk.NewDec(100000).BigInt()))))
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

	tenThousandDec := sdk.NewDec(10000)
	// fmt.Println("B2", tenThousandDec)
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
	suite.Require().True(app.BankKeeper.HasBalance(ctx, lp, NewBankBalance("USDC", 90000)))

	// Confirm Pool Balance
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), NewBankBalance("USDC", 10000)))

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

	// Slightly better price than 1:1
	createResponse3, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     lp.String(),
		TokenA:      "JUNO",
		TokenB:      "STARS",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{sdk.NewDec(5), sdk.NewDec(5)},
		TickIndexes: []int64{-1, -2},
		FeeIndexes:  []uint64{0},
		Receiver:    lp.String(),
	})

	suite.Require().Nil(err)

	_ = createResponse3

	// fmt.Println("Making it to route")
	swapResponse, err := suite.msgServer.Route(goCtx, &types.MsgRoute{
		Creator:  alice.String(),
		TokenIn:  "JUNO",
		TokenOut: "STARS",
		AmountIn: sdk.NewDec(10),
		MinOut:   sdk.NewDec(0),
		Receiver: alice.String(),
	})
	// fmt.Println("Post Route")

	_ = swapResponse
	// TODO: Figure out way to determine correct amount out
	print(app.BankKeeper.GetBalance(ctx, alice, "STARS").Amount.String())
	suite.Require().Nil(err)

	_ = pairId
	_ = goCtx

}

func (suite *IntegrationTestSuite) TestMultiHopRoute() {

}

func (suite *IntegrationTestSuite) TestSwapRoute2() {

}
