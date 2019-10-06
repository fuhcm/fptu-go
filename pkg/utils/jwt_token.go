package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// CreateJWTToken ...
func CreateJWTToken(user string) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	claims := &jwt.StandardClaims{
		ExpiresAt: 0,
		Issuer:    user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}
