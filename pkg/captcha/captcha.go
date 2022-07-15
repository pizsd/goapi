package captcha

import (
	"github.com/mojocn/base64Captcha"
	"goapi/pkg/app"
	"goapi/pkg/config"
	"goapi/pkg/redis"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}
		dirver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.maxskew"),
			config.GetInt("captcha.dotcount"),
		)
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(dirver, &store)
	})
	return internalCaptcha
}

func (c *Captcha) GenerateCaptcha() (id, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id, answer string) (match bool) {
	if !app.IsProd() && id == config.GetString("captcha.testing_key") {
		return true
	}
	return c.Base64Captcha.Verify(id, answer, false)
}
