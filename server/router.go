package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"boilerplate/handlers"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Handle 404
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	router.Methods("GET").Path("/").HandlerFunc(handlers.GetInfo)

	// Authenticate
	router.Methods("POST").Path("/auth").HandlerFunc(handlers.TokenAuthenticate)

	// Testing routes
	router.Methods("GET").Path("/test/write").HandlerFunc(handlers.CreateNewUser)
	router.Methods("GET").Path("/test/read").HandlerFunc(handlers.ReadUsers)

	return router
}
