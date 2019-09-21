package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"fptugo/internal/auth"
	"fptugo/internal/confession"
	"fptugo/internal/crawl"
	"fptugo/internal/handlers"
	"fptugo/internal/radio"
	"fptugo/internal/user"
	"fptugo/pkg/middlewares"
	"fptugo/pkg/websocket"
)

func tokenRequired(controller http.HandlerFunc) http.Handler {
	return middlewares.JWTMiddleware().Handler(http.HandlerFunc(controller))
}

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

	// API Version
	apiPath := "/api"
	apiVersion := "/v1"
	apiPrefix := apiPath + apiVersion

	// Auth routes
	router.Methods("POST").Path("/auth/login").HandlerFunc(auth.LoginHandler)
	router.Methods("POST").Path("/auth/login_facebook").HandlerFunc(auth.LoginHandlerWithoutPassword)

	// User
	router.Methods("POST").Path(apiPrefix + "/users").Handler(tokenRequired(user.CreateUserHandler))
	router.Methods("POST").Path(apiPrefix + "/users").HandlerFunc(user.CreateUserHandler)
	router.Methods("GET").Path(apiPrefix + "/users/{id}").Handler(tokenRequired(user.GetUserByIDHandler))
	router.Methods("PUT").Path(apiPrefix + "/users/{id}").Handler(tokenRequired(user.UpdateUserHandler))
	router.Methods("DELETE").Path(apiPrefix + "/users/{id}").Handler(tokenRequired(user.DeleteUserHandler))

	// Confession
	router.Methods("GET").Path(apiPrefix + "/admincp/confessions").Handler(tokenRequired(confession.GetAllConfessionsHandler))
	router.Methods("POST").Path(apiPrefix + "/confessions").HandlerFunc(confession.CreateConfessionHandler)
	router.Methods("POST").Path(apiPrefix + "/myconfess").HandlerFunc(confession.GetConfessionsBySenderHandler)
	router.Methods("GET").Path(apiPrefix + "/confessions/approved").HandlerFunc(confession.GetApprovedConfessionsHandler)
	router.Methods("GET").Path(apiPrefix + "/confessions/overview").HandlerFunc(confession.GetConfessionsOverviewHandler)
	router.Methods("PUT").Path(apiPrefix + "/admincp/confessions/approve").Handler(tokenRequired(confession.ApproveConfessionHandler))
	router.Methods("PUT").Path(apiPrefix + "/admincp/confessions/reject").Handler(tokenRequired(confession.RejectConfessionHandler))
	router.Methods("GET").Path(apiPrefix + "/confessions/search").HandlerFunc(confession.SearchConfessionsHandler)
	router.Methods("POST").Path(apiPrefix + "/push/sync").HandlerFunc(confession.SyncPushIDHandler)

	// Crawl
	router.Methods("GET").Path("/crawl/{name}").HandlerFunc(crawl.GetHomeFeedHandler)
	router.Methods("GET").Path("/crawl/{name}/{id}").HandlerFunc(crawl.GetPostFeedHandler)
	router.Methods("GET").Path("/gist").HandlerFunc(crawl.GetResolveGithubGist)

	// Radio
	router.Methods("GET").Path(apiPrefix + "/radios").HandlerFunc(radio.GetRadio)
	router.Methods("POST").Path(apiPrefix + "/radios").Handler(tokenRequired(radio.SetRadio))

	return router
}
