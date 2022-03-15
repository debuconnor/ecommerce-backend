package relation

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func GetProductsByCategoryId(catId int) (productIds []int) {
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"
	_catId := convert.IntToString(catId)
	result := db.Call(GetCategoryProducts, []string{_catId}, KEY, db.DML_SELECT)

	for k := range result {
		productIds = append(productIds, convert.StringToInt(k))
	}

	return
}
