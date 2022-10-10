package common

import "reflect"

func GetFieldsValuesAsSlice(obj interface{}) (values []interface{}) {
	reflected := reflect.ValueOf(obj)

	if reflected.Kind() != reflect.Struct {
		reflected = reflect.ValueOf(obj).Elem()
	}

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

func Map[In any,Out any](in []In,maper func(In)Out) []Out {

	result := make([]Out,len(in))
	for _,item := range in {
		result = append(result, maper(item))
	}
	return result
}