package config

import "github.com/pizsd/goapi/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			"code_length":        config.Env("VERIFY_CODE_LENGTH", 6),
			"expire_time":        config.Env("VERIFY_CODE_EXPIRE", 15),
			"debug_expire_time":  10080,
			"debug_code":         666888,
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
