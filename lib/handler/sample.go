package handler

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/hiroyukky/gae.go.template/lib/utils"
)

var SampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Debugf(c, "SampleHandler start")

	type ResponseJWT struct {
		Result bool `json:"result"`
	}

	res := ResponseJWT{
		Result: true,
	}
	utils.ResponseJson(w, res)
})
