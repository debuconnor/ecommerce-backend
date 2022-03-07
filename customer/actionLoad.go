package customer

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func GetUser(id int) customer{
	db.Connect()
	defer db.Disconnect()

	const KEY = "id"

	_id := common.IntToString(id)
	result := db.Call(GetCustomer, []string{_id}, KEY, db.DML_SELECT)

	user := New()
	user.setId(id)
	user.SetEmail(result[_id]["email"])
	user.SetFirstname(result[_id]["firstname"])
	user.SetMiddlename(result[_id]["middlename"])
	user.SetLastname(result[_id]["lastname"])
	user.SetUsername(result[_id]["username"])
	user.SetPhone(result[_id]["phone"])
	user.SetAddress1(result[_id]["address1"])
	user.SetAddress2(result[_id]["address2"])
	user.SetAddress3(result[_id]["address3"])
	user.setCreatedAt(result[_id]["createdAt"])
	user.setUpdatedAt(result[_id]["updatedAt"])

	return user
}