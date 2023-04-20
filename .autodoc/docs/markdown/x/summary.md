[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/x)

The `.autodoc/docs/json/x` folder contains the core modules for the Duality project, including `dex`, `epochs`, `incentives`, and `mev`. These modules are responsible for managing various aspects of the project, such as decentralized exchange functionality, periodic event handling, incentives distribution, and miner-extractable value tracking.

The `dex` module, located in the `dex` folder, manages the trading of assets, liquidity pools, and limit orders within the project. It initializes and exports the genesis state of the DEX module, handles various message types related to DEX functionality, and provides utility functions for error handling and mathematical operations.

Example usage of the `dex` module:

```go
// Initialize the DEX module's state with specific parameters
ctx := sdk.Context(...)
keeper := keeper.NewKeeper(...)
genesisState := types.GenesisState{...}
InitGenesis(ctx, keeper, genesisState)

// Create a new handler for processing messages related to the DEX module
handler := NewHandler(keeper)

// Send a message to the DEX module
msg := types.MsgDeposit{...}
res, err := handler(ctx, msg)

// Export the DEX module's state
exportedGenesisState := ExportGenesis(ctx, keeper)
```

The `epochs` module, located in the `epochs` folder, allows other modules to run code periodically by providing a generalized epoch interface. It manages the state of epochs and provides hooks for custom actions at the start or end of an epoch.

Example usage of the `epochs` module:

```go
// Initialize the epochs module
appModule := epochs.NewAppModule(keeper.NewKeeper(...))

// Register the epochs module with the application
app.RegisterModule(appModule)

// Define a custom hook function to be executed at the start of an epoch
func myEpochStartHook(ctx sdk.Context, epoch types.EpochInfo) {
    // Perform custom actions here
}

// Register the custom hook function with the epochs module
appModule.GetKeeper().SetHooks(types.NewMultiEpochHooks(myEpochStartHook))
```

The `incentives` module, located in the `incentives` folder, manages the incentives system for the Duality blockchain. It provides functionality for creating, modifying, and retrieving gauges, which are used to distribute rewards to users based on certain conditions.

Example usage of the `incentives` module:

```go
ctx := types.NewContext(nil, types.Header{}, false, nil)
req := types.RequestBeginBlock{}
k := keeper.NewKeeper()

incentives.BeginBlocker(ctx, req, k)
updates := incentives.EndBlocker(ctx, k)
// do something with updates
```

The `mev` module, located in the `mev` folder, manages and tracks various aspects of the project's functionality, such as handling incoming messages and executing appropriate actions. It also interacts with the state of the blockchain, including reading and writing data, managing parameters, and handling transactions and events.

Example usage of the `mev` module:

```go
// Initialize the mev module's state with specific parameters
ctx := sdk.Context(...)
keeper := keeper.NewKeeper(...)
genesisState := types.GenesisState{...}
InitGenesis(ctx, keeper, genesisState)

// Create a new handler for processing messages related to the mev module
handler := NewHandler(keeper)

// Send a message to the mev module
msg := types.MsgSend{...}
res, err := handler(ctx, msg)

// Export the mev module's state
exportedGenesisState := ExportGenesis(ctx, keeper)
```

In summary, the `.autodoc/docs/json/x` folder contains the core modules for the Duality project, which work together to provide a robust and flexible system for handling various aspects of the project's functionality.
