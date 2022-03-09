package category

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func (item *category) Delete(){
	db.Connect()
	defer db.Disconnect()

	category_id := common.IntToString(item.GetId())
	db.Call(DeleteCategory, []string{category_id}, "", db.DML_DELETE)
	item = nil
}