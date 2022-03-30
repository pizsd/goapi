package app

import "github.com/pizsd/goapi/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProd() bool {
	return config.Get("app.env") == "prod"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}
