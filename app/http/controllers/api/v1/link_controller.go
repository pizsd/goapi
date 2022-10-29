package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/link"
	"goapi/app/requests"
	"goapi/pkg/response"
)

type LinksController struct {
	BaseApiController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}

func (ctrl *LinksController) Show(c *gin.Context) {
	linkModel := link.Get(c.Param("id"))
	if linkModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, linkModel)
}

func (ctrl *LinksController) Store(c *gin.Context) {

	request := requests.LinkRequest{}
	if ok := requests.Validate(c, &request, requests.LinkSave); !ok {
		return
	}

	linkModel := link.Link{
		Name: request.Name,
		Url:  request.Url,
		Logo: request.Logo,
	}
	linkModel.Create()
	if linkModel.ID > 0 {
		response.Created(c, linkModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *LinksController) Update(c *gin.Context) {

	linkModel := link.Get(c.Param("id"))
	if linkModel.ID == 0 {
		response.Abort404(c)
		return
	}
	// TODO 判断是否有更改友情链接权限

	request := requests.LinkRequest{}
	if ok := requests.Validate(c, &request, requests.LinkSave); !ok {
		return
	}

	linkModel.Name = request.Name
	linkModel.Url = request.Url
	linkModel.Logo = request.Logo
	rowsAffected := linkModel.Save()
	if rowsAffected > 0 {
		response.Data(c, linkModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *LinksController) Delete(c *gin.Context) {

	linkModel := link.Get(c.Param("id"))
	if linkModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := linkModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
