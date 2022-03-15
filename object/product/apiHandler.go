package product

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/route"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetItemById(key)

	json.NewEncoder(w).Encode(item)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	item := New()
	item.Sku = r.FormValue("sku")
	item.Name = r.FormValue("name")
	item.Price = convert.StringToFloat(r.FormValue("price"))
	item.Description = r.FormValue("description")
	item.Enabled = convert.StringToBool(r.FormValue("enabled"))

	item.Create()
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.FormValue("sku")
	oldItem := GetItemBySku(sku)

	newitem := New()
	newitem.setId(oldItem.GetId())
	newitem.Sku = sku
	newitem.Name = r.FormValue("name")
	newitem.Price = convert.StringToFloat(r.FormValue("price"))
	newitem.Description = r.FormValue("description")
	newitem.Enabled = convert.StringToBool(r.FormValue("enabled"))

	newitem.Update()
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetItemById(key)

	item.Delete()
}

func AddHandlers() {
	router := route.GetRouter()
	router.HandleFunc("/products/get/{id}", GetProductHandler).Methods("GET")
	router.HandleFunc("/products/create/", CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/update/", UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", DeleteProductHandler).Methods("DELETE")
}
