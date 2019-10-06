package utils

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	ResponseWriter http.ResponseWriter
}

// SendOK ...
func (r *Response) SendOK(body interface{}) {
	setJSON(r.ResponseWriter)
	encodeJSON(r.ResponseWriter, body)
}

// SendCreated ...
func (r *Response) SendCreated(body interface{}) {
	setJSON(r.ResponseWriter)
	setHTTPStatus(r.ResponseWriter, http.StatusCreated)
	encodeJSON(r.ResponseWriter, body)
}

// SendNoContent ...
func (r *Response) SendNoContent() {
	setJSON(r.ResponseWriter)
	setHTTPStatus(r.ResponseWriter, http.StatusNoContent)
	encodeJSON(r.ResponseWriter, nil)
}

type badRequestMessage struct {
	Message string `json:"message"`
}

// SendBadRequest ...
func (r *Response) SendBadRequest(message string) {
	badMessage := badRequestMessage{Message: message}
	messageJSON, _ := json.Marshal(badMessage)
	http.Error(r.ResponseWriter, string(messageJSON), http.StatusBadRequest)
}

type notFoundMessage struct {
	Message string `json:"message"`
}

// SendNotFound ...
func (r *Response) SendNotFound() {
	notFoundMessage := notFoundMessage{Message: "Not found"}
	messageJSON, _ := json.Marshal(notFoundMessage)
	http.Error(r.ResponseWriter, string(messageJSON), http.StatusNotFound)
}

// SendNotImplemented ...
func (r *Response) SendNotImplemented() {
	http.Error(r.ResponseWriter, "Not Implemented", http.StatusNotImplemented)
}

func setHTTPStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func setJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func encodeJSON(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}
