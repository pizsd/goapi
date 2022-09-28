package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/app/requests/validators"
)

type PasswordResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

type PasswordResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetPasswordByPhone(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:8"},
	}
	messages := govalidator.MapData{
		"phone":       []string{"required:手机号不能为空", "digits:手机号格式不正确"},
		"verify_code": []string{"required:验证码不能为空", "digits:验证码格式不正确"},
		"password":    []string{"required:密码不能为空", "min:密码长度应大于8"},
	}
	errs := validate(data, rule, messages)
	_data := data.(*PasswordResetByPhoneRequest)
	errs = validators.VerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

func ResetPasswordByEmail(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"email":       []string{"required", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:8"},
	}
	messages := govalidator.MapData{
		"phone":       []string{"required:Email不能为空", "digits:Email格式不正确"},
		"verify_code": []string{"required:验证码不能为空", "digits:验证码格式不正确"},
		"password":    []string{"required:密码不能为空", "min:密码长度应大于8"},
	}
	errs := validate(data, rule, messages)
	_data := data.(*PasswordResetByEmailRequest)
	errs = validators.VerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}
