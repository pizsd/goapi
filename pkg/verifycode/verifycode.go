package verifycode

import (
	"fmt"
	"goapi/pkg/app"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mail"
	"goapi/pkg/redis"
	"goapi/pkg/sms"
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

func (vc *VerifyCode) SendSmsCode(phone string) bool {
	code := vc.GenerateCode(phone)
	if !app.IsProd() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	s := sms.NewSms()
	return s.Send(phone, sms.Message{
		Template: s.Driver.Config()["template_code"],
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) SendEmailCode(email string) bool {
	code := vc.GenerateCode(email)
	if !app.IsProd() && strings.HasPrefix(email, config.GetString("verifycode.debug_email_suffix")) {
		return true
	}
	e := mail.NewMailer()
	content := fmt.Sprintf("<h1>您的Email验证码是：%v</h1>", code)
	return e.Send(mail.Email{
		From: mail.From{
			Address: e.Driver.Config()["address"],
			Name:    e.Driver.Config()["name"],
		},
		To:      []string{email},
		Subject: "GoApi 邮件验证码",
		HTML:    []byte(content),
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
