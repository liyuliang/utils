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

func Set(obj interface{}, m map[string]interface{}) {
	for k, v := range m {
		SetField(obj, k, v)
	}
}
