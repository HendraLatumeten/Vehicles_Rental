package config

import "github.com/spf13/cobra"

var initCommend = cobra.Command{
	Short: "simple backend with golang",
}

func init() {
	initCommend.AddCommand(ServeCmd)
}
func Run(args []string) error {
	initCommend.SetArgs(args)
	return initCommend.Execute()
}
