package common

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func newRouter(){
	router = mux.NewRouter()
}

func GetRouter() (r *mux.Router){
	if router == nil{
		newRouter()
	}
	r = router
	return
}

func Listen(port int, r *mux.Router){
	log.Fatal(http.ListenAndServe(":" + IntToString(port), r))
}