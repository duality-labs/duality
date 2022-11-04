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
	app       *dualityapp.App
	msgServer types.MsgServer
	// TODO: keeping ctx in struct is bad practice: https://pkg.go.dev/context#pkg-overview
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

// Helpers for calculating the amount of shares that should be minted
// These are here for the sake of more concise unit tests and the corresponding code in core.go
// should eventually be refactored so that core.go is modularized for easier unit testing and readability

func calculateSharesEmpty(amount0 sdk.Dec, amount1 sdk.Dec, price sdk.Dec) sdk.Dec {
	return amount0.Add(amount1.Mul(price))
}

func calculateSharesNonEmpty(amount sdk.Dec, reserve sdk.Dec, totalShares sdk.Dec) sdk.Dec {
	return amount.Quo(reserve).Mul(totalShares)
}

func calculateSharesPure(
	amount0 sdk.Dec,
	trueAmount0 sdk.Dec,
	amount1 sdk.Dec,
	trueAmount1 sdk.Dec,
	price sdk.Dec,
	feeIndex uint64,
	lowerTickFound bool,
	lowerReserve1 sdk.Dec,
	upperTickFound bool,
	upperReserve0 sdk.Dec,
	upperTotalShares sdk.Dec,
) sdk.Dec {
	// calculating shares minted in DepositHelper
	if !lowerTickFound || !upperTickFound || upperTotalShares.Equal(sdk.ZeroDec()) {
		// this case corresponds to lines 129-132 in function DepositHelper of core.go
		return calculateSharesEmpty(amount0, amount1, price)
	} else {
		// these cases correspond to lines 228-234 in function DepositHelper of core.go
		if trueAmount0.GT(sdk.ZeroDec()) {
			return calculateSharesNonEmpty(trueAmount0, upperReserve0, upperTotalShares)
		} else {
			return calculateSharesNonEmpty(trueAmount1, lowerReserve1, upperTotalShares)
		}
	}
}

func calculateShares(amount0 sdk.Dec, amount1 sdk.Dec, pairId string, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) sdk.Dec {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx

	price, err := k.Calc_price(tickIndex, false)
	if err != nil {
		t.Errorf("TODO: calc price error format")
	}

	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee

	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-fee)
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+fee)
	lowerReserve1 := lowerTick.TickData.Reserve1[feeIndex]
	upperReserve0, upperTotalShares := upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares

	trueAmount0, trueAmount1 := amount0, amount1
	if upperReserve0.GT(sdk.ZeroDec()) {
		// this corresponds to lines 217-221  in function DepositHelper of core.go
		trueAmount1 = k.Min(amount1, lowerReserve1.Mul(amount0).Quo(upperReserve0))
	}
	if lowerReserve1.GT(sdk.ZeroDec()) {
		// this corresponds to lines 223-226 in function DepositHelper of core.go
		trueAmount0 = k.Min(amount0, upperReserve0.Mul(amount1).Quo(lowerReserve1))
	}

	return calculateSharesPure(
		amount0,
		trueAmount0,
		amount1,
		trueAmount1,
		price,
		feeIndex,
		lowerTickFound,
		lowerReserve1,
		upperTickFound,
		upperReserve0,
		upperTotalShares,
	)
}

func makePairId(coinA sdk.Coin, coinB sdk.Coin, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) string {
	// TODO: this really needs to be cleaned up
	app, ctx, goCtx, k := env.cosmos.app, env.cosmos.ctx, sdk.WrapSDKContext(env.cosmos.ctx), env.cosmos.app.DexKeeper
	token0, token1, err := k.SortTokens(ctx, coinA.Denom, coinB.Denom)
	if err != nil {
		t.Errorf("TODO: token sort error")
	}
	feelist := k.GetAllFeeList(ctx)
	pairId, err := app.DexKeeper.PairInit(goCtx, token0, token1, tickIndex, feelist[feeIndex].Fee)
	if err != nil {
		t.Errorf("TODO: pairId error format")
	}
	return pairId
}

func testSingleDeposit(t *testing.T, coinA sdk.Coin, coinB sdk.Coin, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64, env *TestEnv) {
	// TODO: modify this for len(tickIndexes) > 1
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
	amount0, amount1 := sdk.NewDecFromIntWithPrec(coinA.Amount, 18), sdk.NewDecFromIntWithPrec(coinB.Amount, 18)
	_, err := env.cosmos.msgServer.Deposit(goCtx, &types.MsgDeposit{ // (discard message response because we don't need it)
		Creator:     acc.String(),
		TokenA:      coinA.Denom,
		TokenB:      coinB.Denom,
		AmountsA:    []sdk.Dec{amount0}, // Coin is already denominated in 1e18
		AmountsB:    []sdk.Dec{amount1},
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
	pairId := makePairId(coinA, coinB, tickIndexes[0], feeTiers[0], t, env)
	accShares := calculateShares(amount0, amount1, pairId, tickIndexes[0], feeTiers[0], t, env)
	mintedShares, found := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])
	if !found {
		t.Errorf("Shares resulting from deposit by %s have not been minted (not found by getter).", acc)
	}
	if !(mintedShares.SharesOwned.Equal(accShares)) {
		t.Errorf("Incorrect amount of shares minted after deposit by %s of %s, %s. Needed %s, minted %s", acc, coinA, coinB, accShares, mintedShares.SharesOwned)
	}
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
