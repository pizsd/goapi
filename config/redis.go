package config

import "github.com/pizsd/goapi/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("redis.host", "127.0.0.1"),
			"port":     config.Env("redis.port", "6379"),
			"password": config.Env("redis.password", ""),
			"database": config.Env("redis.db", 1),
		}
	})
}
