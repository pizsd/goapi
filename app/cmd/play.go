package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path"
	"strings"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 测试代码
	ext := path.Ext("routes/api.go")
	name := strings.TrimSuffix("api.go", ext)
	fmt.Println(name)
}
