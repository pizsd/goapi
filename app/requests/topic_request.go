package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `valid:"title" json:"title"`
	Content    string `valid:"content" json:"content"`
	CategoryID string `valid:"category_id" json:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:8", "max_cn:100"},
		"content":     []string{"required", "min_cn:15"},
		"category_id": []string{"required", "numeric", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:标题为必填项",
			"min_cn:名称长度需至少 8 个字",
			"max_cn:名称长度不能超过 100 个字",
		},
		"content": []string{
			"required:内容为必填项",
			"min_cn:内容长度需至少 15 个字",
		},
		"category_id": []string{
			"required:请选择分类",
			"exists:分类不存在",
		},
	}
	return validate(data, rules, messages)
}
