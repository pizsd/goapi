package migrations

import (
	"database/sql"
	"goapi/app/models"
	"goapi/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(20);not null"`
		Introduction string `gorm:"type:varchar(255);default:''"`
		Avatar       string `gorm:"type:varchar(255);default:''"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Introduction")
		migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2022_11_13_211128_add_fields_to_user", up, down)
}
