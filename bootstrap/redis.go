package bootstrap

import (
	"fmt"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.user"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
