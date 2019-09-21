package auth

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"fptugo/internal/user"
	"fptugo/pkg/core"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Params ...
type Params struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Response ...
type Response struct {
	JWT       string `json:"token"`
	ExpiresAt int64  `json:"expire_at"`
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
}

// Verify password
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	authParams := new(Params)
	req.GetJSONBody(authParams)

	user := user.User{
		Email: authParams.Email,
	}

	if err := user.FetchByEmail(); err != nil {
		res.SendBadRequest("User not found")
		return
	}

	verifyPassword := comparePasswords(user.Password, []byte(authParams.Password))

	if !verifyPassword {
		res.SendBadRequest("Password is wrong")
		return
	}

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expireTime,
		Id:        strconv.Itoa(user.ID),
		Subject:   user.Admin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		res.SendBadRequest("Unknown error")
		return
	}

	jwtResponse := Response{
		JWT:       ss,
		ExpiresAt: expireTime,
		ID:        user.ID,
		Nickname:  user.Nickname,
	}

	res.SendOK(jwtResponse)
}

// LoginHandlerWithoutPassword ...
func LoginHandlerWithoutPassword(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}
	authParams := new(Params)
	req.GetJSONBody(authParams)

	user := user.User{
		Email: authParams.Email,
	}

	if err := user.FetchByEmail(); err != nil {
		res.SendBadRequest("User not found")
		return
	}

	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, _ := core.HTTPGet(googleAuthURL, authParams.Token)

	if statusCode != 200 {
		res.SendBadRequest("Token Unauthorized")
		return
	}

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expireTime,
		Id:        strconv.Itoa(user.ID),
		Subject:   user.Admin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		res.SendBadRequest("Unknown error")
		return
	}

	jwtResponse := Response{
		JWT:       ss,
		ExpiresAt: expireTime,
		ID:        user.ID,
		Nickname:  user.Nickname,
	}

	res.SendOK(jwtResponse)
}
