package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/user"
	"goapi/app/requests"
	"goapi/pkg/auth"
	"goapi/pkg/config"
	"goapi/pkg/file"
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

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {
	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}
	userModel := auth.User(c)

	userModel.Name = request.Email
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {
	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}
	userModel := auth.User(c)

	userModel.Name = request.Phone
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	userModel := auth.User(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(userModel.Name, request.Password)
	if err != nil {
		errs := make(map[string][]string)
		errs["password"] = []string{"原密码不正确"}
		// 失败，显示错误提示
		response.ValidationError(c, errs)
	} else {
		userModel.Password = request.NewPassword
		rowsAffected := userModel.Save()
		if rowsAffected > 0 {
			response.Success(c)
		} else {
			response.Abort500(c, "更新失败，请稍后尝试~")
		}
	}
}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := auth.User(c)
	currentUser.Avatar = config.GetString("app.url") + avatar
	currentUser.Save()

	response.Data(c, currentUser)
}
