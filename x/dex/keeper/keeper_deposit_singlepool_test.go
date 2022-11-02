package keeper_test

import (
	// stdlib
	"fmt"

	// cosmos SDK
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// duality
	// "github.com/NicholasDotSol/duality/x/dex/types"
)

// TODO: move these to type utils folder or something
type TestEnv struct {
	addrs    []sdk.AccAddress
	balances []sdk.Coins
	feeTiers []types.FeeList
}

func (suite *IntegrationTestSuite) SinglePoolSetup() TestEnv {
	fmt.Println("[UnitTests|Keeper|SinglePool|MinFeeTier] Starting test.")
	app, ctx := suite.app, suite.ctx

	// initialize accounts
	alice, bob := sdk.AccAddress([]byte("alice")), sdk.AccAddress([]byte("bob"))
	accAlice, accBob := app.AccountKeeper.NewAccountWithAddress(ctx, alice), app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	app.AccountKeeper.SetAccount(ctx, accBob)

	// init balances & fund the accounts
	balancesAlice := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("10000000000000000000")))
	balancesBob := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("10000000000000000000")))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, alice, balancesAlice))
	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, bob, balancesBob))

	// add the fee tiers of 1, 3, 5 ticks
	feeTiers := []types.FeeList{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
	}

	// TODO: why does the append require a FeeList object with an Id if Id is assigned to count in append?
	app.DexKeeper.AppendFeeList(ctx, feeTiers[0])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[1])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[2])

	return TestEnv{
		addrs:    []sdk.AccAddress{alice, bob},
		balances: []sdk.Coins{balancesAlice, balancesBob},
		feeTiers: feeTiers,
	}
}

func (suite *IntegrationTestSuite) testFeeTiers(feeTiers []uint64, env *TestEnv) {
	app, ctx := suite.app, suite.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// GIVEN inital balances
	alice := env.addrs[0]
	aliceInitialBalances := env.balances[0]
	// verify alice has 10*1e18, i.e. unchanged from setup
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, aliceInitialBalances[0]))
	// verify verify alice has 10*1e18, bob has 20*1e18 of CoinB, i.e. unchanged from setup
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, aliceInitialBalances[1]))

	// WHEN depositing 10*1e18 of token A and 10*1e18 of token B
	depositADec, _ := sdk.NewDecFromStr("10") // deposit 10 of token A
	depositBDec, _ := sdk.NewDecFromStr("10") // deposit 10 of token B
	// convert FeeList into []uint64

	// discard message response because we don't need it
	_, err := suite.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     alice.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    []sdk.Dec{depositADec},
		AmountsB:    []sdk.Dec{depositBDec},
		TickIndexes: []int64{0},
		FeeIndexes:  feeTiers,
		Receiver:    alice.String(),
	})

	// THEN
	// verify no error
	suite.Require().Nil(err)
	// verify balances changed only by the amount deposited
	aliceFinalBalanceA, aliceFinalBalanceB := sdk.NewDecFromInt(aliceInitialBalances[0].Amount).Sub(depositADec), sdk.NewDecFromInt(aliceInitialBalances[1].Amount).Sub(depositBDec)
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newACoin(aliceFinalBalanceA.RoundInt())))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, alice, newBCoin(aliceFinalBalanceB.RoundInt())))
}

func (suite *IntegrationTestSuite) TestMinFeeTier() {
	fmt.Println("[UnitTests|Keeper|SinglePool|MinFeeTier] Starting test.")
	env := suite.SinglePoolSetup()

	// deposit with min fee tier
	minFeeTier := []uint64{uint64(env.feeTiers[0].Fee)}

	suite.testFeeTiers(minFeeTier, &env)
}

func (suite *IntegrationTestSuite) TestMaxFeeTier() {
	fmt.Println("[UnitTests|Keeper|SinglePool|MinFeeTier] Starting test.")
	env := suite.SinglePoolSetup()

	// deposit with min fee tier
	minFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Fee)}

	suite.testFeeTiers(minFeeTier, &env)
}
