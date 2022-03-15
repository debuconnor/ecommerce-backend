package category

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/route"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetItemById(key)

	json.NewEncoder(w).Encode(item)
}

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request){
	items := GetAllItems()

	json.NewEncoder(w).Encode(items)	
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	item := New()
	item.Code = r.FormValue("code")
	item.Name = r.FormValue("name")
	item.Order = convert.StringToInt(r.FormValue("order"))
	item.Enabled = convert.StringToBool(r.FormValue("enabled"))

	item.Create()
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	oldItem := GetItemByCode(code)

	newitem := New()
	newitem.setId(oldItem.GetId())
	newitem.Code = code
	newitem.Name = r.FormValue("name")
	newitem.Order = convert.StringToInt(r.FormValue("order"))
	newitem.Enabled = convert.StringToBool(r.FormValue("enabled"))

	newitem.Update()
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetItemById(key)

	item.Delete()
}

func AddHandlers() {
	router := route.GetRouter()
	router.HandleFunc("/categories/get/{id}", GetCategoryHandler).Methods("GET")
	router.HandleFunc("/categories/get", GetAllCategoriesHandler).Methods("GET")
	router.HandleFunc("/categories/create/", CreateCategoryHandler).Methods("POST")
	router.HandleFunc("/categories/update/", UpdateCategoryHandler).Methods("PUT")
	router.HandleFunc("/categories/delete/{id}", DeleteCategoryHandler).Methods("DELETE")
}
