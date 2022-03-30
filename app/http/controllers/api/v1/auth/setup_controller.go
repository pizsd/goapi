package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/pizsd/goapi/app/http/controllers/api/v1"
	"github.com/pizsd/goapi/app/models/user"
	"net/http"
)

type SetupController struct {
	v1.BaseApiController
}

func (sc *SetupController) IsPhoneExist(c *gin.Context) {
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}
