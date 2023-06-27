package e2e_test

// import (
// 	"encoding/json"

// 	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"

// 	cmdb "github.com/cometbft/cometbft-db"
// 	"github.com/cometbft/cometbft/libs/log"
// 	"github.com/tendermint/spm/cosmoscmd"

// 	"github.com/duality-labs/duality/app"
// 	appConsumer "github.com/duality-labs/duality/app"
// )

// // DualityAppIniter implements ibctesting.AppIniter for the duality consumer app
// func DualityAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
// 	encoding := cosmoscmd.MakeEncodingConfig(appConsumer.ModuleBasics)
// 	testApp := appConsumer.New(log.NewNopLogger(), cmdb.NewMemDB(), nil, true, map[int64]bool{},
// 		app.DefaultNodeHome, 5, encoding, app.EmptyAppOptions{}).(ibctesting.TestingApp)

// 	return testApp, appConsumer.NewDefaultGenesisState(encoding.Marshaler)
// }
