package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/bootstrap"
	btsConfig "github.com/pizsd/goapi/config"
	"github.com/pizsd/goapi/pkg/captcha"
	"github.com/pizsd/goapi/pkg/config"
	"github.com/pizsd/goapi/pkg/logger"
)

func init() {
	// 可以使用匿名导入包，这样就可以不需要config/config.go了
	// 这里可能是显式的调用，可能是可读性更好
	btsConfig.Initialize()
}
func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(engine)
	logger.Dump(captcha.NewCaptcha().VerifyCaptcha("dm9vQ0gsZYmXoumdq8N9", "218949"), "正确的答案")
	err := engine.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
