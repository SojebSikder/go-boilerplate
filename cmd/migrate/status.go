package migrate

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of all migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}
		defer db.Close()

		if err := goose.SetDialect("postgres"); err != nil {
			return err
		}
		return goose.Status(db, dir)
	},
}
