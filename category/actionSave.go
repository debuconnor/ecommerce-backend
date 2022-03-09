package category

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func (item *category) Create(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateCategory, []string{item.Code}, KEY, db.DML_INSERT)
	for k := range createResult{
		newId = k
		break
	}
	
	db.Call(SetName, []string{newId, item.Name}, KEY, db.DML_INSERT)
	db.Call(SetOrder, []string{newId, common.IntToString(item.Order)}, KEY, db.DML_INSERT)
	db.Call(SetEnabled, []string{newId, common.BoolToString(item.Enabled)}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId}, KEY, db.DML_INSERT)
	db.Call(SetUpdatedAt, []string{newId}, KEY, db.DML_INSERT)
}

func (item *category) Update(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := common.IntToString(item.GetId())

	db.Call(SetName, []string{id, item.Name}, KEY, db.DML_UPDATE)
	db.Call(SetOrder, []string{id, common.IntToString(item.Order)}, KEY, db.DML_UPDATE)
	db.Call(SetEnabled, []string{id, common.BoolToString(item.Enabled)}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}