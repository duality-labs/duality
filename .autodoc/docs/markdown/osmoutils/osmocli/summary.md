[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/osmoutils/osmocli)

The `osmocli` package provides a command-line interface (CLI) for interacting with the Osmocom cellular network stack and a Cosmos SDK-based blockchain. It offers a flexible and extensible way to handle command-line flags, create query and transaction commands, and generate formatted long descriptions for CLI commands.

The `flag_advice.go` file defines types and functions for parsing and handling command-line flags. It allows users to customize the behavior of the CLI by defining custom flag names and parsers. For example:

```go
flagAdvice := osmocli.FlagAdvice{
    HasPagination: true,
    CustomFlagOverrides: map[string]string{
        "flag1": "custom_flag1",
    },
    CustomFieldParsers: map[string]osmocli.CustomFieldParserFn{
        "field1": osmocli.FlagOnlyParser(func() interface{} { return new(int) }),
    },
}
```

The `index_cmd.go` file defines a CLI command that can be used to query information about a specific module in an Osmocom cellular network. Here's an example of how to create a root command and add a module command:

```go
rootCmd := &cobra.Command{Use: "duality"}
moduleCmd := osmocli.IndexCmd("module_name")
rootCmd.AddCommand(moduleCmd)
```

The `parsers.go` file provides functionality for parsing command-line arguments and flags for the duality project. It uses reflection to dynamically parse the fields of a given struct based on the provided arguments and flags:

```go
type MyStruct struct {
    Name string
    Age  int
}

flagAdvice := osmocli.FlagAdvice{}
flags := pflag.NewFlagSet("my-command", pflag.ContinueOnError)
args := []string{"John", "25"}

parsedStruct, err := osmocli.ParseFieldsFromFlagsAndArgs[MyStruct](flagAdvice, flags, args)
```

The `query_cmd_wrap.go` file provides a way to create query commands for a Cosmos SDK-based blockchain. It allows developers to define the properties of a query command and the function to call on the blockchain to retrieve the data:

```go
queryDesc := osmocli.QueryDescriptor{
    Name: "myquery",
    Desc: "My custom query command",
    QueryFn: func(clientCtx client.Context, req *myquery.Request) (*myquery.Response, error) {
        // Call the blockchain to retrieve the data
    },
}
queryCmd := osmocli.BuildQueryCli(queryDesc, createGrpcClient)
```

The `string_formatter.go` file generates long descriptions for CLI commands. Developers can use the `FormatLongDescDirect` function to generate a formatted long description string for a specific module:

```go
longDesc := "This command does something.\n\nUsage: {{.CommandPrefix}} command [args]\n\n{{.ExampleHeader}}\n{{.CommandPrefix}} command arg1 arg2"
moduleName := "mymodule"
formattedLongDesc := osmocli.FormatLongDescDirect(longDesc, moduleName)
```

The `tx_cmd_wrap.go` file provides a framework for building CLI commands for interacting with a Cosmos SDK-based blockchain. It allows developers to easily define the properties of a transaction command and build a Cobra command from those properties:

```go
desc := &osmocli.TxCliDesc{
    Use: "mytx",
    Short: "My custom transaction command",
    Long: "This command does something cool",
    NumArgs: 2,
    ParseAndBuildMsg: func(clientCtx client.Context, args []string, flags *pflag.FlagSet) (sdk.Msg, error) {
        // Parse arguments and build a message
    },
    TxSignerFieldName: "from",
    Flags: osmocli.FlagDesc{
        {"flag1", "", "Flag 1", true},
        {"flag2", "", "Flag 2", true},
    },
}
cmd := desc.BuildCommandCustomFn()
```
