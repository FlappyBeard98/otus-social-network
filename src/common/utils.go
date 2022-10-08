package common

import "reflect"

func GetFieldsValuesAsSlice(obj interface{}) (values []interface{}) {
	reflected := reflect.ValueOf(obj)

	if reflected.Kind() != reflect.Struct {
		return
	}
	l := reflected.NumField()

	for i := 0; i < l; i++ {
		field := reflected.Field(i)
		if field.CanInterface() {
			values = append(values, field.Interface())
		}
	}

	return
}
