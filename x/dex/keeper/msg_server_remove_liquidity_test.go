package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *IntegrationTestSuite) TestHasBalanceRemoveLiqudity() {
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

func (suite *IntegrationTestSuite) TestSingleWithdrawl() {
	fmt.Println("Testing TestSingleWithdrawl")
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

	createResponse, err := suite.msgServer.CreatePair(goCtx, &types.MsgCreatePair{
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
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	// Note this is the "correct"sorted order based of sha256 sort in sort_tokens
	expectedTick := types.Ticks{
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
	tickactual, _ := app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "300", "LP")

	expectedIndexQueue, _ := app.DexKeeper.GetIndexQueue(ctx, "TokenB", "TokenA", 0)

	actualIndexQueueArray := []*types.IndexQueueType{

		&types.IndexQueueType{
			Price: newDec("1"),
			Fee:   newDec("0.03"),
			Orderparams: &types.OrderParams{
				OrderRule:   "",
				OrderType:   "LP",
				OrderShares: newDec("1.5"),
			},
		},
	}

	actualIndexQueue := types.IndexQueue{
		Index: 0,
		Queue: actualIndexQueueArray,
	}

	suite.Require().Equal(expectedIndexQueue, actualIndexQueue)

	suite.Require().Equal(expectedTick, tickactual)

	_ = createResponse

	withdrawResponse, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1.5",
		OrderType: "LP",
		Receiver:  alice.String(),
	})

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("1000000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	tickactual2, _ := app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "300", "LP")

	suite.Require().True(tickactual2.Reserve0.IsNil())
	suite.Require().True(tickactual2.Reserve1.IsNil())
	suite.Require().True(tickactual2.TotalShares.IsNil())
	suite.Require().True(tickactual2.PairPrice.IsNil())
	suite.Require().True(tickactual2.PairFee.IsNil())
	suite.Require().True(tickactual2.Price == "")
	suite.Require().True(tickactual2.Fee == "")
	suite.Require().True(tickactual2.OrderType == "")

	indexQueue2, _ := app.DexKeeper.GetIndexQueue(ctx, "TokenB", "TokenA", 0)

	suite.Require().True(len(indexQueue2.Queue) == 0)

	_ = withdrawResponse

	withdrawResponse2, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1.5",
		OrderType: "LP",
		Receiver:  alice.String(),
	})

	suite.Require().Error(err)

	_ = withdrawResponse2

	createResponse2, err := suite.msgServer.AddLiquidity(goCtx, &types.MsgAddLiquidity{
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
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("80000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("20000000000000000000"))))

	_ = createResponse2

	createResponse3, err := suite.msgServer.AddLiquidity(goCtx, &types.MsgAddLiquidity{
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

	fmt.Println(app.DexKeeper.GetTicks(ctx, "TokenB", "TokenA", "1.0", "300", "LP"))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("40000000000000000000"))))

	_ = createResponse3

	withdrawResponse3, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1333.333333333333333321",
		OrderType: "LP",
		Receiver:  bob.String(),
	})

	_ = withdrawResponse3

	suite.Require().Error(err)

	withdrawResponse4, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   bob.String(),
		TokenA:    "TokenC",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1333.333333333333333321",
		OrderType: "LP",
		Receiver:  bob.String(),
	})

	_ = withdrawResponse4

	suite.Require().Error(err)

	withdrawResponse5, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1333.333333333333333320",
		OrderType: "Limit",
		Receiver:  bob.String(),
	})

	_ = withdrawResponse5

	suite.Require().Error(err)

	withdrawResponse6, err := suite.msgServer.RemoveLiquidity(goCtx, &types.MsgRemoveLiquidity{
		Creator:   alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		Index:     0,
		Price:     "1.0",
		Fee:       "300",
		Shares:    "1333.333333333333333320",
		OrderType: "LP",
		Receiver:  alice.String(),
	})

	_ = withdrawResponse6

	suite.Require().Error(err)
	_ = ctx

}
