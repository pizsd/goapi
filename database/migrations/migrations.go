package migrations

// 防止在migrations目录下没有.go文件导致第一次执行 make migration 编译报错
// 因为 是在执行migrate up 时，必须执行生成的迁移文件的init方法
// 已便将MigrateFile 添加到 migrationFiles 中去 ，所以在migrate.go中 使用了 _ "goapi/database/migrations"
