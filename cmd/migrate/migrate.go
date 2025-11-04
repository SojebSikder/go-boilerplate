package migrate

import (
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/spf13/cobra"
)

var (
	dbURL string
	dir   string
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration management using goose",
}

func init() {
	ctg, _ := config.NewConfig()
	dbURL = ctg.Database.DatabaseURL

	MigrateCmd.PersistentFlags().StringVar(&dir, "dir", "migrations", "Directory containing migration files")

	// Register subcommands
	MigrateCmd.AddCommand(upCmd)
	MigrateCmd.AddCommand(downCmd)
	MigrateCmd.AddCommand(statusCmd)
	MigrateCmd.AddCommand(createCmd)
	MigrateCmd.AddCommand(resetCmd)
}
