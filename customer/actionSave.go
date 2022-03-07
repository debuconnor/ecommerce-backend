package customer

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func (user *customer) Create(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateCustomer, []string{user.GetEmail()}, KEY, db.DML_INSERT)
	for k := range createResult{
		newId = k
		break
	}

	db.Call(SetName, []string{newId, user.GetFirstname(), user.GetMiddlename(), user.GetLastname()}, KEY, db.DML_INSERT)
	db.Call(SetUsername, []string{newId, user.GetUsername()}, KEY, db.DML_INSERT)
	db.Call(SetPhone, []string{newId, user.GetPhone()}, KEY, db.DML_INSERT)
	db.Call(SetAddress, []string{newId, user.GetAddress1(), user.GetAddress2(), user.GetAddress3()}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId}, KEY, db.DML_INSERT)
	db.Call(SetUpdatedAt, []string{newId}, KEY, db.DML_INSERT)
}

func (user *customer) Update(){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := common.IntToString(user.GetId())

	db.Call(SetName, []string{id, user.GetFirstname(), user.GetMiddlename(), user.GetLastname()}, KEY, db.DML_UPDATE)
	db.Call(SetUsername, []string{id, user.GetUsername()}, KEY, db.DML_UPDATE)
	db.Call(SetPhone, []string{id, user.GetPhone()}, KEY, db.DML_UPDATE)
	db.Call(SetAddress, []string{id, user.GetAddress1(), user.GetAddress2(), user.GetAddress3()}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}