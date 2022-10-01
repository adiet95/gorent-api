package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adiet95/gorent-api/src/helpers"
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
		var addrs string

		if pr := helpers.Godotenv("PORT"); pr != "" {
			addrs = "127.0.0.1:" + pr
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      mainRoute,
		}

		fmt.Println("Gorent is running on PORT", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
