package relation

import (
	"ecommerce/core/route"
	cat_pro "ecommerce/relation/CategoryProducts"
	cust_pro "ecommerce/relation/CustomerProducts"
)

func AddHandlers() {
	router := route.GetRouter()
	router.HandleFunc("/categoryproducts/get/{id}", cat_pro.GetCategoryProductsHandler).Methods("GET")
	router.HandleFunc("/customerproducts/get/{id}", cust_pro.GetCustomerProductsHandler).Methods("GET")
}