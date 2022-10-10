package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hendralatumeten/vehicles_rental/src/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "start",
	Short: "Aplication Start",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "0.0.0.0:8080"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}
		t := cors.AllowAll()
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      t.Handler(mainRoute),
		}
		fmt.Println("run on", addrs)

		srv.ListenAndServe()
		return nil
	} else {
		return err
	}

}
