package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Returns -1 if error
func newDec(amount string) sdk.Dec {
	amt, err := sdk.NewDecFromStr(amount)

	if err != nil {
		return sdk.NewDec(-1)
	}

	return amt

}

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
	fmt.Println("Testing TestSingleDeposit")
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

	createResponse, err := suite.msgServer.AddLiquidity(goCtx, &types.MsgAddLiquidity{
		Creator:        alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		TokenDirection: "TokenB",
		Index:          0,
		Price:          "1.0",
		Fee:            "0",
		AmountA:        "0",
		AmountB:        "50",
		OrderType:      "LP",
		Receiver:       alice.String(),
	})

	suite.Require().Error(err)

	_ = createResponse

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().False(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))

	createResponse2, err := suite.msgServer.CreatePair(goCtx, &types.MsgCreatePair{
		Creator:        alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		TokenDirection: "TokenB",
		Index:          0,
		Price:          "1.0",
		Fee:            "0",
		AmountA:        "0",
		AmountB:        "50",
		OrderType:      "LP",
		Receiver:       alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	expectedTick := types.Ticks{
		Price:       "1.0",
		Fee:         "0",
		OrderType:   "LP",
		Reserve0:    newDec("50"),
		Reserve1:    newDec("0"),
		PairPrice:   newDec("1"),
		PairFee:     newDec("0"),
		TotalShares: newDec("0"),
		Orderparams: &types.OrderParams{
			OrderRule:   "",
			OrderType:   "LP",
			OrderShares: newDec("0"),
		},
	}
	tickactual, _ := app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "0", "LP")

	suite.Require().Equal(expectedTick, tickactual)

	_ = createResponse2

	createResponse3, err := suite.msgServer.AddLiquidity(goCtx, &types.MsgAddLiquidity{
		Creator:        alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		TokenDirection: "TokenB",
		Index:          0,
		Price:          "1.0",
		Fee:            "300",
		AmountA:        "0",
		AmountB:        "50",
		OrderType:      "LP",
		Receiver:       alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	expectedTick2 := types.Ticks{
		Price:       "1.0",
		Fee:         "300",
		OrderType:   "LP",
		Reserve0:    newDec("50"),
		Reserve1:    newDec("0"),
		PairPrice:   newDec("1"),
		PairFee:     newDec("0.03"),
		TotalShares: newDec("1.5"),
		Orderparams: &types.OrderParams{
			OrderRule:   "",
			OrderType:   "LP",
			OrderShares: newDec("1.5"),
		},
	}
	tickactual2, _ := app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "300", "LP")

	suite.Require().Equal(expectedTick2, tickactual2)

	_ = createResponse3

	// suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	// suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	// suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	// suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))
	// suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))
	// suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))

	createResponse4, err := suite.msgServer.CreatePair(goCtx, &types.MsgCreatePair{
		Creator:        alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		TokenDirection: "TokenB",
		Index:          0,
		Price:          "1.0",
		Fee:            "0",
		AmountA:        "0",
		AmountB:        "50",
		OrderType:      "LP",
		Receiver:       alice.String(),
	})
	suite.Require().Error(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	_ = createResponse4
	createResponse5, err := suite.msgServer.AddLiquidity(goCtx, &types.MsgAddLiquidity{
		Creator:        bob.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		TokenDirection: "TokenA",
		Index:          0,
		Price:          "1.0",
		Fee:            "300",
		AmountA:        "20",
		AmountB:        "0",
		OrderType:      "LP",
		Receiver:       bob.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("80000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("20000000000000000000"))))

	expectedTick3 := types.Ticks{
		Price:       "1.0",
		Fee:         "300",
		OrderType:   "LP",
		Reserve0:    newDec("50"),
		Reserve1:    newDec("20"),
		PairPrice:   newDec("1"),
		PairFee:     newDec("0.03"),
		TotalShares: newDec("668.166666666666666660"),
		Orderparams: &types.OrderParams{
			OrderRule:   "",
			OrderType:   "LP",
			OrderShares: newDec("668.166666666666666660"),
		},
	}
	tickactual3, _ := app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "300", "LP")

	suite.Require().Equal(expectedTick3, tickactual3)

	_ = createResponse5

	_ = goCtx

}
