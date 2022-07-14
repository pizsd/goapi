package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/app/requests/validators"
)

type VerifyCodePhoneRequest struct {
	Phone         string `json:"phone,omitempty" valid:"phone"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

type VerifyCodeEmailRequest struct {
	Email         string `json:"email,omitempty" valid:"email"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号格式不正确",
		},
		"captcha_id": []string{
			"required:验证码ID不能为空",
		},
		"captcha_answer": []string{
			"required:验证码不能为空",
			"digits:验证码格式不正确",
		},
	}
	errs := validate(data, rules, messages)

	if len(errs) != 0 {
		return errs
	}
	_data := data.(*VerifyCodePhoneRequest)
	errs = validators.VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}

func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":          []string{"required", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email不能为空",
			"email:Email格式不正确",
		},
		"captcha_id": []string{
			"required:验证码ID不能为空",
		},
		"captcha_answer": []string{
			"required:验证码不能为空",
			"digits:验证码格式不正确",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*VerifyCodeEmailRequest)
	errs = validators.VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
