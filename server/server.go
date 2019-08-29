package server

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

// NewServer ...
func NewServer() *negroni.Negroni {

	server := negroni.Classic()
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	server.UseHandler(NewRouter())

	return server
}
