package category

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func GetItemById(id int) category {
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"

	_id := convert.IntToString(id)
	result := db.Call(GetCategoryById, []string{_id}, KEY, db.DML_SELECT)

	item := New()
	item.setId(id)
	item.Code = result[_id]["code"]
	item.Name = result[_id]["name"]
	item.Order = convert.StringToInt(result[_id]["order"])
	item.ParentCategory = convert.StringToInt((result[_id]["parent_id"]))
	item.Enabled = convert.StringToBool(result[_id]["enabled"])
	item.CreatedAt = result[_id]["createdAt"]
	item.UpdatedAt = result[_id]["updatedAt"]

	return item
}

func GetItemByCode(code string) category {
	db.Connect()
	defer db.Disconnect()

	const KEY = "code"

	result := db.Call(GetCategoryByCode, []string{code}, KEY, db.DML_SELECT)

	item := New()
	item.setId(convert.StringToInt(result[code]["id"]))
	item.Code = code
	item.Name = result[code]["name"]
	item.Order = convert.StringToInt(result[code]["order"])
	item.ParentCategory = convert.StringToInt((result[code]["parent_id"]))
	item.Enabled = convert.StringToBool(result[code]["enabled"])
	item.CreatedAt = result[code]["createdAt"]
	item.UpdatedAt = result[code]["updatedAt"]

	return item
}

func GetItems(ids []int) (items []category) {
	for _, id := range ids {
		items = append(items, GetItemById(id))
	}

	return
}

func GetAllItems() (items []category){
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"

	result := db.Call(GetAllCategories, []string{}, KEY, db.DML_SELECT)

	for _, v := range result{
		item := New()
		item.Code = v["code"]
		item.Name = v["name"]
		item.Order = convert.StringToInt(v["order"])
		item.ParentCategory = convert.StringToInt((v["parent_id"]))
		item.Enabled = convert.StringToBool(v["enabled"])
		item.CreatedAt = v["createdAt"]
		item.UpdatedAt = v["updatedAt"]

		items = append(items, item)
	}

	return
}