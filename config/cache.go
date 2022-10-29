package config

import (
	"goapi/pkg/config"
)

func init() {
	config.Add("cache", func() map[string]interface{} {
		return map[string]interface{}{
			"driver": config.Env("CACHE_DRIVER", "file"),
			"file": map[string]interface{}{
				"path": "storage/cache/",
			},
		}
	})
}
