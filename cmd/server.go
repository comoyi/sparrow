package cmd

import (
	"github.com/comoyi/sparrow/cmd/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
