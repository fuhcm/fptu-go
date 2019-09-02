package handlers

import (
	"net/http"

	"fptugo/pkg/core"
)

// NotFound ...
func NotFound(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}
	res.SendNotFound()
}
