[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/app)

The `.autodoc/docs/json/app` folder contains essential code files for the duality project, focusing on transaction validation, encoding configuration, exporting application state, and proposal whitelisting.

`ante_handler.go` is responsible for creating an AnteHandler middleware that validates transactions before they are processed by the blockchain. The AnteHandler is composed of a series of AnteDecorators, each performing specific validation tasks. The `NewAnteHandler` function takes in a `HandlerOptions` struct and returns the AnteHandler created by chaining together the AnteDecorators.

Example usage:

```go
import (
    "github.com/duality-labs/duality/app"
)

func main() {
    handlerOptions := app.HandlerOptions{...}
    anteHandler, err := app.NewAnteHandler(handlerOptions)
    // use anteHandler to validate transactions
}
```

`encoding.go` provides a deprecated function `MakeTestEncodingConfig()` for creating an `EncodingConfig` object for testing purposes. Instead, the `AppCodec` object should be used to create new codecs.

Example usage:

```go
import (
    "github.com/duality-labs/duality/app"
    "github.com/tendermint/spm/cosmoscmd"
)

func main() {
    encodingConfig := app.MakeTestEncodingConfig()
    codec := encodingConfig.Marshaler
    // use codec to encode and decode data structures
}
```

`export.go` contains the `App` struct with the `ExportAppStateAndValidators` method, which exports the state of the application for a genesis file. This method is crucial for exporting the application state and initializing the system.

`genesis.go` defines the `GenesisState` type, representing the initial state of the blockchain. The `NewDefaultGenesisState` function generates the default state by calling the `DefaultGenesis` function of the `ModuleBasics` object.

Example usage:

```go
package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/duality/app"
)

func main() {
	cdc := codec.New()
	genesisState := app.NewDefaultGenesisState(cdc)
	// use genesisState to initialize the system
}
```

`proposals_whitelisting.go` contains the `IsProposalWhitelisted` function, which checks if a proposal is whitelisted based on its content. This is useful for ensuring that only certain types of proposals are allowed to be submitted and voted on by the governance module.

Example usage:

```go
import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/duality-solutions/duality/app"
)

func main() {
	// create a new proposal
	proposal := types.NewParameterChangeProposal("title", "description", []types.ParamChange{
		{Subspace: "bank", Key: "SendEnabled", Value: "true"},
		{Subspace: "staking", Key: "MaxValidators", Value: "100"},
	})

	// check if the proposal is whitelisted
	if app.IsProposalWhitelisted(proposal) {
		// submit the proposal for voting
		// ...
	} else {
		// reject the proposal
		// ...
	}
}
```

The `params` subfolder contains the `proto.go` file, which creates an `EncodingConfig` for a non-amino based test configuration. This is important for testing the project's functionality without relying on amino.

Example usage:

```go
import (
    "github.com/duality/params"
)

func main() {
    encodingConfig := params.MakeTestEncodingConfig()
    // use encodingConfig to test project functionality
}
```
