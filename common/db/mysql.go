package db

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const DML_SELECT = "SELECT"
const DML_UPDATE = "UPDATE"
const DML_INSERT = "INSERT"
const DML_DELETE = "DELETE"
const STORE_PROCEDURE = "CALL"

var db *sql.DB
var err error
var params map[string]string

func Test(from string){
	//log.Println("mysql.go [Test()] called from [" + from + "]")
}

func SetOption(options map[string]string){
	params = make(map[string]string)
	for k, v := range options{
		params[k] = v
	}
}

func Connect(){
	driverName := "mysql"
	dataSourceName := params["username"] + ":" + params["password"] + "@tcp(" + params["address"] + ":" + params["port"] + ")/" + params["dbName"]

	db, err = sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect(){
	db.Close()
}

func isConnected() bool{
	return db != nil
}

func checkError(e error) {
	if e != nil{
		log.Fatal(e)
	}
}

func validateQuery(query string, dml string) (result bool){
	result = true

	if !strings.HasPrefix(query, dml) && !strings.HasPrefix(query, STORE_PROCEDURE){
		result = false
	}

	if strings.Contains(query, ";"){
		result = false
	}

	return
}

func Get(query string, key string) map[string]map[string]string{
	if !isConnected() {
		log.Fatal("Database not connected.")
	}

	if !validateQuery(query, DML_SELECT){
		log.Fatal("Vaildate query failed")
	}
	
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError(err)

	result := make(map[string]map[string]string)

	for rows.Next(){
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		data := make(map[string]string)
		
		for i := range columns{
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols{
			data[colName] = columns[i]
		}

		id := data[key]
		delete(data, key)

		result[id] = make(map[string]string)
		result[id] = data
	}

	return result
}

func Save(query string) int{
	if !isConnected() {
		log.Fatal("Database not connected.")
	}

	if !validateQuery(query, DML_INSERT) && !validateQuery(query, DML_UPDATE){
		log.Fatal("Vaildate query failed")
	}

	result, err := db.Exec(query)
	checkError(err)

	rowCount, err := result.RowsAffected()
	checkError(err)

	return int(rowCount)
}

func Call(proc string, params []string, key string, dml string) (result map[string]map[string]string){
	if !isConnected() {
		log.Fatal("Database not connected.")
	}

	if dml == ""{
		log.Fatal("DML not selected when CALL")
	}

	if key == ""{
		key = "id"
	}

	query := "CALL " + proc + "("
	for _, v := range params{
		query += "'" + v + "'"
	}
	query += ")"

	result = Get(query, key)
	return
}