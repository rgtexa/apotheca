package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rgtexa/apotheca/cmd/server"
	"github.com/spf13/cobra"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		cfg := &server.Configuration{}

		cfgReader, err := os.ReadFile("apotheca.json")
		if err != nil {
			fmt.Println("failed to read configuration file apotheca.json")
		}

		err = json.Unmarshal(cfgReader, cfg)
		if err != nil {
			fmt.Println("failed to parse configuration file apotheca.json")
			fmt.Println(err)
		}

		dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.Database.DBProvider, cfg.Database.DBUser, cfg.Database.DBPass, cfg.Database.DBHost, cfg.Database.DBPort, cfg.Database.DBName)
		m, err := migrate.New("file://migrations", dsn)
		if err != nil {
			fmt.Println(err)
		}
		err = m.Up()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Migration successful!")
	},
}
