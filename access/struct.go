package access

import (
	"reflect"
	"errors"
	"strings"
)

func SetField(obj interface{}, name string, value interface{}) error {
	name = strings.Title(name)

	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return errors.New("No such field: " + name + " in obj")
	}

	if !structFieldValue.CanSet() {
		return errors.New("Cannot set " + name + " field value")
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func Set(obj interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
