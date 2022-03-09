package category

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func GetItemById(id int) category{
	db.Connect()
	defer db.Disconnect()
	
	const KEY = "id"

	_id := common.IntToString(id)
	result := db.Call(GetCategoryById, []string{_id}, KEY, db.DML_SELECT)

	item := New()
	item.setId(id)
	item.Code = result[_id]["code"]
	item.Name = result[_id]["name"]
	item.Order = common.StringToInt(result[_id]["order"])
	item.ParentCategory = common.StringToInt((result[_id]["parent_id"]))
	item.Enabled = common.StringToBool(result[_id]["enabled"])
	item.CreatedAt = result[_id]["createdAt"]
	item.UpdatedAt = result[_id]["updatedAt"]
	
	return item
}

func GetItemByCode(code string) category{
	db.Connect()
	defer db.Disconnect()
	
	const KEY = "code"

	result := db.Call(GetCategoryByCode, []string{code}, KEY, db.DML_SELECT)

	item := New()
	item.setId(common.StringToInt(result[code]["id"]))
	item.Code = code
	item.Name = result[code]["name"]
	item.Order = common.StringToInt(result[code]["order"])
	item.ParentCategory = common.StringToInt((result[code]["parent_id"]))
	item.Enabled = common.StringToBool(result[code]["enabled"])
	item.CreatedAt = result[code]["createdAt"]
	item.UpdatedAt = result[code]["updatedAt"]
	
	return item
}

func GetItems(ids []int) (items []category){
	for _, id := range ids{
		items = append(items, GetItemById(id))
	}

	return
}