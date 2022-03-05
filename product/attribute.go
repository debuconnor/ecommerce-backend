package product

import "ecommerce/common/db"

func CreateAttribute(attrCode string, attrName string){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	db.Call(CreateNewAttribute, []string{attrCode, attrName}, KEY, db.DML_INSERT)
}