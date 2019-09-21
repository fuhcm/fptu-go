package radio

import (
	"fptugo/configs/redis"
	"fptugo/pkg/core"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SetRadio ...
func SetRadio(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	radioRequest := new(Radio)
	req.GetJSONBody(radioRequest)

	err := redis.Set("radios", radioRequest.Radios)
	if err != nil {
		logrus.Println(err.Error())
	}

	res.SendOK(radioRequest)
}

// GetRadio ...
func GetRadio(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}

	value, err := redis.Get("radios")
	if err != nil {
		logrus.Println(err.Error())
	}

	radiosObj := Radio{Radios: value}

	res.SendOK(radiosObj)
}
