package migrate

import (
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <name> [sql|go]",
	Short: "Create a new migration file",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		migrationType := "sql"
		if len(args) > 1 {
			migrationType = args[1]
		}
		fmt.Printf("Creating migration '%s' of type '%s' in %s\n", name, migrationType, dir)
		return goose.Create(nil, dir, name, migrationType)
	},
}
