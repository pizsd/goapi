package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goapi/bootstrap"
	"goapi/pkg/config"
	"goapi/pkg/console"
	"goapi/pkg/logger"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	bootstrap.SetupRoute(engine)
	err := engine.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
