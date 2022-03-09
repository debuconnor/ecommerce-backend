package category

import (
	"ecommerce/core/convert"
	"ecommerce/core/db"
)

func (item *category) Delete() {
	db.Connect()
	defer db.Disconnect()

	category_id := convert.IntToString(item.GetId())
	db.Call(DeleteCategory, []string{category_id}, "", db.DML_DELETE)
	item = nil
}
