package cmd

import (
	"github.com/spf13/cobra"
	"goapi/database/seeders"
	"goapi/pkg/console"
	"goapi/pkg/seed"
)

var CmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeed,
	Args:  cobra.MaximumNArgs(1),
}

func runSeed(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		name := args[0]
		sdr := seed.GetSeeder(name)
		if len(sdr.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
	}
	console.Success("Done seeding.")
}
