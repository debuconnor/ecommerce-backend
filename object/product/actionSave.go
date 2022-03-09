package product

import (
	"ecommerce/core/convert"
	"ecommerce/core/db"
)

func (item *product) Create() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateProduct, []string{item.Sku}, KEY, db.DML_INSERT)
	for k := range createResult {
		newId = k
		break
	}

	db.Call(SetName, []string{newId, item.Name}, KEY, db.DML_INSERT)
	db.Call(SetPrice, []string{newId, convert.FloatToString(item.Price)}, KEY, db.DML_INSERT)
	db.Call(SetDescription, []string{newId, item.Description}, KEY, db.DML_INSERT)
	db.Call(SetEnabled, []string{newId, convert.BoolToString(item.Enabled)}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId}, KEY, db.DML_INSERT)
	db.Call(SetUpdatedAt, []string{newId}, KEY, db.DML_INSERT)
}

func (item *product) Update() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := convert.IntToString(item.GetId())

	db.Call(SetName, []string{id, item.Name}, KEY, db.DML_UPDATE)
	db.Call(SetPrice, []string{id, convert.FloatToString(item.Price)}, KEY, db.DML_UPDATE)
	db.Call(SetDescription, []string{id, item.Description}, KEY, db.DML_UPDATE)
	db.Call(SetEnabled, []string{id, convert.BoolToString(item.Enabled)}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}
