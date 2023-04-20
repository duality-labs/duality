[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/epochs/keeper)

The `keeper` package in the `duality` project is responsible for managing the state of the `epochs` module, which deals with epochs - periods of time during which certain actions can be taken in the system. The package contains several files that define functions and structs for managing epoch information, such as adding, retrieving, and deleting `EpochInfo` objects, as well as iterating through all the epochs.

The `abci.go` file contains the `BeginBlocker` function, which is responsible for managing the start and end of epochs. It checks whether a new epoch should begin and performs the necessary actions to start it. This function is critical for the proper functioning of the system, as it manages the timing of epochs.

The `epoch.go` file defines the `Keeper` struct and its methods for managing the state of `EpochInfo` objects. It provides methods for adding, retrieving, and deleting `EpochInfo` objects, as well as iterating through all the epochs and getting the number of blocks since the epoch started.

The `genesis.go` file contains functions for initializing and exporting the blockchain's genesis state, which includes epoch information. The `InitGenesis` function sets the epoch information from the genesis state, while the `ExportGenesis` function exports the current epoch information to the genesis state.

The `grpc_query.go` file defines a gRPC query server for the `x/epochs` module, allowing external clients to query epoch information in a standardized way. The `Querier` struct wraps around the `Keeper` struct and provides gRPC method handlers for retrieving running epoch information and the current epoch of a specified identifier.

The `hooks.go` file defines two functions, `AfterEpochEnd` and `BeforeEpochStart`, which are called at the end and start of an epoch, respectively. Developers can define their own hook functions and register them with the `hooks` object to perform custom actions at the start or end of an epoch.

The `keeper.go` file provides the implementation of the `Keeper` struct, which is responsible for managing the state of the `epochs` module. It contains methods for initializing a new `Keeper` instance, setting hooks for the `epochs` module, and getting a logger instance for the `epochs` module.

Here's an example of how to use the `Keeper` struct to add a new `EpochInfo` object:

```go
k := keeper.NewKeeper(storeKey)
epochInfo := types.EpochInfo{
    Identifier: "epoch-1",
    StartTime:  time.Now(),
}
err := k.AddEpochInfo(ctx, epochInfo)
if err != nil {
    log.Fatalf("Failed to add epoch info: %v", err)
}
```

Overall, the `keeper` package plays a crucial role in the `duality` project by managing the state of the `epochs` module, which is essential for the proper functioning of the system. Other parts of the system can use the epoch information stored in the system to determine when certain actions can be taken, such as executing a smart contract.
