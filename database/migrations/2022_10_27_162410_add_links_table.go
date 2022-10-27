package migrations

import (
	"database/sql"
	"goapi/app/models"
	"goapi/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(50);not null"`
		Url  string `gorm:"type:varchar(255);not null"`
		Logo string `gorm:"type:varchar(255);not null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_10_27_162410_add_links_table", up, down)
}
