package googleoauth

import (
	"errors"
	"fptugo/pkg/core"
)

// VerifyGoogleOAuth ...
func VerifyGoogleOAuth(token string) error {
	googleAuthURL := "https://www.googleapis.com/userinfo/v2/me"
	statusCode, _ := core.HTTPGet(googleAuthURL, token)

	if statusCode != 200 {
		return errors.New("Token Unauthorized")
	}

	return nil
}
