package handlers

import (
	"net/http"

	"boilerplate/core"
)

type authenticateBody struct {
	Token string `json:"token"`
}

type authenticateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// TokenAuthenticate ...
func TokenAuthenticate(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	authenticateBody := new(authenticateBody)
	req.GetJSONBody(authenticateBody)

	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, body := core.HTTPGet(googleAuthURL, authenticateBody.Token)

	if statusCode != 200 {
		res.SendBadRequest("Unauthorized")
		return
	}

	res.SendOK(body)
}
