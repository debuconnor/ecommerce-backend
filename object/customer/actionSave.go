package customer

import (
	"ecommerce-backend/core/convert"
	"ecommerce-backend/core/db"
)

func (user *customer) Create() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	newId := "1"

	createResult := db.Call(CreateCustomer, []string{user.Email}, KEY, db.DML_INSERT)
	for k := range createResult {
		newId = k
		break
	}

	db.Call(SetName, []string{newId, user.Firstname, user.Middlename, user.Lastname}, KEY, db.DML_INSERT)
	db.Call(SetUsername, []string{newId, user.Username}, KEY, db.DML_INSERT)
	db.Call(SetPhone, []string{newId, user.Phone}, KEY, db.DML_INSERT)
	db.Call(SetAddress, []string{newId, user.Address1, user.Address2, user.Address3}, KEY, db.DML_INSERT)
	db.Call(SetCreatedAt, []string{newId}, KEY, db.DML_INSERT)
	db.Call(SetUpdatedAt, []string{newId}, KEY, db.DML_INSERT)
}

func (user *customer) Update() {
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	id := convert.IntToString(user.GetId())

	db.Call(SetName, []string{id, user.Firstname, user.Middlename, user.Lastname}, KEY, db.DML_UPDATE)
	db.Call(SetUsername, []string{id, user.Username}, KEY, db.DML_UPDATE)
	db.Call(SetPhone, []string{id, user.Phone}, KEY, db.DML_UPDATE)
	db.Call(SetAddress, []string{id, user.Address1, user.Address2, user.Address3}, KEY, db.DML_UPDATE)
	db.Call(SetUpdatedAt, []string{id}, KEY, db.DML_UPDATE)
}
