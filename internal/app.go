package internal

import (
	"crypto/tls"
	"log"
	"net/http"
)

type App struct {
	server *http.Server
}

type ServerOptions struct {
	Address string
	Handler http.Handler
	Logger  *log.Logger
	TLS     *tls.Config
}

func NewApp(options ServerOptions) *App {
	server := &http.Server{
		Addr:      options.Address,
		Handler:   options.Handler,
		ErrorLog:  options.Logger,
		TLSConfig: options.TLS,
	}
	return &App{
		server: server,
	}
}

func CreateRouter() *http.ServeMux {

}

func (a *App) Serve() error {
	return a.server.ListenAndServe()
}
