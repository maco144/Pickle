package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	pickle "github.com/maco144/pickle"
)

func main() {
	rootCmd, _ := svrcmd.CreateRootCommand(
		pickle.Name,
		"Pickle - Data Preservation Engine",
		pickle.NewApp,
		svrcmd.DefaultServerConfigurator(),
	)

	if err := svrcmd.Execute(rootCmd, "", pickle.DefaultNodeHome); err != nil {
		server.PrintErr(err)
		os.Exit(1)
	}
}
