package handlers

import (
	"net/http"

	"boilerplate/core"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type authenticateBody struct {
	Token string `json:"token"`
}

// TokenAuthenticate ...
func TokenAuthenticate(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	authenticateBody := new(authenticateBody)
	req.GetJSONBody(&authenticateBody)

	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, body := core.HTTPGet(googleAuthURL, authenticateBody.Token)

	if statusCode != 200 {
		res.SendBadRequest("Unauthorized")
		return
	}

	res.SendOK(body)
}

type usernamePasswordBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authenticateResponseBody struct {
	Token string `json:"token"`
}

// UsernamePasswordAuthenticate ...
func UsernamePasswordAuthenticate(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	usernamePasswordBody := new(usernamePasswordBody)
	req.GetJSONBody(&usernamePasswordBody)

	if usernamePasswordBody.Username == "admin" && usernamePasswordBody.Password == "admin" {
		mySigningKey := []byte("tuhm")

		claims := &jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signed, err := token.SignedString(mySigningKey)
		if err != nil {
			logrus.Fatal(err)
		}

		res.SendOK(authenticateResponseBody{Token: signed})
		return
	}

	res.SendBadRequest("Unauthorized")
}
