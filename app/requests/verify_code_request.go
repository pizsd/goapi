package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/pkg/captcha"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	Phone         string `json:"phone,omitempty" valid:"phone"`
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
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = []string{"图片验证码错误"}
	}
	return errs

}
