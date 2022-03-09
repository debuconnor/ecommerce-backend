package customer

import (
	"ecommerce/common"
	"ecommerce/common/db"
)

func (user *customer) Delete(){
	db.Connect()
	defer db.Disconnect()

	customer_id := common.IntToString(user.GetId())
	db.Call(DeleteCustomer, []string{customer_id}, "", db.DML_DELETE)
	user = nil
}