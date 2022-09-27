package make

import (
	"github.com/spf13/cobra"
	"goapi/pkg/console"
	"os"
	"path/filepath"
	"strings"
)

var CmdMakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，exmaple: make controller v1/user",
	Run:   runMakeController,
	Args:  cobra.ExactArgs(1),
}

func runMakeController(cmd *cobra.Command, args []string) {

	array := strings.Split(args[0], "/")
	for k, v := range array {
		array[k] = strings.ReplaceAll(v, "/", "")
	}
	var apiVersion, dirName, name, dir string
	version := make(map[string]string)
	switch len(array) {
	case 2:
		apiVersion, name = array[0], array[1]
		dir = filepath.Join("app/http/controllers/api/", apiVersion)
		os.MkdirAll(dir, os.ModePerm)
	case 3:
		apiVersion, dirName, name = array[0], array[1], array[2]
		// version["{{PackageName}}"] = dirName
		dir = filepath.Join("app/http/controllers/api/", apiVersion, dirName)
		os.MkdirAll(dir, os.ModePerm)
	default:
		console.Exit("api controller name format: [v1/user|v1/auth/login]")
	}
	model := makeModelFromString(name)
	version["{{Version}}"] = apiVersion
	createFileFromStub(filepath.Join(dir, model.TableName+"_controller.go"), "controller", model, version)

}
