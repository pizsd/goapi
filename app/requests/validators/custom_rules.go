package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goapi/pkg/database"
	"goapi/pkg/logger"
	"strconv"
	"strings"
	"unicode/utf8"
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

	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))

		valLen := utf8.RuneCountInString(value.(string))

		if valLen < min {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("名称长度需至少 %d 个字", min)
		}
		return nil
	})

	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))

		valLen := utf8.RuneCountInString(value.(string))

		if valLen > max {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("名称长度不能超过 %d 个字", max)
		}
		return nil
	})

	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		s := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")
		tableName := s[0]
		tableField := s[1]
		reqValue := value.(string)
		var count int64
		database.DB.Table(tableName).Where(tableField+" = ?", reqValue).Count(&count)
		if count == 0 {
			if message != "" {
				return errors.New(message)
			} else {
				logger.DebugString("validator", "reqValue", "111")
				return fmt.Errorf("%v 不存在", reqValue)
			}
		}
		return nil
	})
}
