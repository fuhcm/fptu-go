package googleoauth

import (
	"encoding/json"
	"errors"
	"fptugo/pkg/utils"
	"strings"
)

// TokenResponse ...
type TokenResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// VerifyGoogleOAuth ...
func VerifyGoogleOAuth(token string) (TokenResponse, error) {
	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, body := utils.HTTPGet(googleAuthURL, token)

	userData := TokenResponse{}
	json.Unmarshal(body, &userData)

	if statusCode != 200 {
		return TokenResponse{}, errors.New("Token Unauthorized")
	}

	return userData, nil
}

// IsFPTEduEmail ...
func IsFPTEduEmail(email string) bool {
	return strings.Contains(email, "fpt.edu.vn")
}
