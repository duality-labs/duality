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
func (env *TestEnv) handleIntentionalFail(t *testing.T, format string, args ...interface{}) {
	if !env.intentionalFail {
		t.Errorf(format, args...)
	} else {
		t.Skipf("Test intentionally failed, skipping rest of execution. Error: "+format, args)
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

// TODO: this was taken from core.go, lines 287-294. should be moved to utils somewhere
func min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

// Helper to convert coins into sorted amount0, amount1
func (env *TestEnv) sortCoins(t *testing.T, coinA sdk.Coin, coinB sdk.Coin) (sdk.Dec, sdk.Dec) {
	app, ctx := env.cosmos.app, env.cosmos.ctx
	denom0, denom1, err := app.DexKeeper.SortTokens(ctx, coinA.Denom, coinB.Denom)
	if err != nil {
		t.Errorf("Failed to sort coins %s, %s", coinA, coinB)
	}
	coins := sdk.NewCoins(coinA, coinB)
	return sdk.NewDecFromIntWithPrec(coins.AmountOf(denom0), 18), sdk.NewDecFromIntWithPrec(coins.AmountOf(denom1), 18)

}

// Helper function to balance amounts to pool ratio
func trueAmounts(amount0 sdk.Dec, amount1 sdk.Dec, lowerReserve1 sdk.Dec, upperReserve0 sdk.Dec) (sdk.Dec, sdk.Dec) {
	trueAmount0, trueAmount1 := amount0, amount1
	if upperReserve0.GT(sdk.ZeroDec()) {
		// this corresponds to lines 217-221 in function DepositHelper of core.go
		trueAmount1 = min(amount1, lowerReserve1.Mul(amount0).Quo(upperReserve0))
	}
	if lowerReserve1.GT(sdk.ZeroDec()) {
		// this corresponds to lines 223-226 in function DepositHelper of core.go
		trueAmount0 = min(amount0, upperReserve0.Mul(amount1).Quo(lowerReserve1))
	}
	return trueAmount0, trueAmount1
}

// Impure function that pulls all the state variables required for calculating the amount of shares to mint.
func calculateShares(amount0 sdk.Dec, amount1 sdk.Dec, pairId string, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) sdk.Dec {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx

	price, err := k.Calc_price(tickIndex, false)
	if err != nil {
		env.handleIntentionalFail(t, "TODO: calc price error format")
	}

	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee

	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-fee)
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+fee)
	lowerReserve1 := lowerTick.TickData.Reserve1[feeIndex]
	upperReserve0, upperTotalShares := upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares

	trueAmount0, trueAmount1 := trueAmounts(amount0, amount1, lowerReserve1, upperReserve0)

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

// Helper function to calculate if current ticks change
func calculateNewCurrentTicksPure(amount0 sdk.Dec, amount1 sdk.Dec, tickIndex int64, fee int64, curr0to1 int64, curr1to0 int64) (int64, int64) {
	// this corresponds to lines 245-253 in function DepositHelper of core.go
	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	new0to1, new1to0 := curr0to1, curr1to0
	if amount0.GT(sdk.ZeroDec()) && ((tickIndex+fee > curr0to1) && (tickIndex+fee < curr1to0)) {
		new1to0 = tickIndex + fee
	}
	if amount1.GT(sdk.ZeroDec()) && ((tickIndex-fee > curr0to1) && (tickIndex-fee < curr1to0)) {
		new0to1 = tickIndex - fee
	}
	return new0to1, new1to0
}

func (env *TestEnv) calculateNewCurrentTicks(amount0 sdk.Dec, amount1 sdk.Dec, tickIndex int64, feeIndex uint64, pair types.PairMap) (new0to1 int64, new1to0 int64) {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx
	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee
	return calculateNewCurrentTicksPure(amount0, amount1, tickIndex, fee, pair.TokenPair.CurrentTick0To1, pair.TokenPair.CurrentTick1To0)
}

// Helper for getting a pair id
func makePairId(coinA sdk.Coin, coinB sdk.Coin, tickIndex int64, feeIndex uint64, t *testing.T, env *TestEnv) string {
	// TODO: this really should be cleaned up
	app, ctx, goCtx, k := env.cosmos.app, env.cosmos.ctx, sdk.WrapSDKContext(env.cosmos.ctx), env.cosmos.app.DexKeeper
	token0, token1, err := k.SortTokens(ctx, coinA.Denom, coinB.Denom)
	if err != nil {
		env.handleIntentionalFail(t, "TODO: token sort error")
	}

	// this corresponds to line 16 in function DepositVerification of verification.go
	feelist := k.GetAllFeeList(ctx)
	// handle invalid fee index
	if feeIndex >= uint64(len(feelist)) {
		env.handleIntentionalFail(t, "Fee index (%d) > fee tier count (%d)", feeIndex, len(feelist))
	}

	// this corresponds to line 304 in function DepositCore of core.go
	// TODO: this might be wrong?
	pairId, err := app.DexKeeper.PairInit(goCtx, token0, token1, tickIndex, feelist[feeIndex].Fee)
	if err != nil {
		env.handleIntentionalFail(t, "TODO: pairId error format")
	}

	return pairId
}

// Template function for testing the execution of deposit.
func testSingleDeposit(t *testing.T, coinA sdk.Coin, coinB sdk.Coin, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64, env *TestEnv) {
	// TODO: modify this for len(tickIndexes) > 1, i.e. testMultipleDeposits
	app, ctx := env.cosmos.app, env.cosmos.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// GIVEN inital balances
	accBalanceAInitial, accBalanceBInitial := app.BankKeeper.GetBalance(ctx, acc, coinA.Denom), app.BankKeeper.GetBalance(ctx, acc, coinB.Denom)
	// verify acc has sufficient balances in env
	if !(accBalanceAInitial.IsGTE(coinA)) {
		env.handleIntentionalFail(t, "%s has insufficient balance. Has %s, expected at most %s", acc, coinA, accBalanceAInitial)
	}
	if !(accBalanceBInitial.IsGTE(coinB)) {
		env.handleIntentionalFail(t, "%s has insufficient balance. Has %s, expected at most %s", acc, coinB, accBalanceBInitial)
	}
	// get Dex initial balance
	dexAllCoinsInitial := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAInitial, dexBalanceBInitial := newACoin(dexAllCoinsInitial.AmountOf(coinA.Denom)), newBCoin(dexAllCoinsInitial.AmountOf(coinB.Denom))
	// get amount of shares before depositing
	pairId := makePairId(coinA, coinB, tickIndexes[0], feeTiers[0], t, env)
	initialShares, initialSharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])

	// get initial pair info for 0to1 and 1to0 ticks
	pairInitial, pairInitialFound := app.DexKeeper.GetPairMap(ctx, pairId)
	if !pairInitialFound {
		env.handleIntentionalFail(t, "TODO: handle pairInitial not found")
	}

	// WHEN depositing the specified amounts coinA and coinB
	amount0, amount1 := env.sortCoins(t, coinA, coinB)
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
		env.handleIntentionalFail(t, "Deposit of %s, %s by %s failed: %s", coinA, coinB, acc, err)
	}

	// verify alice's resulting balances is aliceBalanceInitial - depositCoin
	accBalanceAFinal, accBalanceBFinal := accBalanceAInitial.Sub(coinA), accBalanceBInitial.Sub(coinB)
	if !(app.BankKeeper.GetBalance(ctx, acc, coinA.Denom).IsEqual(accBalanceAFinal)) {
		env.handleIntentionalFail(t, "%s's final balance of %s does not reflect deposit", acc, coinA)
	}
	if !(app.BankKeeper.GetBalance(ctx, acc, coinB.Denom).IsEqual(accBalanceBFinal)) {
		env.handleIntentionalFail(t, "%s's final balance of %s does not reflect deposit", acc, coinB)
	}

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexAllCoinsFinal := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAFinal, dexBalanceBFinal := dexBalanceAInitial.Add(coinA), dexBalanceBInitial.Add(coinB)
	if !(newACoin(dexAllCoinsFinal.AmountOf(coinA.Denom)).IsEqual(dexBalanceAFinal)) {
		env.handleIntentionalFail(t, "Dex module's final balance of %s does not reflect deposit", coinA.Denom)
	}
	if !(newBCoin(dexAllCoinsFinal.AmountOf(coinB.Denom)).IsEqual(dexBalanceBFinal)) {
		env.handleIntentionalFail(t, "Dex module's final balance of %s does not reflect deposit", coinB.Denom)
	}

	// verify amount of shares minted for alice
	accSharesCalc := calculateShares(amount0, amount1, pairId, tickIndexes[0], feeTiers[0], t, env)
	finalShares, found := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[0], feeTiers[0])
	if !found {
		env.handleIntentionalFail(t, "Shares resulting from deposit by %s have not been minted (not found by getter).", acc)
	} else if !initialSharesFound && !(finalShares.SharesOwned.Equal(accSharesCalc)) {
		// Handle the case when no shares held by account initially but mintedShares != accSharesCalc
		env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit by %s of %s, %s. Needed %s, final %s", acc, coinA, coinB, accSharesCalc, finalShares.SharesOwned)
	} else if initialSharesFound && !finalShares.SharesOwned.Equal(initialShares.SharesOwned.Add(accSharesCalc)) {
		// Handle the case when account had an initial balance of shares but finalShares != initalShares + accSharesCalc
		env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit by %s of %s, %s. Needed %s, final %s", acc, coinA, coinB, initialShares.SharesOwned.Add(accSharesCalc), finalShares.SharesOwned)
	}

	// verify fee tier of minted shares
	if finalShares.FeeIndex != feeTiers[0] {
		env.handleIntentionalFail(t, "Shares minted in the wrong fee tier. Needed %d, final %d", feeTiers[0], finalShares.FeeIndex)
	}

	// verify current ticks set properly
	tick0to1Calc, tick1to0Calc := env.calculateNewCurrentTicks(amount0, amount1, tickIndexes[0], feeTiers[0], pairInitial)
	pairFinal, pairInitialFound := app.DexKeeper.GetPairMap(ctx, pairId)
	if !pairInitialFound {
		env.handleIntentionalFail(t, "TODO: handle pairFinal not found")
	}
	if pairFinal.TokenPair.CurrentTick0To1 != tick0to1Calc {
		env.handleIntentionalFail(t, "Invalid CurrentTick0To1 resulted from deposit. Needed %d, final %d", tick0to1Calc, pairFinal.TokenPair.CurrentTick0To1)
	}
	if pairFinal.TokenPair.CurrentTick1To0 != tick1to0Calc {
		env.handleIntentionalFail(t, "Invalid CurrentTick1To0 resulted from deposit. Needed %d, final %d", tick1to0Calc, pairFinal.TokenPair.CurrentTick1To0)
	}
}

func TestMinFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/MinFee")

	// GIVEN initial balances and fee tiers from the setup
	env := singlePoolSetup(t, cosmosEnvSetup(), false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	minFeeTier := []uint64{0}

	testSingleDeposit(t, coinA, coinB, acc, tickIndex, minFeeTier, &env)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestMaxFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/MaxFee")

	// GIVEN initial balances and fee tiers from the setup
	env := singlePoolSetup(t, cosmosEnvSetup(), false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	maxFeeTier := []uint64{uint64(len(env.feeTiers) - 1)}

	testSingleDeposit(t, coinA, coinB, acc, tickIndex, maxFeeTier, &env)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestInvalidFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/Invalid Fee Tier")

	// GIVEN initial balances and fee tiers from the setup
	env := singlePoolSetup(t, cosmosEnvSetup(), true)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	invalidFeeTier := []uint64{uint64(len(env.feeTiers))}

	testSingleDeposit(t, coinA, coinB, acc, tickIndex, invalidFeeTier, &env)

	// THEN the transaction should fail midway (SkipNow)
}

func TestInitPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitPair")

	// GIVEN initial balances and fee tiers from the setup
	env := singlePoolSetup(t, cosmosEnvSetup(), false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}

	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestInitTick(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitTick")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := singlePoolSetup(t, cosmosEnvSetup(), false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// WHEN alice deposits at tick 1 in fee tier 0
	newTickIndex := []int64{1}
	testSingleDeposit(t, coinA, coinB, acc, newTickIndex, feeTier, &env)

	// THEN the transaction should execute successfully
}

func TestInitFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitFeeTier")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := singlePoolSetup(t, cosmosEnvSetup(), false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// WHEN alice deposits at tick 0 in fee tier 1
	newFeeTier := []uint64{1}
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, newFeeTier, &env)

	// THEN the transaction should execute successfully
}

func TestExistingPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/ExistingPair")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := singlePoolSetup(t, cosmosEnvSetup(), false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// WHEN deposit in the same pair, tick and fee tier again
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// THEN the transaction should execute successfully
}

func TestBehindEnemyLines(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/BehindEnemyLines")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := singlePoolSetup(t, cosmosEnvSetup(), false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	testSingleDeposit(t, coinA, coinB, acc, tickIndex, feeTier, &env)

	// WHEN alice deposits at tick 0 in fee tier 1
	newTickIndex := []int64{-3}
	newFeeTier := []uint64{1}
	testSingleDeposit(t, coinA, coinB, acc, newTickIndex, newFeeTier, &env)

	// THEN the transaction should execute successfully
}
