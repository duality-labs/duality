package keeper_test

import (
	// stdlib
	"fmt"
	"testing"

	// cosmos SDK
	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	// duality
	// "github.com/NicholasDotSol/duality/x/dex/types"
)

// TODO: better name
type CosmostTestEnv struct {
	app              *dualityapp.App
	msgServer        types.MsgServer
	ctx              sdk.Context
	queryClient      types.QueryClient
	dexModuleAddress string
}

func cosmosEnvSetup() CosmostTestEnv {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	dexModuleAddress := app.AccountKeeper.GetModuleAddress(types.ModuleName).String()

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DexKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	msgServer := keeper.NewMsgServerImpl(app.DexKeeper)

	return CosmostTestEnv{
		app,
		msgServer,
		ctx,
		queryClient,
		dexModuleAddress,
	}
}

// TODO: move these to type utils folder or something
type TestEnv struct {
	cosmos   CosmostTestEnv
	addrs    []sdk.AccAddress
	balances map[string]sdk.Coins
	feeTiers []types.FeeList
}

func singlePoolSetup(t *testing.T, cosmos CosmostTestEnv) TestEnv {
	app, ctx := cosmos.app, cosmos.ctx

	// initialize accounts
	alice, bob := sdk.AccAddress([]byte("alice")), sdk.AccAddress([]byte("bob"))
	accAlice, accBob := app.AccountKeeper.NewAccountWithAddress(ctx, alice), app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	app.AccountKeeper.SetAccount(ctx, accBob)

	// init balances & fund the accounts
	balancesAlice := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("10000000000000000000")))
	balancesBob := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("10000000000000000000")))
	// TODO: don't use simapp
	if err := (simapp.FundAccount(app.BankKeeper, ctx, alice, balancesAlice)); err != nil {
		t.Errorf("Failed to fund %s with %s", alice, balancesAlice)
	}
	if err := (simapp.FundAccount(app.BankKeeper, ctx, bob, balancesBob)); err != nil {
		t.Errorf("Failed to fund %s with %s", bob, balancesBob)
	}

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

	addrs := []sdk.AccAddress{alice, bob}
	balances := map[string]sdk.Coins{
		addrs[0].String(): balancesAlice,
		addrs[1].String(): balancesBob,
	}

	return TestEnv{
		cosmos,
		addrs,
		balances,
		feeTiers,
	}
}

func calculateShares() sdk.Coin {
	// calculating shares minted in DepositHelper
	// TODO: comment what this corresponds to
	// if !lowerTickFound || !upperTickFound || upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Equal(sdk.ZeroDec()) {
	// sharesMinted = amount0.Add(amount1.Mul(price))
	// } else {
	// // If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	// if trueAmount0.GT(sdk.ZeroDec()) && ((tickIndex+fee > pair.TokenPair.CurrentTick0To1) && (tickIndex+fee < pair.TokenPair.CurrentTick1To0)) {
	// pair.TokenPair.CurrentTick1To0 = tickIndex + fee
	// }
	// if trueAmount1.GT(sdk.ZeroDec()) && ((tickIndex-fee > pair.TokenPair.CurrentTick0To1) && (tickIndex-fee < pair.TokenPair.CurrentTick1To0)) {
	// pair.TokenPair.CurrentTick0To1 = tickIndex - fee
	// }
	// }
	i, _ := sdk.NewIntFromString("1")
	return sdk.NewCoin("TickShares", i)
}

func testSingleDeposit(t *testing.T, coinA sdk.Coin, coinB sdk.Coin, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64, env *TestEnv) {
	// TODO
	// - take in pair of Coins to allow for initializing a pair or using existing pair (use Coin.Denom when calling Deposit)
	// - take in receiver instead of hardcoding alice
	// - modify this for len(tickIndexes) > 1
	app, ctx := env.cosmos.app, env.cosmos.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// GIVEN inital balances
	accBalanceAInitial, accBalanceBInitial := newACoin(env.balances[acc.String()].AmountOf(coinA.Denom)), newBCoin(env.balances[acc.String()].AmountOf(coinB.Denom))
	// verify acc has exactly the balance passed in from env
	if !(app.BankKeeper.GetBalance(ctx, acc, coinA.Denom).IsEqual(accBalanceAInitial)) {
		t.Errorf("%s's initial balance of %s does not match env: %s", acc, coinA, accBalanceAInitial)
	}
	if !(app.BankKeeper.GetBalance(ctx, acc, coinB.Denom).IsEqual(accBalanceBInitial)) {
		t.Errorf("%s's initial balance of %s does not match env: %s", acc, coinB, accBalanceBInitial)
	}
	// get bank initial balance
	dexAllCoinsInitial := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAInitial, dexBalanceBInitial := newACoin(dexAllCoinsInitial.AmountOf(coinA.Denom)), newBCoin(dexAllCoinsInitial.AmountOf(coinB.Denom))
	// get amount of shares before depositing

	// WHEN depositing the specified amounts coinA and coinB
	// (discard message response because we don't need it)
	_, err := env.cosmos.msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator:     acc.String(),
		TokenA:      coinA.Denom,
		TokenB:      coinB.Denom,
		AmountsA:    []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, // Coin is already denominated in 1e18
		AmountsB:    []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)},
		TickIndexes: tickIndexes, // single deposit
		FeeIndexes:  feeTiers,
		Receiver:    acc.String(),
	})

	// THEN no error, alice's balances changed only by the amount depoisited, funds transfered to dex module, and position minted with appropriate fee tier
	// verify no error
	if err != nil {
		t.Errorf("Deposit of %s, %s by %s failed: %s", coinA, coinB, acc, err)
	}

	// verify alice's resulting balances is aliceBalanceInitial - depositCoin
	accBalanceAFinal, accBalanceBFinal := accBalanceAInitial.Sub(coinA), accBalanceBInitial.Sub(coinB)
	if !(app.BankKeeper.GetBalance(ctx, acc, coinA.Denom).IsEqual(accBalanceAFinal)) {
		t.Errorf("%s's final balance of %s does not reflect deposit", acc, coinA)
	}
	if !(app.BankKeeper.GetBalance(ctx, acc, coinB.Denom).IsEqual(accBalanceBFinal)) {
		t.Errorf("%s's final balance of %s does not reflect deposit", acc, coinB)
	}

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexAllCoinsFinal := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAFinal, dexBalanceBFinal := dexBalanceAInitial.Add(coinA), dexBalanceBInitial.Add(coinB)
	if !(newACoin(dexAllCoinsFinal.AmountOf(coinA.Denom)).IsEqual(dexBalanceAFinal)) {
		t.Errorf("Dex module's final balance of %s does not reflect deposit", coinA.Denom)
	}
	if !(newBCoin(dexAllCoinsFinal.AmountOf(coinB.Denom)).IsEqual(dexBalanceBFinal)) {
		t.Errorf("Dex module's final balance of %s does not reflect deposit", coinB.Denom)
	}

	// verify shares minted for alice
	// accShares, _ := calculateShares()
	// pairId := "0"
	// mintedShares, found := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])
	// suite.Require().True(found)
	// suite.Require().True(mintedShares.SharesOwned.Equal(accShares))
}

func TestMinFeeTier(t *testing.T) {
	fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/MinFeeTier")
	env := singlePoolSetup(t, cosmosEnvSetup())

	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with min fee tier
	tickIndex := []int64{0}
	minFeeTier := []uint64{uint64(env.feeTiers[0].Id)}

	// validity Requires are done inside testSingleDeposit
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, minFeeTier, &env)
}

func TestMaxFeeTier(t *testing.T) {
	fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/MaxFeeTier")
	env := singlePoolSetup(t, cosmosEnvSetup())

	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with min fee tier
	tickIndex := []int64{0}
	maxFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Id)}

	// validity Requires are done inside testSingleDeposit
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)
}
