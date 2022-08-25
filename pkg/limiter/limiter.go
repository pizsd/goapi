package limiter

import (
	"github.com/gin-gonic/gin"
	libLimiter "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"strings"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

func CheckRate(c *gin.Context, key string, formatted string) (libLimiter.Context, error) {
	var context libLimiter.Context
	rate, err := libLimiter.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, libLimiter.StoreOptions{
		Prefix: config.GetString("app.name", "go-api"),
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	limiterObj := libLimiter.New(store, rate)
	if c.GetBool("limit-once") {
		return limiterObj.Peek(c, key)
	} else {
		c.Set("limit-once", true)
		return limiterObj.Get(c, key)
	}
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
