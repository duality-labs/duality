package e2e_test

import (
	"encoding/json"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"
	"github.com/duality-labs/duality/app"
	appConsumer "github.com/duality-labs/duality/app"
)

// DualityAppIniter implements ibctesting.AppIniter for the duality consumer app
func DualityAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := app.MakeEncodingConfig()
	testApp := appConsumer.NewApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		simtestutil.EmptyAppOptions{},
		encoding,
	)

	return testApp, appConsumer.NewDefaultGenesisState(encoding.Marshaler)
}
