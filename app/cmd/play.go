package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 测试代码
	var i int64 = 1
	s := strconv.FormatInt(i, 10)
	fmt.Printf("%T - %[1]s\n", s)
}
