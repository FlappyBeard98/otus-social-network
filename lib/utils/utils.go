package utils

import (
	"crypto/sha256"
	"log"
	"os/exec"
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

func Retry[T any](fn func() (T, error), delays []time.Duration) (T, error) {
	var r T
	var err error
	for i := 0; i < len(delays); i++ {
		r, err = fn()
		if err == nil {
			return r, nil
		}
		time.Sleep(delays[i])
	}
	return r, err
}

func ExecOrFail(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	o, err := cmd.Output()
	if err != nil {
		log.Printf("COMMAND: %v\n", cmd)
		log.Printf("OUTPUT : %v\n", string(o))
		log.Panic(err.Error())
	}
	log.Printf("%v", o)
}

func ComposeUp(path string) {
	ExecOrFail("docker-compose", "-f", path, "up", "--detach")
}

func ComposeDown(path string) {
	ExecOrFail("docker-compose", "-f", path, "down", "-v")
}

func ComposeMigrate(path string) {
	ExecOrFail("docker-compose", "-f", path, "build", "--no-cache")
}
