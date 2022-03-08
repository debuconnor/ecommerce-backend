package product

import (
	"ecommerce/common"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func preventError2(){
	log.Println("")
}

func GetProductHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := common.StringToInt(vars["id"])
	item := GetItemById(key)

	json.NewEncoder(w).Encode(item)
}

func CreateProductHandler(w http.ResponseWriter, r  *http.Request){
	item := New()
	item.Sku = r.FormValue("sku")
	item.Name = r.FormValue("name")
	item.Price = common.StringToFloat(r.FormValue("price"))
	item.Description = r.FormValue("description")
	item.Enabled = common.StringToBool(r.FormValue("enabled"))

	item.Create()
}

func UpdateProductHandler(w http.ResponseWriter, r  *http.Request){
	sku := r.FormValue("sku")
	oldItem := GetItemBySku(sku)

	newitem := New()
	newitem.setId(oldItem.GetId())
	newitem.Sku = sku
	newitem.Name = r.FormValue("name")
	newitem.Price = common.StringToFloat(r.FormValue("price"))
	newitem.Description = r.FormValue("description")
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
	router.HandleFunc("/products/get/{id}", GetProductHandler).Methods("GET")
	router.HandleFunc("/products/create/", CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/update/", UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", DeleteProductHandler).Methods("DELETE")
}