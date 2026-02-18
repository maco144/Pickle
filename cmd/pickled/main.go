package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/maco144/pickle/app"
)

func main() {
	rootCmd, _ := svrcmd.CreateRootCommand(
		app.Name,
		"Pickle - Data Preservation Engine",
		app.NewApp,
		svrcmd.DefaultServerConfigurator(),
	)

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		server.PrintErr(err)
		os.Exit(1)
	}
}
