package make

import (
	"github.com/spf13/cobra"
	"goapi/pkg/app"
	"goapi/pkg/console"
	"path/filepath"
)

var CmdMakeMigrate = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	fileName := timeStr + "_" + model.PackageName
	path := filepath.Join("database/migrations", fileName+".go")
	createFileFromStub(path, "migration", model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")

}
