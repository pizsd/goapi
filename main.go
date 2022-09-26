package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"goapi/app/cmd"
	"goapi/bootstrap"
	btsConfig "goapi/config"
	"goapi/pkg/config"
	"goapi/pkg/console"
	"os"
)

func init() {
	// 可以使用匿名导入包，这样就可以不需要config/config.go了
	// 这里可能是显式的调用，可能是可读性更好
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "GoApi",
		Short: "A simple api project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			// 初始化 Logger
			bootstrap.SetupLogger()
			// 初始化数据库
			bootstrap.SetupDB()
			// 初始化 Redis
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
