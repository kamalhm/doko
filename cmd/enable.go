package cmd

import (
	"github.com/egoist/doko/services"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable <service>",
	Short: "Enable a service",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "postgres":
			services.EnablePostgres(false)
		case "timescale":
			services.EnablePostgres(true)
		case "redis":
			services.EnableRedis()
		case "mysql":
			services.EnableMysql()
		case "chrome":
			services.EnableChrome()
		}
	},
}
