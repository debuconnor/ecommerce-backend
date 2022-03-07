package product

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func (item *product) Delete(){
	db.Connect()
	defer db.Disconnect()

	product_id := common.IntToString(item.GetId())
	db.Call(DeleteProduct, []string{product_id}, "", db.DML_DELETE)
	item = nil
}