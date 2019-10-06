package radio

import (
	"fptugo/configs/redis"
	"fptugo/pkg/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SetRadio ...
func SetRadio(w http.ResponseWriter, r *http.Request) {
	req := utils.Request{ResponseWriter: w, Request: r}
	res := utils.Response{ResponseWriter: w}

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
	res := utils.Response{ResponseWriter: w}

	value, err := redis.Get("radios")
	if err != nil {
		logrus.Println(err.Error())
	}

	radiosObj := Radio{Radios: value}

	res.SendOK(radiosObj)
}
