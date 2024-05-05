package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apotheca",
	Short: "Apotheca is a web-based document repository management system",
	Long: `A web-based document repository management system with
			role-based permissions and training mangement.
			Complete docs at apotheca.doc`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
