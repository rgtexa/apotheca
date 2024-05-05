package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testdbCmd)
}

var testdbCmd = &cobra.Command{
	Use:   "testdb",
	Short: "Test connection to the database",
	Long: `This command will read the apotheca.json config,
	assemble the data source connection string, and
	attempt to connect to the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Testing database connection...")
	},
}
