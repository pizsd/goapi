package policies

import (
	"github.com/gin-gonic/gin"
	"goapi/app/models/topic"
	"goapi/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return _topic.UserID == auth.Uid(c)
}
