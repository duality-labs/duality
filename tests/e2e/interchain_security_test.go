package e2e_test

import (
	"encoding/json"
	"testing"

	appConsumer "github.com/NicholasDotSol/duality/app"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	appProvider "github.com/cosmos/interchain-security/app/provider"
	"github.com/cosmos/interchain-security/tests/e2e"
	icstestingutils "github.com/cosmos/interchain-security/testutil/ibc_testing"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// Executes the standard group of ccv tests against a consumer and provider app.go implementation.
func TestCCVTestSuite(t *testing.T) {

	ccvSuite := e2e.NewCCVTestSuite[*appProvider.App, *appConsumer.App](
		icstestingutils.ProviderAppIniter,
		SetupTestingAppConsumer,
		[]string{},
	)
	suite.Run(t, ccvSuite)
}

// NewCoordinator initializes Coordinator with interchain security dummy provider and duality consumer chain
func NewProviderConsumerCoordinator(t *testing.T) (*ibctesting.Coordinator, *ibctesting.TestChain, *ibctesting.TestChain) {
	coordinator := ibctesting.NewCoordinator(t, 0)
	chainID := ibctesting.GetChainID(1)
	coordinator.Chains[chainID] = ibctesting.NewTestChain(t, coordinator, icstestingutils.ProviderAppIniter, chainID)
	providerChain := coordinator.GetChain(chainID)
	chainID = ibctesting.GetChainID(2)
	coordinator.Chains[chainID] = ibctesting.NewTestChainWithValSet(t, coordinator,
		SetupTestingAppConsumer, chainID, providerChain.Vals, providerChain.Signers)
	consumerChain := coordinator.GetChain(chainID)
	return coordinator, providerChain, consumerChain
}

func SetupTestingAppConsumer() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := appConsumer.MakeTestEncodingConfig()
	app := appConsumer.NewApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		appConsumer.DefaultNodeHome,
		5,
		encCdc,
		appConsumer.EmptyAppOptions{})

	return app, appConsumer.NewDefaultGenesisState(encCdc.Marshaler)
}
