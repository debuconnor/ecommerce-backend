package product

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func getNewId() (id int){
	//Call this method after opening DB
	const KEY = "id"

	result := db.Get("SELECT * from product order by id desc limit 1", KEY)
	
	for k := range result{
		id = common.StringToInt(k) + 1
		break
	}

	return
}

func (item *product) Create(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateProduct, []string{item.GetSku()}, KEY, db.DML_INSERT)
	for k := range createResult{
		newId = k
		break
	}
	db.Call(SetName, []string{newId, item.GetName()}, KEY, db.DML_INSERT)
	db.Call(SetPrice, []string{newId, common.FloatToString(item.GetPrice())}, KEY, db.DML_INSERT)
	db.Call(SetDescription, []string{newId, item.GetDescription()}, KEY, db.DML_INSERT)
	db.Call(SetColor, []string{newId, item.GetColor()}, KEY, db.DML_INSERT)
	db.Call(SetEnabled, []string{newId, common.BoolToString(item.GetEnabled())}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId, }, KEY, db.DML_INSERT)
}

func (item *product) Update(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := common.IntToString(item.GetId())

	db.Call(SetName, []string{id, item.GetName()}, KEY, db.DML_UPDATE)
	db.Call(SetPrice, []string{id, common.FloatToString(item.GetPrice())}, KEY, db.DML_UPDATE)
	db.Call(SetDescription, []string{id, item.GetDescription()}, KEY, db.DML_UPDATE)
	db.Call(SetColor, []string{id, item.GetColor()}, KEY, db.DML_UPDATE)
	db.Call(SetEnabled, []string{id, common.BoolToString(item.GetEnabled())}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}