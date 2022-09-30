package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/category"
	"goapi/app/requests"
	"goapi/pkg/response"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}

	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := category.Paginate(c, 10)
	response.Data(c, gin.H{
		"list":  data,
		"pager": pager,
	})
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c, "分类不存在")
		return
	}

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()
	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

/*func (ctrl *CategoriesController) Delete(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyCategory(c, categoryModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}*/
