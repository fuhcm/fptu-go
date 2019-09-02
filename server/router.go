package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"fptugo/handlers"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Handle 404
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	router.Methods("GET").Path("/").HandlerFunc(handlers.GetInfo)

	// Authenticate
	router.Methods("POST").Path("/auth").HandlerFunc(handlers.UsernamePasswordAuthenticate)
	router.Methods("POST").Path("/auth/oauth").HandlerFunc(handlers.TokenAuthenticate)
	router.Methods("POST").Path("/auth/new").HandlerFunc(handlers.CreateNewUser)

	// Users
	router.Methods("GET").Path("/users").HandlerFunc(handlers.ListUsers)

	return router
}
