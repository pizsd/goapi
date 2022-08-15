package middlewares

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/user"
	"goapi/pkg/jwt"
	"goapi/pkg/response"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJwt().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, "Unauthorized")
			return
		}
		userModel := user.Find(claims.UserId)
		if userModel.ID == 0 {
			response.Abort404(c, "model is not found")
		}
		c.Set("uid", userModel.GetStringId())
		c.Set("name", userModel.Name)
		c.Set("user", userModel)
		c.Next()
	}
}
