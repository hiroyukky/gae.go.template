package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiroyukky/gae.go.template/lib/handler"
)

//var secretKey []byte = []byte(os.Getenv("JWT_SECRET"))

func init(){
	r := Router()
	http.Handle("/", r)
}

func Router() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/v{version:[0-9]+}/sample", 
		handler.SampleHandler).Methods("GET")

	return r
}
