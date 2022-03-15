package relation

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func GetProductsByCustomerId(customerId int) (productIds []int) {
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"
	_customerId := convert.IntToString(customerId)
	result := db.Call(GetCartItems, []string{_customerId}, KEY, db.DML_SELECT)

	for k := range result {
		productIds = append(productIds, convert.StringToInt(k))
	}

	return
}
