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

func (ctrl *UsersController) Show(c *gin.Context) {
	userModel := user.Find(c.Param("id"))
	if userModel.ID > 0 {
		response.Data(c, userModel)
		return
	}
	response.Abort404(c, "用户不存在或已删除")
}
