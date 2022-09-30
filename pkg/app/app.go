package app

import (
	"goapi/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProd() bool {
	return config.Get("app.env") == "prod"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}

func IsDebug() bool {
	return config.GetBool("app.debug")
}

func TimenowInTimezone() time.Time {
	timezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(timezone)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
