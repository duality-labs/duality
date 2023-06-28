package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/cmd/dualityd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	rootCmd.AddCommand(AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
