package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up the Apotheca database tables",
	Long: `This command will connect to the database specified
	in the configuration file and run the database
	migrations to setup the Apotheca tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running database setup...")
	},
}
