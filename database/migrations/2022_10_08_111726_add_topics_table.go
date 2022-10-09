package migrations

import (
	"database/sql"
	"goapi/app/models"
	"goapi/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(100);not null"`
		Content    string `gorm:"type:text;not null"`
		UserID     string `gorm:"type:int(11);index;not null"`
		CategoryID string `gorm:"type:int(11);index;not null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_10_08_111726_add_topics_table", up, down)
}
