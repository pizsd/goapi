package verifycode

import (
	"github.com/pizsd/goapi/pkg/app"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (rs *RedisStore) Get(key string, clear bool) string {
	key = rs.KeyPrefix + key
	val := rs.RedisClient.Get(key)
	if clear {
		rs.RedisClient.Del(key)
	}
	return val
}

func (rs *RedisStore) Verify(key string, answer string, clear bool) bool {
	val := rs.Get(key, clear)
	return val == answer
}

func (rs *RedisStore) Set(id string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt("verifycode.debug_expire_time"))
	}
	return rs.RedisClient.Set(rs.KeyPrefix+id, value, ExpireTime)
}
