package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/http/controllers/api/v1"
	"goapi/app/models/user"
	"goapi/app/requests"
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
