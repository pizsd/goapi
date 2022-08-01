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

type LoginByMultiRequest struct {
	CaptchaId     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	LoginId  string `json:"login_id,omitempty" valid:"login_id"`
	Password string `json:"password,omitempty" valid:"password"`
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

func LoginByMulti(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"login_id":       []string{"required"},
		"password":       []string{"required"},
	}
	messages := govalidator.MapData{
		"captcha_id":     []string{"required:图片验证码ID不能为空"},
		"captcha_answer": []string{"required:图片验证码不能为空", "digits:图片验证码格式不正确"},
		"login_id":       []string{"required:登录ID不能为空"},
		"password":       []string{"required:密码不能为空"},
	}
	errs := validate(data, rule, messages)
	_data := data.(*LoginByMultiRequest)
	errs = validators.VerifyCaptcha(_data.CaptchaId, _data.CaptchaAnswer, errs)

	return errs
}
