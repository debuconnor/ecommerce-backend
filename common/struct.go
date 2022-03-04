package common

import "reflect"

func GetStructFields(object interface{}) (fields []string){
	element := reflect.ValueOf(object).Elem()

	for i := 0; i < element.NumField(); i++{
		field := element.Type().Field(i).Name
		fields = append(fields, field)
	}

	return
}

func GetStructFieldType(object interface{}, fieldName string) (fieldType string){
	element := reflect.ValueOf(object).Elem()

	for i := 0; i < element.NumField(); i++{
		if field := element.Type().Field(i).Name; field == fieldName{
			fieldType = element.Type().Field(i).Type.String()
			
			if fieldType == "float64"{
				fieldType = "decimal"
			}
			break
		}
	}

	return
}