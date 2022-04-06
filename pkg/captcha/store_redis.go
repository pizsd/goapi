package captcha

import (
	"errors"
	"github.com/pizsd/goapi/pkg/app"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (rs *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_tim"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := rs.RedisClient.Set(key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	key = rs.KeyPrefix + key
	val := rs.RedisClient.Get(key)
	if clear {
		rs.RedisClient.Del(key)
	}
	return val
}

func (rs *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}
