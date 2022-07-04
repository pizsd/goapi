package response

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Success(c *gin.Context) {
	JSON(c, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}

func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"data": data,
	})
}

func CreatedJosn(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("资源不存在，请确定请求正确", msg...),
	})
}

func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("权限不足，请确定您有对应的权限", msg...),
	})
}

func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}

func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。", msg...),
		"errors":  err.Error(),
	})
}

func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	if err == gorm.ErrRecordNotFound {
		Abort404(c, "无效的资源")
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("参数错误"),
		"error":   err.Error(),
	})
}

func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("未认证请求", msg...),
	})
}

func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("参数错误"),
		"errors":  errors,
	})

}

func defaultMessage(defaultMessage string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMessage
	}
	return
}
