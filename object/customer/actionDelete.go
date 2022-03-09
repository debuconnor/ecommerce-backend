package customer

import (
	"ecommerce/core/convert"
	"ecommerce/core/db"
)

func (user *customer) Delete() {
	db.Connect()
	defer db.Disconnect()

	customer_id := convert.IntToString(user.GetId())
	db.Call(DeleteCustomer, []string{customer_id}, "", db.DML_DELETE)
	user = nil
}
