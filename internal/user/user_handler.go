package user

import (
	"fptugo/pkg/utils"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Response ...
type Response struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	NickName  string     `json:"nickname"`
	Resolved  int        `json:"resolved"`
}

// Hash password
func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		// Log error
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// CreateUserHandler ...
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := utils.Request{ResponseWriter: w, Request: r}
	res := utils.Response{ResponseWriter: w}

	user := new(User)
	req.GetJSONBody(user)

	// Hash password
	user.Password = hashAndSalt([]byte(user.Password))

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendCreated(user)
}

// GetUserByIDHandler ...
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	req := utils.Request{ResponseWriter: w, Request: r}
	res := utils.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := User{
		ID: id,
	}

	if err := user.FetchByID(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendOK(user)
}

// UpdateUserHandler ...
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := utils.Request{ResponseWriter: w, Request: r}
	res := utils.Response{ResponseWriter: w}

	id, _ := req.GetVarID()

	user := new(User)
	req.GetJSONBody(user)
	user.ID = id

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(user)
}

// DeleteUserHandler ...
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	req := utils.Request{ResponseWriter: w, Request: r}
	res := utils.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := User{
		ID: id,
	}

	if err := user.Delete(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendNoContent()
}
