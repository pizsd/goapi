package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goapi/pkg/cache"
	"goapi/pkg/console"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
	Args:  cobra.NoArgs,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget --key|k=xxx",
	Run:   runCacheForget,
	Args:  cobra.NoArgs,
}

var cacheKey string

func init() {
	CmdCache.AddCommand(
		CmdCacheClear,
		CmdCacheForget,
	)
	// 设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}
func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
