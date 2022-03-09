package relation

import (
	"ecommerce/core/convert"
	"ecommerce/core/db"
)

func GetProductsByCategoryId(customerId int) (productIds []int) {
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"
	_customerId := convert.IntToString(customerId)
	result := db.Call(GetCustomerProducts, []string{_customerId}, KEY, db.DML_SELECT)

	for k := range result {
		productIds = append(productIds, convert.StringToInt(k))
	}

	return
}
