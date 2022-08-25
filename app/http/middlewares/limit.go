package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"goapi/pkg/app"
	"goapi/pkg/limiter"
	"goapi/pkg/logger"
	"goapi/pkg/response"
)

func LimitIP(limit string) gin.HandlerFunc {
	if !app.IsProd() {
		limit = "10000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}
	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))
	if rate.Reached {
		response.Abort429(c)
		return false
	}
	return true
}
