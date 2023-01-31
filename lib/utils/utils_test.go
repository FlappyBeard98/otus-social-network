package utils

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHashReturnsSameHashForSameString(t *testing.T) {
	str1 := "string"
	str2 := "string"
	h1 := GetHash(str1)
	h2 := GetHash(str2)
	assert.Equal(t, h1, h2)
}

func TestGetFieldsValuesAsSliceWithStructReturnsSliceWithStructValues(t *testing.T) {
	type test struct {
		X int
		Y string
	}
	arg := test{X: 1, Y: "1"}

	act := GetFieldsValuesAsSlice(arg)

	assert.Equal(t, len(act), 2)
	assert.Equal(t, arg.X, act[0])
	assert.Equal(t, arg.Y, act[1])
}

func TestGetFieldsValuesAsSliceWithPointerReturnsSliceWithStructValues(t *testing.T) {
	type test struct {
		X int
		Y string
	}
	arg := test{X: 1, Y: "1"}

	act := GetFieldsValuesAsSlice(&arg)

	assert.Equal(t, len(act), 2)
	assert.Equal(t, arg.X, act[0])
	assert.Equal(t, arg.Y, act[1])
}

func TestGetFieldsValuesAsSliceWithPointerOfPointerReturnsEmptySlice(t *testing.T) {
	type test struct {
		X int
		Y string
	}
	arg := &test{X: 1, Y: "1"}

	act := GetFieldsValuesAsSlice(&arg)

	assert.Equal(t, len(act), 0)
}

func TestRetryReturnsNotNil(t *testing.T) {

	retries := []time.Duration{time.Microsecond, 2 * time.Millisecond}

	act, err := Retry(func() (any, error) { return 1, nil }, retries...)

	assert.Nil(t, err)
	assert.NotNil(t, act)
}

func TestRetryWithSingleErrorReturnsNotNil(t *testing.T) {

	retries := []time.Duration{time.Microsecond, 2 * time.Millisecond}
	iteration := 0

	act, err := Retry(func() (any, error) {
		iteration++
		if iteration < 2 {
			return nil, errors.New("test")
		} else {
			return 1, nil
		}
	}, retries...)

	assert.Nil(t, err)
	assert.NotNil(t, act)
}

func TestRetryWithErrorOnAllRetriesReturnsError(t *testing.T) {

	retries := []time.Duration{time.Microsecond, 2 * time.Millisecond}
	fn := func() (any, error) { return nil, errors.New("test") }

	act, err := Retry(fn, retries...)

	assert.NotNil(t, err)
	assert.Nil(t, act)
}
