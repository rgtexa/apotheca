package cli

import (
	"github.com/rgtexa/apotheca/cmd/server"
	"github.com/spf13/cobra"
)

func init() {
	var debug bool

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the Apotheca server",
		Long:  `This command will start and run the Apotheca web server.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			server.RunServer(args)
		},
	}

	runCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	rootCmd.AddCommand(runCmd)
}
