package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"reflect"
)

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

	result := make([]Out,0)
	for _,item := range in {
		result = append(result, maper(item))
	}
	return result
}

func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}