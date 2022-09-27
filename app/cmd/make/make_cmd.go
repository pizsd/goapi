package make

import (
	"github.com/spf13/cobra"
	"goapi/pkg/console"
	"path/filepath"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filePath := filepath.Join("app/cmd/", model.PackageName+".go")
	createFileFromStub(filePath, "cmd", model)
	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
