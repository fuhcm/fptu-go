package handlers

import (
	"net/http"

	"boilerplate/core"
	"boilerplate/db"
	"boilerplate/models"
)

type responseMessage struct {
	Message string `json:"message"`
}

// CreateNewUser ...
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	_, err := db.DB.NamedExec("INSERT INTO users VALUES(null, :name, :email)", models.User{Name: "Tu Huynh", Email: "tuhmse62531@fpt.edu."})

	if err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(responseMessage{
		Message: "Success",
	})
}

// ReadUsers ...
func ReadUsers(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	users := []models.User{}
	db.DB.Select(&users, "SELECT name, email FROM users")

	res.SendOK(users)
}
