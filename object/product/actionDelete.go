package product

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func (item *product) Delete(){
	db.Connect()
	defer db.Disconnect()

	product_id := convert.IntToString(item.GetId())
	db.Call(DeleteProduct, []string{product_id}, "", db.DML_DELETE)
	item = nil
}