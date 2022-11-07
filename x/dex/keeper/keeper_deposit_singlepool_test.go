package keeper_test

import (
	// stdlib

	"testing"

	// cosmos SDK
	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
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
	cosmos          CosmostTestEnv
	addrs           []sdk.AccAddress
	balances        map[string]sdk.Coins
	feeTiers        []types.FeeList
	intentionalFail bool
}

// handle checking for intentional failure of test
func (env *TestEnv) checkIntentionalFail(condition bool) bool {
	if condition && !env.intentionalFail {
		return true
	} else {
		return false
	}
}

func singlePoolSetup(t *testing.T, cosmos CosmostTestEnv, intentionalFail bool) TestEnv {
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
	if err := (FundAccount(app.BankKeeper, ctx, alice, balancesAlice)); err != nil {
		t.Errorf("Failed to fund %s with %s", alice, balancesAlice)
	}
	if err := (FundAccount(app.BankKeeper, ctx, bob, balancesBob)); err != nil {
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
		intentionalFail,
	}
}

// Helpers for calculating the amount of shares that should be minted
// These are here for the sake of more concise unit tests and the corresponding code in core.go
// should eventually be refactored so that core.go is modularized for easier unit testing and readability

// Calculation of shares when depositing the initial amount (no reserves)
func calculateSharesEmpty(amount0 sdk.Dec, amount1 sdk.Dec, price sdk.Dec) sdk.Dec {
	return amount0.Add(amount1.Mul(price))
}

// Calculation of shares when there are pre-existing reserves
func calculateSharesNonEmpty(amount sdk.Dec, reserve sdk.Dec, totalShares sdk.Dec) sdk.Dec {
	return amount.Quo(reserve).Mul(totalShares)
}

// Pure func that takes all the parameters requires to compute the amount of minted shares and handles the different cases accordingly.
// This is probably excessive as keeping only the calculation pure is reasonable enough, but it's here for posterity.
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

// Impure function that pulls all the state variables required for calculating the amount of shares to mint.
func calculateShares(amount0 sdk.Dec, amount1 sdk.Dec, pairId string, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) sdk.Dec {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx

	price, err := k.Calc_price(tickIndex, false)
	if env.checkIntentionalFail(err != nil) {
		t.Errorf("TODO: calc price error format")
	}

	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee

	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-fee)
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+fee)
	// TODO: this won't work if not found
	lowerReserve1 := lowerTick.TickData.Reserve1[feeIndex]
	upperReserve0, upperTotalShares := upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares

	trueAmount0, trueAmount1 := amount0, amount1
	if upperReserve0.GT(sdk.ZeroDec()) {
		// this corresponds to lines 217-221 in function DepositHelper of core.go
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

// Helper for getting a pair id
func makePairId(coinA sdk.Coin, coinB sdk.Coin, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) string {
	// TODO: this really should be cleaned up
	app, ctx, goCtx, k := env.cosmos.app, env.cosmos.ctx, sdk.WrapSDKContext(env.cosmos.ctx), env.cosmos.app.DexKeeper
	token0, token1, err := k.SortTokens(ctx, coinA.Denom, coinB.Denom)
	if env.checkIntentionalFail((err != nil)) {
		t.Errorf("TODO: token sort error")
	}

	// this corresponds to line 16 in function DepositVerification of verification.go
	feelist := k.GetAllFeeList(ctx)
	// handle invalid fee index
	if env.checkIntentionalFail(feeIndex > uint64(len(feelist))) {
		t.Errorf("Fee index (%d) > fee tier count (%d)", feeIndex, len(feelist))
	}

	// this corresponds to line 304 in function DepositCore of core.go
	// TODO: this might be wrong?
	pairId, err := app.DexKeeper.PairInit(goCtx, token0, token1, tickIndex, feelist[feeIndex].Fee)
	if env.checkIntentionalFail(err != nil) {
		t.Errorf("TODO: pairId error format")
	}

	return pairId
}

// Template function for testing the execution of deposit.
func testSingleDeposit(t *testing.T, coinA sdk.Coin, coinB sdk.Coin, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64, env *TestEnv) {
	// TODO: modify this for len(tickIndexes) > 1, i.e. testMultipleDeposits
	app, ctx := env.cosmos.app, env.cosmos.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// GIVEN inital balances
	accBalanceAInitial, accBalanceBInitial := newACoin(env.balances[acc.String()].AmountOf(coinA.Denom)), newBCoin(env.balances[acc.String()].AmountOf(coinB.Denom))
	// verify acc has exactly the balance passed in from env
	if env.checkIntentionalFail(!(app.BankKeeper.GetBalance(ctx, acc, coinA.Denom).IsEqual(accBalanceAInitial))) {
		t.Errorf("%s's initial balance of %s does not match env: %s", acc, coinA, accBalanceAInitial)
	}
	if env.checkIntentionalFail(!(app.BankKeeper.GetBalance(ctx, acc, coinB.Denom).IsEqual(accBalanceBInitial))) {
		t.Errorf("%s's initial balance of %s does not match env: %s", acc, coinB, accBalanceBInitial)
	}
	// get Dex initial balance
	dexAllCoinsInitial := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAInitial, dexBalanceBInitial := newACoin(dexAllCoinsInitial.AmountOf(coinA.Denom)), newBCoin(dexAllCoinsInitial.AmountOf(coinB.Denom))
	// get amount of shares before depositing
	pairId := makePairId(coinA, coinB, tickIndexes[0], feeTiers[0], t, env)
	initialShares, initialSharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])

	// WHEN depositing the specified amounts coinA and coinB
	// TODO: need to sort coinA, coinB
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
	if env.checkIntentionalFail(err != nil) {
		t.Errorf("Deposit of %s, %s by %s failed: %s", coinA, coinB, acc, err)
	}

	// verify alice's resulting balances is aliceBalanceInitial - depositCoin
	accBalanceAFinal, accBalanceBFinal := accBalanceAInitial.Sub(coinA), accBalanceBInitial.Sub(coinB)
	if env.checkIntentionalFail(!(app.BankKeeper.GetBalance(ctx, acc, coinA.Denom).IsEqual(accBalanceAFinal))) {
		t.Errorf("%s's final balance of %s does not reflect deposit", acc, coinA)
	}
	if env.checkIntentionalFail(!(app.BankKeeper.GetBalance(ctx, acc, coinB.Denom).IsEqual(accBalanceBFinal))) {
		t.Errorf("%s's final balance of %s does not reflect deposit", acc, coinB)
	}

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexAllCoinsFinal := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAFinal, dexBalanceBFinal := dexBalanceAInitial.Add(coinA), dexBalanceBInitial.Add(coinB)
	if env.checkIntentionalFail(!(newACoin(dexAllCoinsFinal.AmountOf(coinA.Denom)).IsEqual(dexBalanceAFinal))) {
		t.Errorf("Dex module's final balance of %s does not reflect deposit", coinA.Denom)
	}
	if env.checkIntentionalFail(!(newBCoin(dexAllCoinsFinal.AmountOf(coinB.Denom)).IsEqual(dexBalanceBFinal))) {
		t.Errorf("Dex module's final balance of %s does not reflect deposit", coinB.Denom)
	}

	// verify amount of shares minted for alice
	accSharesCalc := calculateShares(amount0, amount1, pairId, tickIndexes[0], feeTiers[0], t, env)
	finalShares, found := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])
	if env.checkIntentionalFail(!found) {
		t.Errorf("Shares resulting from deposit by %s have not been minted (not found by getter).", acc)
	} else if env.checkIntentionalFail(!initialSharesFound && !(finalShares.SharesOwned.Equal(accSharesCalc))) {
		// Handle the case when no shares held by account initially but mintedShares != accSharesCalc
		t.Errorf("Incorrect amount of shares minted after deposit by %s of %s, %s. Needed %s, final %s", acc, coinA, coinB, accSharesCalc, finalShares.SharesOwned)
	} else if env.checkIntentionalFail(initialSharesFound && !finalShares.SharesOwned.Equal(initialShares.SharesOwned.Add(accSharesCalc))) {
		// Handle the case when account had an initial balance of shares but finalShares != initalShares + accSharesCalc
		t.Errorf("Incorrect amount of shares minted after deposit by %s of %s, %s. Needed %s, final %s", acc, coinA, coinB, initialShares.SharesOwned.Add(accSharesCalc), finalShares.SharesOwned)
	}

	// verify fee tier of minted shares
	if env.checkIntentionalFail(finalShares.FeeIndex != feeTiers[0]) {
		t.Errorf("Shares minted in the wrong fee tier. Needed %d, final %d", feeTiers[0], finalShares.FeeIndex)
	}
}

// func TestMinFeeTier(t *testing.T) {
// 	fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/MinFeeTier")

// 	// GIVEN initial balances and fee tiers from the setup
// 	env := singlePoolSetup(t, cosmosEnvSetup(), false)

// 	// WHEN alice deposits her genesis balance of tokenA and tokenB into the minimal fee tier
// 	// prep deposit args
// 	acc := env.addrs[0]
// 	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

// 	// deposit with min fee tier
// 	tickIndex := []int64{0}
// 	minFeeTier := []uint64{uint64(env.feeTiers[0].Id)} // grab the index of minimal fee tier

// 	testSingleDeposit(t, coinA, coinB, acc, tickIndex, minFeeTier, &env)

// 	// THEN the transaction should execute successfully
// }

// func TestMaxFeeTier(t *testing.T) {
// 	fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/MaxFeeTier")

// 	// GIVEN initial balances and fee tiers from the setup
// 	env := singlePoolSetup(t, cosmosEnvSetup(), false)

// 	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
// 	// prep deposit args
// 	acc := env.addrs[0]
// 	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

// 	// deposit with max fee tier
// 	tickIndex := []int64{0}
// 	maxFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Id)} // grab the index of max fee tier

// 	testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)

// 	// THEN the transaction should execute successfully
// }

func testFeeTier(t *testing.T, name string, tier uint64, intentionalFail bool) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/%s", name)

	// GIVEN initial balances and fee tiers from the setup
	env := singlePoolSetup(t, cosmosEnvSetup(), intentionalFail)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	invalidFeeTier := []uint64{tier}

	testSingleDeposit(t, coinA, coinB, acc, tickIndex, invalidFeeTier, &env)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestFeeTiers(t *testing.T) {
	tests := []struct {
		name            string
		tier            uint64
		intentionalFail bool
	}{
		{"Min Tier", 0, false},
		{"Max Tier", 2, false}, // amount of tiers added in setupDepositTest (line 92)
		// {"Invalid Tier", 3, true}, // Max Tier + 1
	}
	for _, tt := range tests {
		testFeeTier(t, tt.name, tt.tier, tt.intentionalFail)
	}
}

func TestInitPair(t *testing.T) {
	// fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/InitPair")

	// GIVEN initial balances and fee tiers from the setup, i.e. no liquidity has been deposited
	// env := singlePoolSetup(t, cosmosEnvSetup())
	// // grab list of pairs for comparison

	// // WHEN alice deposits her setup balance of tokenA and tokenB into some fee tier (shouldn't matter)
	// // prep deposit args
	// acc := env.addrs[0]
	// coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	// tickIndex := []int64{0}
	// maxFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Id)}

	// testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)

	// THEN the transaction should execute successfully, i.e. the new pair has liquidity
	// verify new pair added
}

func TestInitTick(t *testing.T) {
	// fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/InitTick")

	// // GIVEN initial balances and fee tiers from the setup, i.e. no liquidity has been deposited
	// env := singlePoolSetup(t, cosmosEnvSetup())
	// // grab list of ticks for pair

	// // WHEN alice deposits her setup balance of tokenA and tokenB into some fee tier (shouldn't matter)
	// // prep deposit args
	// acc := env.addrs[0]
	// coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	// maxFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Id)}

	// // deposit into new tick
	// // TODO
	// tickIndex := []int64{0}

	// testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)

	// THEN the transaction should execute successfully, i.e. the pair has liquidity in the new tick
	// verify new tick added
	// verify tick balance
}

func TestInitFeeTier(t *testing.T) {
	// fmt.Println("[ UnitTests|Keeper ] Starting test: SinglePool/InitFeeTier")

	// // GIVEN initial balances and fee tiers from the setup, i.e. no liquidity has been deposited
	// env := singlePoolSetup(t, cosmosEnvSetup())
	// // grab list of fee tiers for pair

	// // WHEN alice deposits her setup balance of tokenA and tokenB into a new fee tier
	// // prep deposit args
	// acc := env.addrs[0]
	// coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// // deposit with min fee tier
	// tickIndex := []int64{0}
	// maxFeeTier := []uint64{uint64(env.feeTiers[len(env.feeTiers)-1].Id)}

	// testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)

	// THEN the transaction should execute successfully, i.e. the pair has liquidity in the new fee tier
	// verify new fee tier added
	// verify liquidity at tick corresponding to fee tier
}
