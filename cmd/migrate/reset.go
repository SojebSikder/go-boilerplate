package migrate

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database by rolling back all migrations and reapplying them",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}
		defer db.Close()

		if err := goose.SetDialect("postgres"); err != nil {
			return err
		}
		return goose.Reset(db, dir)
	},
}
