package routes

import (
	"github.com/gin-gonic/gin"
	"goapi/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "PONG",
		})
	})
	v1 := r.Group("/v1")
	{

		authGroup := v1.Group("/auth")
		{
			sc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", sc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", sc.IsEmailExist)
			authGroup.POST("/signup/using-phone", sc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", sc.SignupUsingEmail)
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-code/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-code/phone", vcc.SendSmsCode)
			authGroup.POST("verify-code/email", vcc.SendEmailCode)
			lc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", lc.LoginByPhone)
			authGroup.POST("/login/using-multi", lc.LoginByMulti)
		}
	}
}
