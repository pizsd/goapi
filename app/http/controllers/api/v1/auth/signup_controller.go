package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pizsd/goapi/app/http/controllers/api/v1"
	"github.com/pizsd/goapi/app/models/user"
	"github.com/pizsd/goapi/app/requests"
	"net/http"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
