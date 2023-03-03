package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/http/controllers/api/v1"
	"goapi/app/models/user"
	"goapi/app/requests"
	"goapi/pkg/jwt"
	"goapi/pkg/response"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		token, expires := jwt.NewJwt().IsuseToken(_user.GetStringId(), _user.Name)
		response.CreatedJosn(c, gin.H{
			"token":    token,
			"expires":  expires,
			"userinfo": _user,
		})
	} else {
		response.Abort500(c, "服务器错误，请稍候再试")
	}
}

func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		token, expires := jwt.NewJwt().IsuseToken(_user.GetStringId(), _user.Name)
		response.CreatedJosn(c, gin.H{
			"token":    token,
			"expires":  expires,
			"userinfo": _user,
		})
	} else {
		response.Abort500(c, "服务器错误，请稍候再试")
	}

}
