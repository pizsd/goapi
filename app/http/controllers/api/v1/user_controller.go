package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/user"
	"goapi/pkg/auth"
	"goapi/pkg/response"
)

type UsersController struct {
	BaseApiController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.User(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
