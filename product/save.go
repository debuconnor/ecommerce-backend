package product

import (
	"ecommerce/common"
	"ecommerce/common/db"
	"log"
)

func getNewId() (id int){
	//Call this method after opening DB
	const KEY = "id"

	result := db.Get("SELECT * from product order by id desc limit 1", KEY)
	
	for k, _ := range result{
		id = common.StringToInt(k) + 1
		break
	}

	return
}

func (item *product) SetItem(){
	db.Connect()
	defer db.Disconnect()

	newId := common.IntToString(getNewId())
	db.Save("INSERT into product (id, sku) values (" + newId + ", '" + item.GetSku() + "')")
	// db.Save("INSERT into product_string (product_id, attribute_id, `value`) values (2, " + newId + ", '" + item.GetName() + "')")
	// db.Save("INSERT into product_decimal (product_id, attribute_id, `value`) values (3, " + newId + ", '" + common.FloatToString(item.GetPrice()) + "')")
	// db.Save("INSERT into product_string (product_id, attribute_id, `value`) values (4, " + newId + ", '" + item.GetDescription() + "')")
	// db.Save("INSERT into product_string (product_id, attribute_id, `value`) values (5, " + newId + ", '" + item.GetColor() + "')")
	
	fields := common.GetStructFields(item)
	attributeId := getAllAttributeId("product")

	for _, fieldName := range fields{
		fieldType := common.GetStructFieldType(item, fieldName)
		query := "INSERT into product_" + fieldType + 
				" (product_id, attribute_id, `value`) values (" +
				attributeId[fieldName]["id"] + ", " + newId + ", '" +
				item[fieldName] + "')"
		db.Save(query)
	}
}

func (item *product) UpdateItem(){
	db.Connect()
	defer db.Disconnect()

	//UPDATE...
}