package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/e-ticket-terminal/internal/routers"
	"github.com/wildanfaz/e-ticket-terminal/migrations"
)

func InitCmd(ctx context.Context) {
	var rootCmd = &cobra.Command{
		Short: "E-Ticket Terminal App",
	}

	rootCmd.AddCommand(startEchoApp)
	rootCmd.AddCommand(migrateTables)
	rootCmd.AddCommand(rollbackTables)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var startEchoApp = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		routers.InitEchoRouter()
	},
}

var migrateTables = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Tables",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.MigrateTables(context.Background())
	},
}

var rollbackTables = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback Tables",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.RollbackTables(context.Background())
	},
}
