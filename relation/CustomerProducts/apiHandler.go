package relation

import (
	"ecommerce/core/convert"
	"ecommerce/object/product"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomerProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := convert.StringToInt(vars["id"])
	productIds := GetProductsByCustomerId(key)
	products := product.GetItems(productIds)

	json.NewEncoder(w).Encode(products)
}
