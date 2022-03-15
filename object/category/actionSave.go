package category

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func (item *category) Create() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateCategory, []string{item.Code, convert.IntToString(item.ParentCategory)}, KEY, db.DML_INSERT)
	for k := range createResult {
		newId = k
		break
	}

	db.Call(SetName, []string{newId, item.Name}, KEY, db.DML_INSERT)
	db.Call(SetOrder, []string{newId, convert.IntToString(item.Order)}, KEY, db.DML_INSERT)
	db.Call(SetEnabled, []string{newId, convert.BoolToString(item.Enabled)}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId}, KEY, db.DML_INSERT)
	db.Call(SetUpdatedAt, []string{newId}, KEY, db.DML_INSERT)
}

func (item *category) Update() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := convert.IntToString(item.GetId())

	db.Call(SetName, []string{id, item.Name}, KEY, db.DML_UPDATE)
	db.Call(SetOrder, []string{id, convert.IntToString(item.Order)}, KEY, db.DML_UPDATE)
	db.Call(SetEnabled, []string{id, convert.BoolToString(item.Enabled)}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}

func (item *category) SetParentId() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := convert.IntToString(item.GetId())
	catId := convert.IntToString(item.ParentCategory)

	db.Call(UpdateParentCategory, []string{id, catId}, KEY, db.DML_UPDATE)
}
