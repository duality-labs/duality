[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x)

The `.autodoc/docs/json/x` folder contains essential modules for the Duality project, including `dex`, `epochs`, `incentives`, and `mev`. These modules provide core functionalities such as decentralized exchange, periodic event scheduling, incentives management, and Maximal Extractable Value (MEV) handling.

The `dex` package offers core functionality for the Duality decentralized exchange (DEX) module. It includes functions for initializing and exporting the DEX module's state, handling incoming messages related to trading, and implementing the AppModuleBasic and AppModule interfaces for the module. Developers can use this package to build more complex trading strategies and applications on top of the DEX module.

```go
import (
    "github.com/duality-labs/duality/x/dex/keeper"
    "github.com/duality-labs/duality/x/dex/types"
)

func main() {
    // create a new DEX keeper
    k := keeper.NewKeeper()

    // create a new handler for the DEX module
    handler := NewHandler(k)

    // create a new deposit message
    depositMsg := types.NewMsgDeposit(...)

    // process the deposit message using the handler
    result, err := handler(ctx, depositMsg)
    if err != nil {
        // handle error
    }

    // handle result
}
```

The `epochs` module allows other modules to run code periodically by providing a generalized epoch interface. This enables modules to schedule and execute code at specified intervals without having to implement their own scheduling logic.

```go
import (
    "github.com/duality/epochs"
    "github.com/duality/epochs/types"
)

// Initialize the epochs module
appModule := epochs.NewAppModule(keeper)

// Register a new epoch
epochInfo := types.EpochInfo{
    Identifier: "weekly-update",
    StartTime:  time.Now(),
    Duration:   7 * 24 * time.Hour,
}
appModule.Keeper.AddEpochInfo(ctx, epochInfo)

// Register a hook to be executed at the start of the epoch
appModule.Keeper.SetHooks(types.MultiEpochHooks{
    BeforeEpochStart: func(ctx sdk.Context, epochInfo types.EpochInfo) {
        // Execute custom code at the start of the epoch
    },
})
```

The `incentives` module manages the incentives system for the Duality blockchain, providing functionalities for creating, modifying, and retrieving gauges and stakes, as well as distributing rewards to users based on certain conditions.

```go
// create a new gauge
keeper.CreateGauge(ctx, ...)

// stake tokens
keeper.Stake(ctx, ...)

// get active gauges
activeGauges := keeper.GetActiveGauges(ctx)
```

The `mev` package manages the Maximal Extractable Value (MEV) module, providing functionalities such as initializing and exporting the module's state, handling messages related to the module, and simulating the module's behavior.

```go
ctx := sdk.Context{}
keeper := keeper.Keeper{}
genesisState := types.GenesisState{}

InitGenesis(ctx, keeper, genesisState)
exportedGenesisState := ExportGenesis(ctx, keeper)
```

These modules play a crucial role in the Duality project, providing essential functionalities and enabling developers to build more complex applications on top of the Duality blockchain.
