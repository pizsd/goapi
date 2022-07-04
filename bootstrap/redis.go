package bootstrap

import (
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.user"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
