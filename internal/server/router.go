package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"fptugo/internal/handlers"
	"fptugo/pkg/websocket"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Handle 404
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	// Handle websocket
	hub := websocket.NewHub()
	go hub.Run()
	router.Path("/ws").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	router.Methods("GET").Path("/").HandlerFunc(handlers.GetInfo)

	// Authenticate
	router.Methods("POST").Path("/auth").HandlerFunc(handlers.UsernamePasswordAuthenticate)
	router.Methods("POST").Path("/auth/oauth").HandlerFunc(handlers.TokenAuthenticate)
	router.Methods("POST").Path("/auth/new").HandlerFunc(handlers.CreateNewUser)

	// Users
	router.Methods("GET").Path("/users").HandlerFunc(handlers.ListUsers)

	// Confession
	router.Methods("GET").Path("/confessions").HandlerFunc(handlers.ListConfessions)

	return router
}
