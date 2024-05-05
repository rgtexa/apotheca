package cli

import (
	"github.com/rgtexa/apotheca/cmd/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Apotheca server",
	Long:  `This command will start and run the Apotheca web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}
