package main

import (
	"ecommerce/common/db"
	"ecommerce/product"
	"fmt"
	"time"
)

func main(){
	dbParams := make(map[string]string)
	dbParams["username"] = "root"
	dbParams["password"] = "penta0611"
	dbParams["address"] = "127.0.0.1"
	dbParams["port"] = "3306"
	dbParams["dbName"] = "ecommerce"
	
	db.SetOption(dbParams)

	item := product.New()
	item.SetSku("insertTest5")
	item.SetName("Test Product")
	item.SetPrice(1.12345678)
	item.SetDescription("This is for testing saving")
	item.SetColor("black")
	item.SetCreatedAt(time.Now().Format("2022-03-02 22:28:30"))
	item.SetEnabled(true)

	fmt.Println(item)

	item.SetItem()
}