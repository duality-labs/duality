package keeper_test

import (
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func newCCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenC", amt)
}

func newDCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenD", amt)
}

// TODO: better name
type CosmosTestEnv struct {
	app       *dualityapp.App
	msgServer types.MsgServer
	// TODO: keeping ctx in struct is bad practice: https://pkg.go.dev/context#pkg-overview
	ctx              sdk.Context
	queryClient      types.QueryClient
	dexModuleAddress string
}

func cosmosEnvSetup() CosmosTestEnv {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	dexModuleAddress := app.AccountKeeper.GetModuleAddress(types.ModuleName).String()

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DexKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	msgServer := keeper.NewMsgServerImpl(app.DexKeeper)

	return CosmosTestEnv{
		app,
		msgServer,
		ctx,
		queryClient,
		dexModuleAddress,
	}
}

type TestEnv struct {
	cosmos          CosmosTestEnv
	addrs           []sdk.AccAddress
	balances        map[string]sdk.Coins
	feeTiers        []types.FeeList
	intentionalFail bool
}

func EnvSetup(t *testing.T, intentionalFail bool) TestEnv {
	cosmos := cosmosEnvSetup()
	app, ctx := cosmos.app, cosmos.ctx

	// initialize accounts
	alice, bob := sdk.AccAddress([]byte("alice")), sdk.AccAddress([]byte("bob"))
	accAlice, accBob := app.AccountKeeper.NewAccountWithAddress(ctx, alice), app.AccountKeeper.NewAccountWithAddress(ctx, bob)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	app.AccountKeeper.SetAccount(ctx, accBob)

	// init balances & fund the accounts
	balancesAlice := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("20000000000000000000")), newCCoin(convInt("30000000000000000000")), newDCoin(convInt("40000000000000000000")))
	balancesBob := sdk.NewCoins(newACoin(convInt("10000000000000000000")), newBCoin(convInt("20000000000000000000")), newCCoin(convInt("30000000000000000000")), newDCoin(convInt("40000000000000000000")))
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

// Handle checking for intentional failure of test
func (env *TestEnv) handleIntentionalFail(t *testing.T, format string, args ...interface{}) {
	if !env.intentionalFail {
		t.Fatalf(format, args...)
	} else {
		t.Skipf("Test intentionally failed, skipping rest of execution. Error: "+format, args)
	}
}

// Helper to convert coins into sorted amount0, amount1
func (env *TestEnv) sortCoins(t *testing.T, denomA string, denomB string, amountsA []sdk.Dec, amountsB []sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec) {
	app, ctx := env.cosmos.app, env.cosmos.ctx
	denom0, denom1, err := app.DexKeeper.SortTokens(ctx, denomA, denomB)
	if err != nil {
		t.Errorf("Failed to sort coins %s, %s", denomA, denomB)
	}
	// this corresponds to lines 45-54 of verification.go
	amounts0, amounts1 := amountsA, amountsB
	// flip amounts if denoms were flipped
	if denom0 != denomA {
		amounts0, amounts1 = amountsB, amountsA
	}
	return denom0, denom1, amounts0, amounts1
}

// TODO: this was taken from core.go, lines 287-294. should be moved to utils somewhere
func min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}
