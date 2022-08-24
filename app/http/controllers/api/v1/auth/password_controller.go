package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/http/controllers/api/v1"
	"goapi/app/models/user"
	"goapi/app/requests"
	"goapi/pkg/response"
)

type PasswordController struct {
	v1.BaseApiController
}

func (pc *PasswordController) PasswordResetByPhone(c *gin.Context) {
	request := requests.PasswordResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetPasswordByPhone); !ok {
		return
	}
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c, "用户不存在")
		return
	}
	userModel.Password = request.Password
	userModel.Save()
	response.Success(c)
}

func (pc *PasswordController) PasswordResetByEmail(c *gin.Context) {
	request := requests.PasswordResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetPasswordByEmail); !ok {
		return
	}
	userModel := user.GetByMulti(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c, "用户不存在")
		return
	}
	userModel.Password = request.Password
	userModel.Save()
	response.Success(c)
}
