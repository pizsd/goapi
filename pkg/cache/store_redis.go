package cache

import (
	"goapi/pkg/config"
	"goapi/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	Prefix      string
}

func NewRedisStore(addr string, user string, pass string, db int) *RedisStore {
	rs := &RedisStore{
		RedisClient: redis.NewClient(addr, user, pass, db),
		Prefix:      config.GetString("app.name") + ":cache:",
	}
	return rs
}

func (rs *RedisStore) Set(key, value string, expireTime time.Duration) {
	rs.RedisClient.Set(rs.Prefix+key, value, expireTime)
}

func (rs *RedisStore) Get(key string) string {
	return rs.RedisClient.Get(rs.Prefix + key)
}

func (rs *RedisStore) Has(key string) bool {
	return rs.RedisClient.Has(rs.Prefix + key)
}

func (rs *RedisStore) Forget(key string) {
	rs.RedisClient.Del(rs.Prefix + key)
}

func (rs *RedisStore) Forever(key, value string) {
	rs.RedisClient.Set(rs.Prefix+key, value, 0)
}

func (rs *RedisStore) Flush() {
	rs.RedisClient.FlushDB()
}

func (rs *RedisStore) IsAlive() error {
	return rs.RedisClient.Ping()
}

func (rs *RedisStore) Increment(parameters ...interface{}) {
	rs.RedisClient.Increment(parameters)
}

func (rs *RedisStore) Decrement(parameters ...interface{}) {
	rs.RedisClient.Decrement(parameters)
}
