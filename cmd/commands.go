package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"multipliers/cmd/serve"
)

func ExecuteAppCmd() {

	appCmd := &cobra.Command{}
	appCmd.AddCommand(createServeCmd())

	if err := appCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func createServeCmd() *cobra.Command {

	serveCmd := &cobra.Command{
		Use: "serve",
		Run: serve.ExecuteCmdFn,
	}

	return serveCmd
}
