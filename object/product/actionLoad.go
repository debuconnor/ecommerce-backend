package product

import (
	"ecommerce/core/convert"
	"ecommerce/core/db"
)

func GetItemById(id int) product {
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"

	_id := convert.IntToString(id)
	result := db.Call(GetProductById, []string{_id}, KEY, db.DML_SELECT)

	item := New()
	item.setId(id)
	item.Sku = result[_id]["sku"]
	item.Name = result[_id]["name"]
	item.Price = convert.StringToFloat(result[_id]["price"])
	item.Description = result[_id]["description"]
	item.Enabled = convert.StringToBool(result[_id]["enabled"])
	item.CreatedAt = result[_id]["createdAt"]
	item.UpdatedAt = result[_id]["updatedAt"]

	return item
}

func GetItemBySku(sku string) product {
	db.Connect()
	defer db.Disconnect()

	const KEY = "sku"

	result := db.Call(GetProductBySku, []string{sku}, KEY, db.DML_SELECT)

	item := New()
	item.setId(convert.StringToInt(result[sku]["id"]))
	item.Sku = sku
	item.Name = result[sku]["name"]
	item.Price = convert.StringToFloat(result[sku]["price"])
	item.Description = result[sku]["description"]
	item.Enabled = convert.StringToBool(result[sku]["enabled"])
	item.CreatedAt = result[sku]["createdAt"]
	item.UpdatedAt = result[sku]["updatedAt"]

	return item
}

func GetItems(ids []int) (items []product) {
	for _, id := range ids {
		items = append(items, GetItemById(id))
	}

	return
}
