package make

import (
	"github.com/spf13/cobra"
	"path/filepath"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	path := filepath.Join("app/requests/", model.PackageName+"_request.go")

	createFileFromStub(path, "request", model)
}
