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

func Fill(targetObj interface{}, fillObj interface{}) {

	targetObject := reflect.ValueOf(targetObj)
	fillObject := reflect.ValueOf(fillObj)

	targetObjStructValue := targetObject.Elem()
	fillObjStructValue := fillObject.Elem()

	targetObjStructType := targetObjStructValue.Type()
	fillObjStructType := fillObjStructValue.Type()

	for i := 0; i < targetObjStructType.NumField(); i++ {
		targetObjFieldName := targetObjStructType.Field(i).Name
		fillObjFieldName := fillObjStructType.Field(i).Name

		targetObjStructFieldValue := targetObjStructValue.FieldByName(targetObjFieldName)
		fillObjStructFieldValue := fillObjStructValue.FieldByName(fillObjFieldName)


		if targetObjStructFieldValue.IsValid() {
			if targetObjStructFieldValue.CanSet() {

				targetObjStructFieldType := targetObjStructFieldValue.Type()
				fillObjStructFieldVal := reflect.ValueOf(fillObjStructFieldValue)
				println(targetObjStructFieldType.String(), fillObjStructFieldVal.Type().String())
				if targetObjStructFieldType.String() == fillObjStructFieldVal.Type().String() {

					targetObjStructFieldValue.Set(fillObjStructFieldVal)
				}
			}
		}
		//println(targetObjFieldName, targetObjStructFieldValue.String())
		//println(fillObjFieldName, fillObjStructFieldValue.String())
	}

}
