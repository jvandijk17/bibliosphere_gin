package utils

import (
	"errors"
	"reflect"
)

func AssignDataToStruct(dest interface{}, data map[string]interface{}) error {
	destVal := reflect.ValueOf(dest).Elem()
	destType := destVal.Type()

	for i := 0; i < destVal.NumField(); i++ {
		field := destVal.Field(i)
		fieldType := destType.Field(i)

		jsonTag := fieldType.Tag.Get("json")

		if jsonTag == "" {
			continue
		}

		value, exists := data[jsonTag]
		if !exists {
			continue
		}

		fieldValue := reflect.ValueOf(value)
		if field.Kind() == fieldValue.Kind() {
			if fieldValue.Type().ConvertibleTo(field.Type()) {
				field.Set(fieldValue.Convert(field.Type()))
			} else {
				return errors.New("type mismatch for field: " + fieldType.Name)
			}
		}
	}

	return nil
}
