package access

import (
	"reflect"
	"strings"
)

func SetField(obj interface{}, name string, value interface{}) {
	name = strings.Title(strings.ToLower(name))

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
