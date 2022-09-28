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

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}
func migrator() *migrate.Migrator {
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
