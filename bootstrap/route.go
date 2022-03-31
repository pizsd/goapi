package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/app/http/middlewares"
	"github.com/pizsd/goapi/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(r)
	// 注册路由
	routes.RegisterApiRoutes(r)
	// 处理404
	setupNotFoundHandler(r)
}

func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(middlewares.Logger(), gin.Recovery())
}

func setupNotFoundHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		acceptStr := c.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			c.String(http.StatusNotFound, "页面不存在")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"errcode": http.StatusNotFound,
				"errmsg":  "Not Found",
			})
		}
	})
}
