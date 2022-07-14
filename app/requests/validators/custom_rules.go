package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goapi/pkg/database"
	"strings"
)

func init() {
	govalidator.AddCustomRule("no_exist", func(field string, rule string, message string, value interface{}) error {
		ruleSlice := strings.Split(strings.TrimPrefix(rule, "no_exist:"), ",")
		tableName := ruleSlice[0]
		tableField := ruleSlice[1]
		var exceptId string
		if len(ruleSlice) > 2 {
			exceptId = ruleSlice[2]
		}
		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(tableField+"= ?", requestValue)
		if exceptId != "" {
			query.Where("id <> ?", exceptId)
		}
		var count int64
		query.Count(&count)
		if count > 0 {
			if message != "" {
				return errors.New(message)
			} else {
				return fmt.Errorf("%v 用户名已存在", requestValue)
			}
		}
		return nil
	})
}
