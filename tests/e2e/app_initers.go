package e2e_test

import (
	"encoding/json"

	"cosmossdk.io/simapp"

	"github.com/cometbft/cometbft/libs/log"
	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"
	appparams "github.com/duality-labs/duality/app/params"
	tmdb "github.com/tendermint/tm-db"

	appConsumer "github.com/duality-labs/duality/app"
)

// DualityAppIniter implements ibctesting.AppIniter for the duality consumer app
func DualityAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := appparams.MakeTestEncodingConfig(appConsumer.ModuleBasics)
	testApp := appConsumer.New(log.NewNopLogger(), tmdb.NewMemDB(), nil, true, map[int64]bool{},
		simapp.DefaultNodeHome, 5, encoding, simapp.EmptyAppOptions{}).(ibctesting.TestingApp)

	return testApp, appConsumer.NewDefaultGenesisState(encoding.Marshaler)
}
