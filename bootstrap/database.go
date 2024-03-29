package bootstrap

import (
	"errors"
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/database"
	"goapi/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.user"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		sqlitedb := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(sqlitedb)
	default:
		panic(errors.New("database connection not supported"))
	}
	database.Connect(dbConfig, logger.NewGormLogger())
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	// 不自动执行，使用migrate命令执行迁移
	// database.DB.AutoMigrate(&user.User{})
}
