package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/adiet95/gorent-api/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start aplikasi",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {

	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "127.0.0.1:8080"

		if pr := os.Getenv("_PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 20,
			ReadTimeout:  time.Second * 20,
			IdleTimeout:  time.Minute * 5,
			Handler:      mainRoute,
		}

		fmt.Println("Gorent is running on PORT", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
