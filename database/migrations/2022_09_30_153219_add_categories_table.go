package migrations

import (
	"database/sql"
	"goapi/app/models"
	"goapi/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel
		Name        string `gorm:"type:varchar(20);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_09_30_153219_add_categories_table", up, down)
}
