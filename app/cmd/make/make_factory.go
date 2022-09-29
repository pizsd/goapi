package make

import (
	"github.com/spf13/cobra"
	"path/filepath"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, example: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1),
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	path := filepath.Join("database/factories/", args[0]+"_factory.go")
	createFileFromStub(path, "factory", model)
}
