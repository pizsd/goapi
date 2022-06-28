package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/pkg/response"
	"github.com/thedevsaddam/govalidator"
)

type ValidateFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。")
		fmt.Println(err.Error())
		return false
	}
	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
