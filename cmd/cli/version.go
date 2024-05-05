package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Apotheca",
	Long: `This command will print the version of Apotheca
	currently installed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Apotheca v0.02")
	},
}
