package handlers

import (
	"log"
	"net/http"

	"fptugo/core"
	"fptugo/db"
	"fptugo/models"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticateBody ...
type AuthenticateBody struct {
	Token string `json:"token"`
}

// TokenAuthenticate ...
func TokenAuthenticate(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	AuthenticateBody := new(AuthenticateBody)
	req.GetJSONBody(&AuthenticateBody)

	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, body := core.HTTPGet(googleAuthURL, AuthenticateBody.Token)

	if statusCode != 200 {
		res.SendBadRequest("Unauthorized")
		return
	}

	res.SendOK(body)
}

func comparePasswords(hashedPwd string, plainPwdStr string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	plainPwd := []byte(plainPwdStr)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// UsernamePasswordBody ...
type UsernamePasswordBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthenticateResponseBody ...
type AuthenticateResponseBody struct {
	Token string `json:"token"`
}

// UsernamePasswordAuthenticate ...
func UsernamePasswordAuthenticate(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	usernamePasswordBody := new(UsernamePasswordBody)
	req.GetJSONBody(&usernamePasswordBody)

	var foundUser models.User
	err := db.DB.QueryRowx("SELECT name, email, password, level FROM users WHERE email = ?", usernamePasswordBody.Email).StructScan(&foundUser)
	if err != nil {
		res.SendBadRequest("Email not found")
		return
	}

	isPasswordValid := comparePasswords(foundUser.Password, usernamePasswordBody.Password)
	if isPasswordValid == true {
		token, err := core.CreateJWTToken(foundUser)
		if err != nil {
			res.SendBadRequest(err.Error())
			return
		}

		res.SendOK(AuthenticateResponseBody{Token: token})
		return
	}

	res.SendBadRequest("Wrong authentication")
}
