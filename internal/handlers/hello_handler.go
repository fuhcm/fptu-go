package handlers

import (
	"net/http"

	"fptugo/pkg/utils"
)

// Info ...
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GetInfo ...
func GetInfo(w http.ResponseWriter, r *http.Request) {
	res := utils.Response{ResponseWriter: w}
	res.SendOK(Info{
		Name:    "fptugo",
		Version: "0.0.0",
	})
}
