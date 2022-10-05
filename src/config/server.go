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
		var addrs string = "127.0.0.1:8080"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = ":" + pr
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      mainRoute,
		}
		fmt.Println("Aplication run on http://", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}

}
