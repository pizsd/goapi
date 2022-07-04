package validators

import "github.com/pizsd/goapi/pkg/captcha"

func VerifyCaptcha(captchaId, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaId, captchaAnswer); !ok {
		errs["captcha_answer"] = []string{"图片验证码错误"}
	}
	return errs
}
