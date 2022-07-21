package config

import "goapi/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			"expire_time":       config.Env("JWT_EXPIRE_TIME"),
			"max_refresh_time":  config.Env("JWT_MAX_REFRESH_TIME"),
			"debug_expire_time": config.Env("JWT_DEBUG_EXPIRE_TIME"),
		}
	})
}
