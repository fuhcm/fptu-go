package handlers

import (
	"net/http"

	"boilerplate/core"
)

// NotFound ...
func NotFound(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}
	res.SendNotFound()
}
