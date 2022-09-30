package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/user"
	"goapi/app/requests"
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
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := user.Paginate(c, 10)
	response.Data(c, gin.H{
		"list":  data,
		"pager": pager,
	})
}
