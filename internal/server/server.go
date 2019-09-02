package server

import (
	"fptugo/internal/router"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

// NewServer ...
func NewServer() *negroni.Negroni {

	server := negroni.Classic()
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	// server.Use(middlewares.JWTMiddleware())

	server.UseHandler(router.NewRouter())

	return server
}
