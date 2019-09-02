package handlers

import (
	"fmt"
	"net/http"

	"fptugo/core"
	"fptugo/db"
	"fptugo/models"

	"golang.org/x/crypto/bcrypt"
)

// ResponseMessage ...
type ResponseMessage struct {
	Message string `json:"message"`
}

// CreateNewUser ...
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	newUser := new(models.User)
	req.GetJSONBody(&newUser)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	fmt.Printf(string(hashedPassword))

	newUser.Password = string(hashedPassword)

	_, err := db.DB.NamedExec("INSERT INTO users VALUES(null, :name, :email, :password, :level)", newUser)
	if err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(ResponseMessage{
		Message: "Success",
	})
}

// ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	users := []models.User{}
	db.DB.Select(&users, "SELECT name, email, password, level FROM users")

	res.SendOK(users)
}
