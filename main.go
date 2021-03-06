package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/bootstrap"
	btsConfig "goapi/config"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
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
	helpers.RandomNumber(6)
	err := engine.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
