package migrate

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all available migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure the DB exists
		if err := ensureDatabaseExists(dbURL); err != nil {
			return err
		}
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			return fmt.Errorf("failed to open database: %w", err)
		}
		defer db.Close()

		if err := goose.SetDialect("postgres"); err != nil {
			return err
		}
		return goose.Up(db, dir)
	},
}

func ensureDatabaseExists(databaseURL string) error {
	// Parse URL
	u, err := url.Parse(databaseURL)
	if err != nil {
		return fmt.Errorf("invalid DATABASE_URL: %w", err)
	}

	dbName := strings.TrimPrefix(u.Path, "/")
	if dbName == "" {
		return fmt.Errorf("no database name in URL")
	}

	// Connect to default postgres database
	u.Path = "/postgres"
	masterURL := u.String()

	db, err := sql.Open("postgres", masterURL)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer db.Close()

	// Try to create the target database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create database: %w", err)
	}

	return nil
}
