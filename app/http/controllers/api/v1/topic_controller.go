package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/topic"
	"goapi/app/requests"
	"goapi/pkg/auth"
	"goapi/pkg/response"
	"strconv"
)

type TopicsController struct {
	BaseApiController
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	topics := topic.All()
	response.Data(c, topics)
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicModel)
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	userModel := auth.User(c)
	cid, _ := strconv.ParseInt(request.CategoryID, 10, 64)
	topicModel := topic.Topic{
		Title:      request.Title,
		Content:    request.Content,
		CategoryID: cid,
		UserID:     userModel.ID,
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c, "话题不存在")
		return
	}
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	cid, _ := strconv.ParseInt(request.CategoryID, 10, 64)
	topicModel.Title = request.Title
	topicModel.Content = request.Content
	topicModel.CategoryID = cid
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

/*func (ctrl *TopicsController) Delete(c *gin.Context) {

    topicModel := topic.Get(c.Param("id"))
    if topicModel.ID == 0 {
        response.Abort404(c)
        return
    }

    if ok := policies.CanModifyTopic(c, topicModel); !ok {
        response.Abort403(c)
        return
    }

    rowsAffected := topicModel.Delete()
    if rowsAffected > 0 {
        response.Success(c)
        return
    }

    response.Abort500(c, "删除失败，请稍后尝试~")
}*/
