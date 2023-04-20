[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/incentives)

The `incentives` module in the Duality project is responsible for managing the incentives system for the Duality blockchain. It provides functionalities for creating, modifying, and retrieving gauges and stakes, as well as distributing rewards to users based on certain conditions. The module is organized into three main subfolders: `client`, `keeper`, and `types`.

The `client` subfolder contains the `cli` package, which provides a command-line interface (CLI) for interacting with the incentives module. Users can create gauges, add to gauges, stake tokens, and unstake tokens using the CLI. For example, users can query gauges and stakes, as well as estimate future rewards:

```sh
$ duality query incentives module-status
$ duality query incentives gauge-by-id 1
$ duality query incentives list-gauges UPCOMING DualityPoolShares-stake-token-t0-f1
$ duality query incentives stake-by-id 1
$ duality query incentives list-stakes cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx
$ duality query incentives reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672
```

The `keeper` subfolder manages the state and provides functions for creating, modifying, and retrieving gauges and stakes. Gauges are used to distribute rewards to users based on certain conditions, while stakes represent the tokens locked by users to participate in the network. Example usage of the `keeper` package might involve creating a new gauge, staking tokens, and retrieving the active gauges to distribute rewards to users:

```go
// create a new gauge
keeper.CreateGauge(ctx, ...)

// stake tokens
keeper.Stake(ctx, ...)

// get active gauges
activeGauges := keeper.GetActiveGauges(ctx)
```

The `types` package contains various types, functions, and interfaces that are used throughout the project, particularly for managing incentives, staking, and rewards distribution. The package provides functionality for registering concrete types and interfaces used for Amino JSON serialization and message services, as well as defining sentinel errors, events, and expected keepers for the `x/incentives` module.

In the larger project, the `incentives` module can be used to incentivize users to participate in the Duality blockchain by rewarding them with tokens for staking their coins. The `BeginBlocker` and `EndBlocker` functions in the `abci.go` file would be used to manage the incentives system by automatically unstaking matured stakes and distributing rewards to users. The module can be used to create and manage gauges, which store the yield to be distributed to stakers. The module can also be used to query gauge information and upcoming gauges. The functionalities provided by the module can be accessed through the CLI or REST service handlers.
