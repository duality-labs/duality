package e2e_test

import (
	"testing"

	appprovider "github.com/cosmos/interchain-security/v3/app/provider"
	"github.com/cosmos/interchain-security/v3/tests/integration"

	"github.com/duality-labs/duality/app"
	"github.com/stretchr/testify/suite"
)

// This file can be used as an example e2e testing instance for any provider/consumer applications.
// In the case of this repo, we're testing the dummy provider/consumer applications,
// but to test any arbitrary app, one only needs to replicate this file and "specific_setup.go",
// then pass in the appropriate types and parameters to the suite. Note that provider and consumer
// applications types must implement the interfaces defined in /testutil/e2e/interfaces.go to compile.

// Executes the standard group of ccv tests against a consumer and provider app.go implementation.
func TestCCVTestSuite(t *testing.T) {
	// Pass in concrete app types that implement the interfaces defined in /testutil/e2e/interfaces.go
	ccvSuite := integration.NewCCVTestSuite[*appprovider.App, *app.App](
		// Pass in ibctesting.AppIniters for provider and consumer.
		ProviderAppIniter, DualityAppIniter, []string{})

	// Run tests
	suite.Run(t, ccvSuite)
}
