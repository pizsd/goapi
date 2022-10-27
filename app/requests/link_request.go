package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LinkRequest struct {
	Name string `valid:"name" json:"name"`
	Url  string `valid:"url" json:"url,omitempty"`
	Logo string `valid:"logo" json:"logo,omitempty"`
}

func LinkSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:50", "not_exists:links,name"},
		"url":  []string{"required", "url"},
		"logo": []string{"url"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 50 个字",
			"not_exists:名称已存在",
		},
		"url": []string{
			"required:链接地址为必填项",
			"url:无效的URL",
		},
		"logo": []string{
			"url:无效的URL",
		},
	}
	return validate(data, rules, messages)
}
