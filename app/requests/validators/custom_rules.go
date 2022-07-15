package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goapi/pkg/database"
	"strings"
)

func init() {
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		ruleSlice := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
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
				return fmt.Errorf("%v 已被占用", requestValue)
			}
		}
		return nil
	})
}
