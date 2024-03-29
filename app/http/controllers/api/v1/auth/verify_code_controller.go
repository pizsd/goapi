package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/http/controllers/api/v1"
	"goapi/app/requests"
	"goapi/pkg/captcha"
	"goapi/pkg/logger"
	"goapi/pkg/response"
	"goapi/pkg/verifycode"
)

type VerifyCodeController struct {
	v1.BaseApiController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.Data(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) VerifySmsCode(c *gin.Context) {
}

func (vc *VerifyCodeController) SendSmsCode(c *gin.Context) {
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	if ok := verifycode.NewVerifyCode().SendSmsCode(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}

func (vc *VerifyCodeController) SendEmailCode(c *gin.Context) {
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}
	if ok := verifycode.NewVerifyCode().SendEmailCode(request.Email); !ok {
		response.Abort500(c, "发送邮件失败~")
	} else {
		response.Success(c)
	}
}
