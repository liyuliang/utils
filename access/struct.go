package access

import (
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) {

	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

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

func Set(obj interface{}, m map[string]interface{}) {
	for k, v := range m {
		SetField(obj, k, v)
	}
}
