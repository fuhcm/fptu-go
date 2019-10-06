package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// HTTPGet ...
func HTTPGet(url string, token string) (int, []byte) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Accept", "application/json")

	if err != nil {
		logrus.Fatal(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode, body
}
