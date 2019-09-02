package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"fptugo/internal/authen"
	"fptugo/internal/confession"
	"fptugo/internal/handlers"
	"fptugo/internal/user"
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
	router.Methods("POST").Path("/auth").HandlerFunc(authen.UsernamePasswordAuthenticate)
	router.Methods("POST").Path("/auth/oauth").HandlerFunc(authen.TokenAuthenticate)
	router.Methods("POST").Path("/auth/new").HandlerFunc(user.CreateNewUser)

	// Users
	router.Methods("GET").Path("/users").HandlerFunc(user.ListUsers)

	// Confession
	router.Methods("GET").Path("/confessions").HandlerFunc(confession.ListConfessions)

	return router
}
