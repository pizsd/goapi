package auth

import (
	"fmt"
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
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	errs := requests.ValidateSignupEmailExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
