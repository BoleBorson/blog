package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"

	"github.com/BoleBorson/blog/internal"
)

func main() {
	options := internal.ServerOptions{
		Address: ":8080",
		Handler: http.NewServeMux(),
		Logger:  log.Default(),
		TLS:     &tls.Config{},
	}
	app := internal.NewApp(options)
	err := app.Serve()
	if err != nil {
		slog.Error(err.Error())
	}
}
