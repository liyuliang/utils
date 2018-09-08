package access

import (
	"reflect"
	"strings"
)

func SetField(obj interface{}, name string, value interface{}) {

	object := reflect.ValueOf(obj)
	structValue := object.Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		if strings.ToLower(fieldName) == strings.ToLower(name) {

			structFieldValue := structValue.FieldByName(fieldName)

			if structFieldValue.IsValid() {

				if structFieldValue.CanSet() {
					structFieldType := structFieldValue.Type()
					val := reflect.ValueOf(value)

					if structFieldType == val.Type() {
						structFieldValue.Set(val)
					}
				}
			}
		}
	}
}

func SetMap(obj interface{}, m map[string]interface{}) {
	for k, v := range m {
		SetField(obj, k, v)
	}
}

func Set(obj interface{}, fill interface{}) {

	targetObj := reflect.ValueOf(obj)
	fillObj := reflect.ValueOf(fill)
	structElem := targetObj.Elem()
	filleStructElem := fillObj.Elem()
	structType := structElem.Type()

	for i := 0; i < structElem.NumField(); i++ {
		
		fieldName := structType.Field(i).Name
		structField := structElem.FieldByName(fieldName)
		fillStructField := filleStructElem.FieldByName(fieldName)
		if structField.IsValid() && structField.CanSet() {
			if structField.Type()  == fillStructField.Type() {
				structField.Set(fillStructField)
			}
		}
	}
}

