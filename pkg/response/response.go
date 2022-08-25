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
		"code": http.StatusOK,
		"data": nil,
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
		"code":    http.StatusNotFound,
		"message": defaultMessage("404 Not found", msg...),
	})
}

func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"code":    http.StatusForbidden,
		"message": defaultMessage("403 Forbidden", msg...),
	})
}

func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": defaultMessage("500 Server error", msg...),
	})
}

func Abort429(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
		"code":    http.StatusTooManyRequests,
		"message": defaultMessage("429 Requests too many", msg...),
	})
}

func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": defaultMessage("400 Bad request", msg...),
		"errors":  err.Error(),
	})
}

func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	if err == gorm.ErrRecordNotFound {
		Abort404(c, "资源未找到")
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"message": defaultMessage("Parameter error", msg...),
		"errors":  err.Error(),
	})
}

func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": defaultMessage("401 Unauthorized", msg...),
	})
}

func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"message": defaultMessage("Parameter error"),
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
