package utils

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"time"
)

// GetHash returns sha256 hash string for input string
func GetHash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	v := h.Sum(nil)

	return string(v)
}

// GetFieldsValuesAsSlice returns the fields of the structure as an slice
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

func Retry[T any](fn func() (T, error), delays []time.Duration) (r T, err error) {

	f := func() (T, error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("%+v", r)
			}
		}()

		return fn()
	}

	for i := 0; i < len(delays); i++ {
		r, err = f()
		if err != nil {
			fmt.Printf("%+v", err)
		} else {
			return r, nil
		}
		time.Sleep(delays[i])
	}
	return *new(T), err
}
