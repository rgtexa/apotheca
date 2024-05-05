package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rgtexa/apotheca/cmd/server"
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
		fmt.Printf("DSN: %s\n", dsn)
		conn, err := pgx.Connect(context.Background(), dsn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Connected successfully!")
		defer conn.Close(context.Background())
	},
}
