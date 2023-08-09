package e2e_test

import (
	"encoding/json"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"
	icsappiniters "github.com/cosmos/interchain-security/v3/testutil/ibc_testing"
	"github.com/duality-labs/duality/app"
	appparams "github.com/duality-labs/duality/app/params"
)

// DualityAppIniter implements ibctesting.AppIniter for the duality consumer app
func DualityAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	// Reset to duality config before initializing duality app
	appparams.InitParams()
	encoding := app.MakeEncodingConfig()
	testApp := app.NewApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		simtestutil.EmptyAppOptions{},
		encoding,
		nil,
	)

	return testApp, app.NewDefaultGenesisState(encoding.Marshaler)
}

func ProviderAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	// Reset to cosmos sdk base config before initializing provider app
	sdk.GetConfig()
	return icsappiniters.ProviderAppIniter()
}
