package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrateFunc func(gorm.Migrator, *sql.DB)

var migrationFiles []migrationFile

type migrationFile struct {
	Up       migrateFunc
	Down     migrateFunc
	FileName string
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up, down migrateFunc) {
	migrationFiles = append(migrationFiles, migrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
