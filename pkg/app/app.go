package app

import "goapi/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProd() bool {
	return config.Get("app.env") == "prod"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}
