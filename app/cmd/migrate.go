package cmd

import (
	"github.com/spf13/cobra"
	_ "goapi/database/migrations" // 为了执行migrations下的所有迁移文件的init
	"goapi/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use: "rollback",
	// 设置别名 migrate down == migrate rollback
	Aliases: []string{"down"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
	)
}
func migrator() *migrate.Migrator {
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Down()
}
