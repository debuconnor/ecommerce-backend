package route

import (
	"ecommerce-backend/core/convert"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":" + convert.IntToString(port), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}