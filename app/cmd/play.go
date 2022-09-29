package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 测试代码
	m := make(map[string]string)
	// 最后给已声明的map赋值
	m["haha"] = "haha"
	m["hehe"] = "hehe"
	m["huhu"] = "valhu"

	if v, ok := m["huhu"]; ok {
		fmt.Println(v, ok)
	}
}
