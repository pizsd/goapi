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
	if userModel.ID == 0 {
		response.Abort404(c, "用户不存在或已删除")
		return
	}
	response.Data(c, userModel)
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {
	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}
	userModel := auth.User(c)

	userModel.Name = request.Name
	userModel.City = request.City
	userModel.Introduction = request.Introduction
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
