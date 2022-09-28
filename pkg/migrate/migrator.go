package migrate

import (
	"goapi/pkg/console"
	"goapi/pkg/database"
	"goapi/pkg/file"
	"gorm.io/gorm"
	"os"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

func (migrator *Migrator) Up() {
	// 读取所有迁移文件，确保按照时间排序
	migrationFiles := migrator.readAllMigrationFiles()

	// 获取当前批次的值
	batch := migrator.getBatch()

	migrations := []Migration{}

	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	runed := false

	for _, mfile := range migrationFiles {
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}
	if !runed {
		console.Success("database is up to date.")
	}
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)
	var migrateFiles []MigrationFile
	for _, f := range files {
		fileName := file.FileNameWithoutExtension(f.Name())
		mfile := getMigrationFile(fileName)
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrateFiles, mfile)
		}
	}
	return migrationFiles
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) getBatch() int {
	batch := 1
	var lastMigration = Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	if lastMigration.Batch > 0 {
		return lastMigration.Batch
	}
	return batch
}

func (migrator *Migrator) runUpMigration(m MigrationFile, batch int) {
	// 执行Up迁移
	if m.Up != nil {
		console.Warning("migrating " + m.FileName)
		m.Up(database.DB.Migrator(), database.SQLDB)
		console.Success("migrated " + m.FileName)
	}
	// 执行完保存执行记录到migrations
	err := migrator.DB.Create(&Migration{Migration: m.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}
