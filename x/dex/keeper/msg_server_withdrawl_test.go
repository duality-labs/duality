package keeper_test

import (
	"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (suite *IntegrationTestSuite) TestHasBalance2() {
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

	// Set Fee List

	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	fmt.Println("FeeList")
	feeList := app.DexKeeper.GetAllFeeList(ctx)
	fmt.Println(feeList)

	fiftyDec, _ := sdk.NewDecFromStr("50")
	createResponse, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{fiftyDec},
		AmountsB:    []sdk.Dec{sdk.ZeroDec()},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{0},
		Receiver:    alice.String(),
	})

	suite.Require().Nil(err)

	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", 1))
	fmt.Println()

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("500000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("0"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse

	pairId := app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fmt.Println(app.DexKeeper.GetShares(ctx, alice.String(), pairId, 0, 0))

	sharesToRemove, _ := sdk.NewDecFromStr("50")

	withdrawlResponse, err := suite.msgServer.Withdrawl(goCtx, &types.MsgWithdrawl{
		Creator:        alice.String(),
		Receiver:       alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sharesToRemove},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{0},
	})

	suite.Require().Nil(err)
	fmt.Println("Post Withdrawl")
	fmt.Println(app.DexKeeper.GetShares(ctx, alice.String(), pairId, 0, 0))

	_ = withdrawlResponse
	_ = goCtx

	createResponse2, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{fiftyDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{1},
		Receiver:    alice.String(),
	})

	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", 2))
	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", -2))

	suite.Require().Nil(err)

	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("0"))))

	_ = createResponse2

	lowerTick, _ := app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", -2)
	upperTick, _ := app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", 2)

	fmt.Println("Upper tic", upperTick)
	suite.Require().Equal(upperTick.TickData.Reserve0AndShares[1].Reserve0, sdk.ZeroDec())
	suite.Require().Equal(upperTick.TickData.Reserve0AndShares[1].TotalShares, sdk.NewDec(50))

	suite.Require().Equal(lowerTick.TickData.Reserve1[1], sdk.NewDec(50))

	createResponse3, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{fiftyDec},
		AmountsB:    []sdk.Dec{sdk.ZeroDec()},
		TickIndexes: []int64{1},
		FeeIndexes:  []uint64{1},
		Receiver:    alice.String(),
	})

	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse3

	createResponse4, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{fiftyDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{1},
		Receiver:    alice.String(),
	})

	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", 2))
	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", -2))

	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("4000000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse3
	_ = createResponse4

	sharesToRemove, _ = sdk.NewDecFromStr("50")

	withdrawlResponse2, err := suite.msgServer.Withdrawl(goCtx, &types.MsgWithdrawl{
		Creator:        alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: []sdk.Dec{sharesToRemove},
		TickIndexes:    []int64{0},
		FeeIndexes:     []uint64{1},
		Receiver:       alice.String(),
	})

	fmt.Println("test here")
	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("450000000000000000000"))))

	fmt.Println(app.DexKeeper.GetShares(ctx, alice.String(), pairId, 0, 1))
	_ = withdrawlResponse2

	createResponse5, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{sdk.ZeroDec()},
		AmountsB:    []sdk.Dec{fiftyDec},
		TickIndexes: []int64{0},
		FeeIndexes:  []uint64{1},
		Receiver:    alice.String(),
	})

	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", 2))
	fmt.Println(app.DexKeeper.GetTickObject(ctx, "TokenA/TokenB", -2))

	suite.Require().Nil(err)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(convInt("50000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("400000000000000000000"))))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, alice, newBCoin(convInt("4000000000000000000001"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newACoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, bob, newBCoin(convInt("200000000000000000000"))))

	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newBCoin(convInt("100000000000000000000"))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, app.AccountKeeper.GetModuleAddress("dex"), newACoin(convInt("50000000000000000000"))))

	_ = createResponse5

	fmt.Println("Withdrawl Tests complete")
}
