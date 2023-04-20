[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/cmd)

The `.autodoc/docs/json/cmd` folder contains code for initializing and executing the duality blockchain node, as well as modifying the genesis state of the blockchain network for testing purposes. The code is organized into two main files: `consumer.go` and `main.go`.

`consumer.go` defines a CLI command for adding a consumer section to the genesis state of the blockchain network. This consumer section is used for testing the interchain security module, which verifies the validity of transactions between different blockchain networks. The `AddConsumerSectionCmd` function returns a Cobra command that can be executed from the command line. When executed, it reads the current genesis state, modifies the consumer section using a callback function, and writes the updated genesis state back to disk. The code utilizes external packages such as `cosmos-sdk`, `tendermint`, and `interchain-security`, as well as a custom package called `duality` for testing purposes.

```go
// Example usage of AddConsumerSectionCmd
rootCmd.AddCommand(AddConsumerSectionCmd(defaultNodeHome))
```

`main.go` initializes and executes the duality blockchain node using the Cosmos SDK-based framework. The `main` function initializes the root command for the duality blockchain node using the `cosmoscmd.NewRootCmd` function, which takes several arguments, including the name of the application, the account address prefix, the default node home directory, the name of the application again, the module basics, and a function that creates a new instance of the application. The `rootCmd` variable is then used to add the `AddConsumerSectionCmd` to the root command. Finally, the `svrcmd.Execute` function is called with the root command and the default node home directory as arguments, which starts the duality blockchain node.

```go
// Example usage of main.go
func main() {
    rootCmd := cosmoscmd.NewRootCmd(appName, prefix, defaultNodeHome, appCreator, app.ModuleBasics())
    rootCmd.AddCommand(AddConsumerSectionCmd(defaultNodeHome))
    svrcmd.Execute(rootCmd, defaultNodeHome)
}
```

Developers can use this code as a starting point for building custom blockchain applications on top of the duality blockchain. They can add new commands to the root command to provide additional functionality to the blockchain node. For example, a developer might create a new command for querying the state of a specific module or for submitting a new transaction to the network.
