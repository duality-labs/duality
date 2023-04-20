[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/x/epochs/keeper)

The `keeper` package in the `duality` project is responsible for managing the state of epochs, which are periods of time during which certain actions can be taken in the system. The package contains several files that implement various functionalities related to epoch management.

`abci.go` contains the `BeginBlocker` function, which is responsible for determining when a new epoch should begin and performing the necessary actions to start it. This function is critical for managing the timing of epochs, which is essential for the proper functioning of the system. Other parts of the system can use the epoch information stored in the system to determine when certain actions can be taken, such as executing a smart contract.

`epoch.go` provides the implementation of the `Keeper` struct, which manages the state of `EpochInfo` objects. The `Keeper` struct offers methods for adding, retrieving, and deleting `EpochInfo` objects, as well as iterating through all the epochs. This package is likely used in conjunction with other packages in the project to manage the overall state of the system.

`genesis.go` manages epoch information in the duality blockchain during the initialization and export of the blockchain's genesis state. The `InitGenesis` function sets the epoch information from the genesis state, while the `ExportGenesis` function exports the epoch information to the genesis state. This code allows for the addition and retrieval of epoch information, which can be used by other parts of the duality project to enforce rules or conditions during specific epochs.

`hooks.go` defines two functions, `AfterEpochEnd` and `BeforeEpochStart`, which are called at the end and start of an epoch, respectively. These functions are designed to be used as hooks in the `duality` project, allowing developers to define their own hook functions and register them with the `hooks` object to perform custom actions at the start or end of an epoch.

`keeper.go` contains the implementation of the `Keeper` struct, which is responsible for managing the state of the `epochs` module in the larger `duality` project. The `NewKeeper` function initializes a new `Keeper` instance, the `SetHooks` method sets the hooks for the `epochs` module, and the `Logger` method gets a logger instance for the `epochs` module.

Overall, the `keeper` package plays a crucial role in the `duality` project by managing the state of epochs and providing functionalities for adding, retrieving, and deleting epoch information. This package enables the implementation of various rules and conditions during specific epochs, ensuring the proper functioning of the system.
