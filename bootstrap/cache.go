package bootstrap

import (
	"fmt"
	"goapi/pkg/cache"
	"goapi/pkg/config"
	"goapi/pkg/logger"
)

func SetupCache() {
	driver := config.GetString("cache.driver")
	switch driver {
	case "redis":
		rs := cache.NewRedisStore(
			fmt.Sprintf("%s:%s", config.GetString("redis.host"), config.GetString("redis.port")),
			config.GetString("redis.user"),
			config.GetString("password"),
			config.GetInt("redis.cache_database"),
		)
		cache.NewCache(rs)
	case "file":
		logger.ErrorString("cache", "file", "file driver not supported")
	}
}
