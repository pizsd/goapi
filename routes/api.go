package routes

import (
	"github.com/gin-gonic/gin"
	controllers "goapi/app/http/controllers/api/v1"
	"goapi/app/http/controllers/api/v1/auth"
	"goapi/app/http/middlewares"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "PONG",
		})
	})
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("1000-H"))
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
			authGroup.POST("/verify-code/phone", middlewares.LimitPerRoute("5-H"), vcc.SendSmsCode)
			authGroup.POST("verify-code/email", middlewares.LimitPerRoute("5-H"), vcc.SendEmailCode)
			lc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", lc.LoginByPhone)
			authGroup.POST("/login/using-multi", lc.LoginByMulti)
			authGroup.POST("/login/refresh-token", middlewares.AuthJwt(), lc.RefreshToken)
			pc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.AuthJwt(), pc.PasswordResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.AuthJwt(), pc.PasswordResetByEmail)
		}
		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJwt(), uc.CurrentUser)
		userGroup := v1.Group("/users").Use(middlewares.AuthJwt())
		{
			userGroup.GET("", uc.Index)
			userGroup.GET("/:id", uc.Show)
		}

		cc := new(controllers.CategoriesController)
		cateGroup := v1.Group("/categories", middlewares.AuthJwt())
		{
			cateGroup.GET("", cc.Index)
			cateGroup.POST("", cc.Store)
			cateGroup.PUT("/:id", cc.Update)
		}
	}
}
