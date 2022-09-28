package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrateFunc func(gorm.Migrator, *sql.DB)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrateFunc
	Down     migrateFunc
	FileName string
}

func (mf *MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, mg := range migrations {
		if mf.FileName == mg.Migration {
			return false
		}
	}
	return true
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up, down migrateFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
