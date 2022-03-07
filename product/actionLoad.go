package product

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func GetItem(id int) product{
	db.Connect()
	defer db.Disconnect()
	
	const KEY = "id"

	_id := common.IntToString(id)
	result := db.Call(GetProduct, []string{_id}, KEY, db.DML_SELECT)

	item := New()
	item.setId(id)
	item.SetSku(result[_id]["sku"])
	item.SetName(result[_id]["name"])
	item.SetPrice(common.StringToFloat(result[_id]["price"]))
	item.SetDescription(result[_id]["description"])
	item.setCreatedAt(result[_id]["createdAt"])
	item.setUpdatedAt(result[_id]["updatedAt"])
	
	return item
}