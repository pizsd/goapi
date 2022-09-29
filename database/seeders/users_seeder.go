package seeders

import (
	"fmt"
	"goapi/database/factories"
	"goapi/pkg/console"
	"goapi/pkg/logger"
	"goapi/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建 10 个用户对象
		users := factories.MakeUsers(10)
		res := db.Table("users").Create(&users)
		// 记录错误
		if err := res.Error; err != nil {
			logger.LogIf(err)
			return
		}
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", res.Statement.Table, res.RowsAffected))
	})
}
