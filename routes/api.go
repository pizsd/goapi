package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "server is ok",
		})
	})
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "PONG",
			})
		})
		authGroup := v1.Group("/auth")
		{
			sc := new(auth.SignupController)
			authGroup.POST("/setup/phone/exist", sc.IsPhoneExist)
			authGroup.POST("/setup/email/exist", sc.IsEmailExist)
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-code/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-code/phone", vcc.SendSmsCode)
			authGroup.POST("verify-code/email", vcc.SendEmailCode)
		}
	}
}
