package config

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm"
	"github.com/spf13/cobra"
)

var initCommend = cobra.Command{
	Short: "simple backend with golang",
}

func init() {
	initCommend.AddCommand(ServeCmd)
	initCommend.AddCommand(orm.MigrateCmd)
}
func Run(args []string) error {
	initCommend.SetArgs(args)
	return initCommend.Execute()
}
