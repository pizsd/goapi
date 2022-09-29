package cmd

import (
	"fmt"
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
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	bootstrap.SetupRoute(engine)
	port := config.Get("app.port")
	err := engine.Run(":" + port)
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
	console.Success(fmt.Sprintf("%s is running at http://localhost:/%s . Press Ctrl+C to stop.", config.Get("app.name"), port))
}
