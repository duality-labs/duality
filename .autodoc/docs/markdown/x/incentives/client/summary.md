[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/incentives/client)

The `.autodoc/docs/json/x/incentives/client` folder contains the `cli` package, which provides a command-line interface (CLI) for interacting with the incentives module of the Duality project. This package allows users to create gauges, add to gauges, stake tokens, and unstake tokens using the CLI.

The `cli` package consists of three main files: `flags.go`, `query.go`, and `tx.go`.

`flags.go` is responsible for defining and creating flag sets for the incentives module transaction commands. It imports the `flag` package from `github.com/spf13/pflag` to create these flag sets. The file defines three constants, `FlagStartTime`, `FlagPerpetual`, and `FlagAmount`, which are used as keys to access the corresponding flag values. The `FlagSetCreateGauge()` function returns a flag set for creating gauges, and the `FlagSetUnSetupStake()` function returns a flag set for unstaking an amount.

`query.go` defines the CLI commands for querying the incentives module in the Duality project. Users can retrieve information about gauges and stakes, as well as estimate future rewards. The `GetQueryCmd` function returns a `cobra.Command` object that groups all the query commands for the incentives module under a subcommand. Functions like `GetCmdGetModuleStatus`, `GetCmdGetGaugeByID`, `GetCmdGauges`, `GetCmdGetStakeByID`, and `GetCmdStakes` define CLI commands for querying gauges and stakes by ID or status.

Example usage of query commands:

```
$ duality query incentives module-status
$ duality query incentives gauge-by-id 1
$ duality query incentives list-gauges UPCOMING DualityPoolShares-stake-token-t0-f1
$ duality query incentives stake-by-id 1
$ duality query incentives list-stakes cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx
$ duality query incentives reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672
```

`tx.go` contains the transaction commands for the incentives module. The `GetTxCmd` function returns a `cobra.Command` that includes all the transaction commands for the incentives module. Functions like `NewCreateGaugeCmd`, `NewAddToGaugeCmd`, `NewStakeCmd`, and `NewUnstakeCmd` return a `osmocli.TxCliDesc` and a corresponding message type (`types.MsgCreateGauge`, `types.MsgAddToGauge`, `types.MsgStake`, or `types.MsgUnstake`). These messages are sent to the blockchain to create a new gauge, add tokens to an existing gauge, stake tokens into the stakeup pool, or unstake tokens from the stakeup pool.

In summary, the `cli` package provides a user-friendly way for users to interact with the incentives module of the Duality project. Users can create gauges, add to gauges, stake tokens, and unstake tokens using the CLI. This package plays a crucial role in enabling users to manage and interact with the incentives module, making it an essential part of the Duality project.
