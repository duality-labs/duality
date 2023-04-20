[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/incentives/keeper)

The `keeper` package in the `incentives` module of the Duality project is responsible for managing the state and providing functions for creating, modifying, and retrieving gauges and stakes. Gauges are used to distribute rewards to users based on certain conditions, while stakes represent the tokens locked by users to participate in the network.

For example, the `gauge.go` file provides functions to create, modify, and retrieve gauges. The `CreateGauge` function creates a gauge and sends coins to it, while the `AddToGaugeRewards` function adds coins to an existing gauge. The `GetGauges` function returns upcoming, active, and finished gauges, which can be used to manage the distribution of rewards to users.

The `genesis.go` file initializes and exports the state of the incentives module, which is crucial for managing the incentives and rewards for users who participate in the network. The `InitializeAllStakes` and `InitializeAllGauges` functions are used to set the state of the stakes and gauges, respectively.

The `hooks.go` file defines hooks that are called at the start and end of each epoch in the Duality blockchain. These hooks are used to perform actions related to distributing rewards to users. The `AfterEpochEnd` hook retrieves and distributes rewards to users based on the active gauges.

The `invariants.go` file registers and executes invariants for the governance module, ensuring the integrity of the system. The `AccumulationStoreInvariant` and `StakesBalancesInvariant` functions are used to detect and prevent errors in the system, which could lead to incorrect calculations or other issues.

The `iterator.go` file provides functions to manage the state of the incentives module, including retrieving and manipulating data stored in the key-value store. The `getStakesFromIterator` function retrieves stakes from the key-value store and returns them as an array.

The `keeper.go` file defines the `Keeper` struct, which manages the storage of the incentives module. The `GetModuleBalance` and `GetModuleStakedCoins` functions return the full balance and staked balance of the module, respectively.

The `lock_refs.go` file provides functions to manage reference keys for staked assets. The `addStakeRefs` and `deleteStakeRefs` functions are used to add and delete reference keys for a given stake, which are used to track the staked assets and calculate the incentives that should be rewarded to the staker.

The `msg_server.go` file implements the message server interface for the incentives module, allowing for the creation and management of gauges and stakes. The `Stake` and `Unstake` functions are used to stake and unstake tokens, respectively.

The `params.go` file provides functions to get and set parameters in the incentive module, which can be useful for adjusting the incentives offered to users or changing the rules around how incentives are earned.

The `query_server.go` file implements the QueryServer struct, which provides gRPC method handlers for querying the status of the module, gauges, stakes, and future reward estimates.

The `stake.go` file provides methods for managing stakes, such as `CreateStake`, which creates a new stake object and stores it in the state.

The `utils.go` file contains functions for managing references to objects in the Duality project, such as `addRefByKey` and `deleteRefByKey`, which can be used to manage references to objects when they are created or deleted.

Example usage of the `keeper` package might involve creating a new gauge, staking tokens, and retrieving the active gauges to distribute rewards to users:

```go
// create a new gauge
keeper.CreateGauge(ctx, ...)

// stake tokens
keeper.Stake(ctx, ...)

// get active gauges
activeGauges := keeper.GetActiveGauges(ctx)
```
