package product

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func GetItem(id int) map[string]map[string]string{
	db.Connect()
	defer db.Disconnect()
	
	const KEY = "id"

	_id := common.IntToString(id)
	result := db.Get("select * from product where id = " + _id, KEY)

	return result
}