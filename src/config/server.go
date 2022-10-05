package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hendralatumeten/vehicles_rental/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "start",
	Short: "Aplication Start",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "127.0.0.1"

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = ":" + pr
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      mainRoute,
		}
		pr := os.Getenv("APP_PORT")
		fmt.Println("Gorent is running on PORT", pr)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}

}
