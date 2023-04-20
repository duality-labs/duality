[View code on GitHub](https://github.com/duality-labs/duality/osmoutils/osmocli/string_formatter.go)

The `osmocli` package contains code that is used to generate long descriptions for CLI commands in the duality project. The `LongMetadata` struct defines the metadata that is used to generate the long description. It contains the binary name, command prefix, short description, and an example header. The `NewLongMetadata` function creates a new instance of the `LongMetadata` struct and sets the command prefix based on the module name. 

The `FormatLongDesc` function takes a long string and a `LongMetadata` struct as input and returns a formatted long description string. It uses the `text/template` package to parse the long string and replace any placeholders with the values from the `LongMetadata` struct. If the parsing fails, it panics with an error message. The formatted long description string is returned after trimming any leading or trailing white space.

The `FormatLongDescDirect` function is a convenience function that takes a long string and a module name as input and returns a formatted long description string. It calls the `FormatLongDesc` function with a new instance of the `LongMetadata` struct created using the module name.

This code is used to generate long descriptions for CLI commands in the duality project. Developers can use the `FormatLongDescDirect` function to generate a formatted long description string for a specific module and use it in the CLI command documentation. For example:

```
longDesc := "This command does something.\n\nUsage: {{.CommandPrefix}} command [args]\n\n{{.ExampleHeader}}\n{{.CommandPrefix}} command arg1 arg2"
moduleName := "mymodule"
formattedLongDesc := FormatLongDescDirect(longDesc, moduleName)
```

The `formattedLongDesc` variable will contain the following string:

```
This command does something.

Usage: $ duality q mymodule command [args]

Example:
$ duality q mymodule command arg1 arg2
```
## Questions: 
 1. What is the purpose of the `LongMetadata` struct?
   
   The `LongMetadata` struct is used to store information about a command's long description, including the binary name, command prefix, and an example header.

2. What is the purpose of the `FormatLongDesc` function?
   
   The `FormatLongDesc` function takes a long description string and a `LongMetadata` struct, and returns a formatted string that includes the long description and example header.

3. What is the difference between `NewLongMetadata` and `FormatLongDescDirect`?
   
   `NewLongMetadata` creates a new `LongMetadata` struct with a given module name, while `FormatLongDescDirect` takes a long description string and a module name, and returns a formatted string that includes the long description and example header.