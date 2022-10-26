package make

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	dir := "app/policies/"
	os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(filepath.Join(dir, model.PackageName+"_policy.go"), "policy", model)
}
