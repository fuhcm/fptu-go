package handlers

import (
	"net/http"

	"fptugo/pkg/utils"
)

// NotFound ...
func NotFound(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{ResponseWriter: w}
	res.SendNotFound()
}
