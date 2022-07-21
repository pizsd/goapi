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
