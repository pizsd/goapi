package make

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	dir := filepath.Join("app/models/", model.PackageName)

	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll(dir, os.ModePerm)

	// 替换变量
	createFileFromStub(filepath.Join(dir, model.PackageName+"_model.go"), "model/model", model)
	createFileFromStub(filepath.Join(dir, model.PackageName+"_util.go"), "model/model_util", model)
	createFileFromStub(filepath.Join(dir, model.PackageName+"_hooks.go"), "model/model_hooks", model)
}
