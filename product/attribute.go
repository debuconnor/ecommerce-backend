package product

import "ecommerce/common/db"

func getAllAttributeId(objectName string) map[string]map[string]string{
	const KEY = "code"

	result := db.Get("SELECT a.type, a.code " +
					"from attribute a " +
					"join attribute_type att on a.type = att.id " +
					"where att.type = '" + objectName + "'", KEY)

	return result
}

func CreateAttribute(attrCode string, attrName string){
	db.Connect()
	defer db.Disconnect()

	const KEY = ""
	db.Call(CreateNewAttribute, []string{attrCode, attrName}, KEY, db.DML_INSERT)
}