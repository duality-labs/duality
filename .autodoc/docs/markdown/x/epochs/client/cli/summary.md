[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/epochs/client/cli)

The `query.go` file in the `epochs/client/cli` folder provides a set of Command Line Interface (CLI) query commands for the epochs module in the duality project. These commands allow developers and users to retrieve information about the current epoch and running epoch information, which can be useful for debugging issues related to epoch transitions or monitoring the progress of the current epoch.

The main function in this file is `GetQueryCmd`, which returns a `cobra.Command` object that can be used to execute the CLI commands. This function first calls `osmocli.QueryIndexCmd` to create a new `cobra.Command` object with the name of the epochs module. It then adds two query commands to the command object using `osmocli.AddQueryCmd`. The first command queries running epoch information using the `EpochInfos` function from the `types` package. The second command queries the current epoch by specified identifier using the `QueryCurrentEpochRequest` function from the `types` package.

In addition to the `GetQueryCmd` function, there are two other functions in this file that return `osmocli.QueryDescriptor` objects that describe the CLI commands for querying epoch information. The `GetCmdEpochInfos` function returns a descriptor for querying running epoch information, while the `GetCmdCurrentEpoch` function returns a descriptor for querying the current epoch by specified identifier.

Here's an example of how these CLI commands might be used:

```sh
# Query running epoch information
dualitycli query epochs epoch-infos

# Query the current epoch by specified identifier
dualitycli query epochs current-epoch --identifier "example-identifier"
```

In summary, the `query.go` file in the `epochs/client/cli` folder provides a set of CLI commands for querying epoch information in the duality project. These commands can be used by developers and users to retrieve information about the current epoch and running epoch information, which can be useful for debugging issues related to epoch transitions or monitoring the progress of the current epoch.
