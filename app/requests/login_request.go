package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/app/requests/validators"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone":       []string{"required:手机号不能为空", "digits:手机号格式不正确"},
		"verify_code": []string{"required:验证码不能为空", "digits:验证码格式不正确"},
	}
	errs := validate(data, rule, messages)
	_data := data.(*LoginByPhoneRequest)
	errs = validators.VerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
