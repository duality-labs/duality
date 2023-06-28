package e2e_test

import (
	"encoding/json"

	"cosmossdk.io/simapp"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	ibctesting "github.com/cosmos/interchain-security/legacy_ibc_testing/testing"
	"github.com/duality-labs/duality/app"
	appparams "github.com/duality-labs/duality/app/params"

	appConsumer "github.com/duality-labs/duality/app"
)

// DualityAppIniter implements ibctesting.AppIniter for the duality consumer app
func DualityAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := appparams.MakeTestEncodingConfig()
	testApp := appConsumer.NewApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, map[int64]bool{},
		simapp.DefaultNodeHome, 5, encoding, app.EmptyAppOptions{})

	return testApp, appConsumer.NewDefaultGenesisState(encoding.Marshaler)
}
