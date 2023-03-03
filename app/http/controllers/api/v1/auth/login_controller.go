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
		response.Error(c, err, "帐号不存在")
	}
	token, expires := jwt.NewJwt().IsuseToken(user.GetStringId(), user.Name)
	response.Data(c, gin.H{
		"token":    token,
		"expires":  expires,
		"userinfo": user,
	})
}

func (lc *LoginController) LoginByMulti(c *gin.Context) {
	request := requests.LoginByMultiRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByMulti); !ok {
		return
	}
	user, err := auth.Attempt(request.LoginId, request.Password)
	if err != nil {
		response.Unauthorized(c, "用户名或密码错误")
	} else {
		token, expires := jwt.NewJwt().IsuseToken(user.GetStringId(), request.LoginId)
		response.Data(c, gin.H{
			"token":    token,
			"expires":  expires,
			"userinfo": user,
		})
	}
}

func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, expires, err := jwt.NewJwt().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "token刷新失败")
	}
	response.Data(c, gin.H{
		"token":   token,
		"expires": expires,
	})
}
