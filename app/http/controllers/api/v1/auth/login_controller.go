package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/http/controllers/api/v1"
	"goapi/app/requests"
	"goapi/pkg/auth"
	"goapi/pkg/jwt"
	"goapi/pkg/response"
)

type LoginController struct {
	v1.BaseApiController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "用户名或密码错误")
	}
	token := jwt.NewJwt().IsuseToken(user.GetStringId(), user.Name)
	response.JSON(c, gin.H{
		"userinfo": user,
		"token":    token,
	})
}
