package app_test

import (
	"encoding/json"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	ibctesting "github.com/cosmos/interchain-security/v3/legacy_ibc_testing/testing"
	"github.com/duality-labs/duality/app"
	"github.com/stretchr/testify/require"
)

func TestConsumerWhitelistingKeys(t *testing.T) {
	chain := ibctesting.NewTestChain(t, ibctesting.NewCoordinator(t, 0), SetupTestingAppConsumer, "test")
	paramKeeper := chain.App.(*app.App).ParamsKeeper
	for paramKey := range app.WhitelistedParams {
		ss, ok := paramKeeper.GetSubspace(paramKey.Subspace)
		require.True(t, ok, "Unknown subspace %s", paramKey.Subspace)
		hasKey := ss.Has(chain.GetContext(), []byte(paramKey.Key))
		require.True(t, hasKey, "Invalid key %s for subspace %s", paramKey.Key, paramKey.Subspace)
	}
}

func SetupTestingAppConsumer() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := app.MakeTestEncodingConfig()
	testApp := app.NewApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, encCdc, app.EmptyAppOptions{})

	return testApp, app.NewDefaultGenesisState(encCdc.Marshaler)
}
