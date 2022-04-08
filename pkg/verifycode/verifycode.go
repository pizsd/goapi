package verifycode

import (
	"github.com/pizsd/goapi/pkg/app"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/helpers"
	"github.com/pizsd/goapi/pkg/logger"
	"github.com/pizsd/goapi/pkg/redis"
	"github.com/pizsd/goapi/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) SendCode(phone string) bool {
	code := vc.GenerateCode(phone)
	if !app.IsProd() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSms().Send(phone, sms.Message{
		Template: config.GetString("sms.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})
	if !app.IsProd() && (strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix")) || strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) GenerateCode(key string) string {
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length", 6))
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}
	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	vc.Store.Set(key, code)
	return code
}
