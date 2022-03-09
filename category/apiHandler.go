package category

import (
	"ecommerce/common"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := common.StringToInt(vars["id"])
	item := GetItemById(key)

	json.NewEncoder(w).Encode(item)
}

func CreateProductHandler(w http.ResponseWriter, r  *http.Request){
	item := New()
	item.Code = r.FormValue("code")
	item.Name = r.FormValue("name")
	item.Order = common.StringToInt(r.FormValue("order"))
	item.Enabled = common.StringToBool(r.FormValue("enabled"))

	item.Create()
}

func UpdateProductHandler(w http.ResponseWriter, r  *http.Request){
	code := r.FormValue("code")
	oldItem := GetItemByCode(code)

	newitem := New()
	newitem.setId(oldItem.GetId())
	newitem.Code = code
	newitem.Name = r.FormValue("name")
	newitem.Order = common.StringToInt(r.FormValue("order"))
	newitem.Enabled = common.StringToBool(r.FormValue("enabled"))

	newitem.Update()
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := common.StringToInt(vars["id"])
	item := GetItemById(key)

	item.Delete()
}


func AddHandlers(){
	router := common.GetRouter()
	router.HandleFunc("/categories/get/{id}", GetProductHandler).Methods("GET")
	router.HandleFunc("/categories/create/", CreateProductHandler).Methods("POST")
	router.HandleFunc("/categories/update/", UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/categories/delete/{id}", DeleteProductHandler).Methods("DELETE")
}