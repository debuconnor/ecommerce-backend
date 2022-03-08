package customer

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func GetUserById(id int) customer{
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"

	_id := common.IntToString(id)
	result := db.Call(GetCustomerById, []string{_id}, KEY, db.DML_SELECT)

	user := New()
	user.setId(id)
	user.Email = result[_id]["email"]
	user.Firstname = result[_id]["firstname"]
	user.Middlename = result[_id]["middlename"]
	user.Lastname = result[_id]["lastname"]
	user.Username = result[_id]["username"]
	user.Phone = result[_id]["phone"]
	user.Address1 = result[_id]["address1"]
	user.Address2 = result[_id]["address2"]
	user.Address3 = result[_id]["address3"]
	user.CreatedAt = result[_id]["createdAt"]
	user.UpdatedAt = result[_id]["updatedAt"]

	return user
}

func GetUserByEmail(email string) customer{
	db.Connect()
	defer db.Disconnect()

	const KEY = "email"

	result := db.Call(GetCustomerByEmail, []string{email}, KEY, db.DML_SELECT)

	user := New()
	user.setId(common.StringToInt(result[email]["id"]))
	user.Email = result[email]["email"]
	user.Firstname = result[email]["firstname"]
	user.Middlename = result[email]["middlename"]
	user.Lastname = result[email]["lastname"]
	user.Username = result[email]["username"]
	user.Phone = result[email]["phone"]
	user.Address1 = result[email]["address1"]
	user.Address2 = result[email]["address2"]
	user.Address3 = result[email]["address3"]
	user.CreatedAt = result[email]["createdAt"]
	user.UpdatedAt = result[email]["updatedAt"]

	return user
}