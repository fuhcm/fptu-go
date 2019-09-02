package core

import (
	"encoding/json"
	"net/http"
)

// Request ...
type Request struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// GetJSONBody ...
func (r *Request) GetJSONBody(model interface{}) {
	decoder := json.NewDecoder(r.Request.Body)
	decoder.Decode(&model)
}
