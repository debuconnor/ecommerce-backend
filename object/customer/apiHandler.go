package customer

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/route"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetUserById(key)

	json.NewEncoder(w).Encode(item)
}

func CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	user := New()
	user.Email = r.FormValue("email")
	user.Firstname = r.FormValue("firstname")
	user.Middlename = r.FormValue("middlename")
	user.Lastname = r.FormValue("lastname")
	user.Username = r.FormValue("username")
	user.Phone = r.FormValue("phone")
	user.Address1 = r.FormValue("address1")
	user.Address2 = r.FormValue("address2")
	user.Address3 = r.FormValue("address3")
	user.CreatedAt = r.FormValue("createdAt")
	user.UpdatedAt = r.FormValue("updatedAt")

	user.Create()
}

func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	oldUser := GetUserByEmail(email)

	newUser := New()
	newUser.setId(oldUser.GetId())
	newUser.Firstname = r.FormValue("firstname")
	newUser.Middlename = r.FormValue("middlename")
	newUser.Lastname = r.FormValue("lastname")
	newUser.Username = r.FormValue("username")
	newUser.Phone = r.FormValue("phone")
	newUser.Address1 = r.FormValue("address1")
	newUser.Address2 = r.FormValue("address2")
	newUser.Address3 = r.FormValue("address3")
	newUser.CreatedAt = r.FormValue("createdAt")
	newUser.UpdatedAt = r.FormValue("updatedAt")

	newUser.Update()
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	item := GetUserById(key)

	item.Delete()
}

func AddHandlers() {
	router := route.GetRouter()
	router.HandleFunc("/customers/get/{id}", GetCustomerHandler).Methods("GET")
	router.HandleFunc("/customers/create/", CreateCustomerHandler).Methods("POST")
	router.HandleFunc("/customers/update/", UpdateCustomerHandler).Methods("PUT")
	router.HandleFunc("/customers/delete/{id}", DeleteCustomerHandler).Methods("DELETE")
}
