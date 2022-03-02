package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Test(){
	db, err = sql.Open("mysql", "root:penta0611@tcp(127.0.0.1:3306)/ecommerce")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func isConnected() bool{
	return db != nil
}

func Connect(username *string, password *string, address *string, port *string, dbName *string){
	driverName := "mysql"
	dataSourceName := *username + ":" + *password + "@tcp(" + *address + ":" + *port + ")/" + *dbName

	db, err = sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect(){
	db.Close()
}

func Select(query string) map[string]interface{}{
	if !isConnected() {
		log.Fatal("Database disconnected.")
	}

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rawColumns, _ := rows.Columns()
	result := make(map[string]interface{})

	for rows.Next(){
		columns := make([]interface{}, len(rawColumns))
		columnPointers := make([]interface{}, len(rawColumns))

		for i, _ := range columns{
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil{
			log.Fatal(err)
		}

		for i, colName := range rawColumns{
			v := columnPointers[i].(*interface{})
			result[colName] = *v
		}
	}

	return result
}

