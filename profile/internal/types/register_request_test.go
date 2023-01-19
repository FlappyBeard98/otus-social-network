package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthReturnsNotNil(t *testing.T) {
	sut := RegisterRequest{
		Login:     validLogin,
		Password:  validPassword,
		FirstName: validFirstName,
		LastName:  validFirstName,
		Age:       validAge,
		Gender:    validGender,
		City:      validFirstName,
		Hobbies:   validFirstName,
	}
	act, _ := sut.NewAuth()
	assert.NotNil(t, act)
}

func TestNewProfileReturnsNotNil(t *testing.T) {
	sut := RegisterRequest{
		Login:     validLogin,
		Password:  validPassword,
		FirstName: validFirstName,
		LastName:  validFirstName,
		Age:       validAge,
		Gender:    validGender,
		City:      validFirstName,
		Hobbies:   validFirstName,
	}
	act, _ := sut.NewProfile()
	assert.NotNil(t, act)
}
