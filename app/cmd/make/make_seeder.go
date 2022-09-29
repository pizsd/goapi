package make

import (
	"github.com/spf13/cobra"
	"path/filepath"
)

var CmdMakeSeed = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example:  make seeder user",
	Run:   runMakeSeed,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeed(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	path := filepath.Join("database/seeders/", args[0]+"_seeder.go")
	createFileFromStub(path, "seeder", model)
}
