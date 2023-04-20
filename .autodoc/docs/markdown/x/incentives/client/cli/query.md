[View code on GitHub](https://github.com/duality-labs/duality/incentives/client/cli/query.go)

The `cli` package contains functions that define the command-line interface (CLI) for the Duality project's incentives module. The `GetQueryCmd` function returns a `cobra.Command` object that groups all the query commands for the incentives module under a subcommand. The `osmocli` package is used to create the CLI commands. 

The `GetCmdGetModuleStatus` function returns a `QueryDescriptor` object and a `GetModuleStatusRequest` object. The former defines the CLI command for querying the status of the incentives module, while the latter is used to specify any parameters required for the query. 

Similarly, the `GetCmdGetGaugeByID`, `GetCmdGauges`, `GetCmdGetStakeByID`, and `GetCmdStakes` functions define CLI commands for querying gauges and stakes by ID or status. These functions return a `QueryDescriptor` object and a corresponding `GetGaugeByIDRequest`, `GetGaugesRequest`, `GetStakeByIDRequest`, or `GetStakesRequest` object, respectively. 

The `GetCmdGetFutureRewardEstimate` function returns a `QueryDescriptor` object and a `GetFutureRewardEstimateRequest` object. This command is used to estimate future rewards for a given set of stakes. The `CustomFieldParsers` field in the `QueryDescriptor` object is used to specify custom parsing functions for the command's parameters. 

Overall, this package defines the CLI commands for querying the incentives module in the Duality project. These commands can be used by users to retrieve information about gauges and stakes, as well as estimate future rewards. 

Example usage:
```
$ duality query incentives module-status
$ duality query incentives gauge-by-id 1
$ duality query incentives list-gauges UPCOMING DualityPoolShares-stake-token-t0-f1
$ duality query incentives stake-by-id 1
$ duality query incentives list-stakes cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx
$ duality query incentives reward-estimate cosmos1chl62vc593p99z2tfh2pp8tl4anm0w4l8h8svx [1,2,3] 1681450672
```
## Questions: 
 1. What is the purpose of the `duality-labs/duality/osmoutils/osmocli` package?
- The `duality-labs/duality/osmoutils/osmocli` package is used to add query commands to the `cobra.Command` object.
2. What is the `GetQueryCmd` function used for?
- The `GetQueryCmd` function returns a `cobra.Command` object that contains several query commands related to incentives.
3. What is the purpose of the `parseGaugeStatus` function?
- The `parseGaugeStatus` function is a custom field parser that is used to parse the `Status` field in the `GetCmdGauges` query command. It converts the string value to the corresponding `types.GaugeStatus` enum value.