package types

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	validFirstName = "firstName"
	validLastName  = "lastName"
	validAge       = 21
	validGender    = 1
	validCity      = "city"
	validHobbies   = "hobbies"
	host           = "http://localhost:1323"
)

func TestUpsertProfileReturnsNotNil(t *testing.T) {
	sut, _ := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		validGender,
		validCity,
		validHobbies)
	act := sut.UpsertProfile()
	assert.NotNil(t, act)
}

func TestNewProfileReturnsValidProfile(t *testing.T) {
	act, _ := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		validGender,
		validCity,
		validHobbies)
	assert.NotNil(t, act)
}

func TestNewProfileWithLongFirstNameReturnsError(t *testing.T) {
	_, err := NewProfile(
		gofakeit.LetterN(101),
		validLastName,
		validAge,
		validGender,
		validCity,
		validHobbies)
	assert.Error(t, err)
}

func TestNewProfileWithLongLastNameReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		gofakeit.LetterN(101),
		validAge,
		validGender,
		validCity,
		validHobbies)
	assert.Error(t, err)
}

func TestNewProfileWithLongCityReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		validGender,
		gofakeit.LetterN(51),
		validHobbies)
	assert.Error(t, err)
}

func TestNewProfileWithLongHobbiesReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		validGender,
		validCity,
		gofakeit.LetterN(5001))
	assert.Error(t, err)
}

func TestNewProfileWithLessThanZeroAgeReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		validLastName,
		-1,
		validGender,
		validCity,
		validHobbies)
	assert.Error(t, err)
}

func TestNewProfileWithLessThanZeroGenderReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		-1,
		validCity,
		validHobbies)
	assert.Error(t, err)
}

func TestNewProfileWithMoreThanTwoGenderReturnsError(t *testing.T) {
	_, err := NewProfile(
		validFirstName,
		validLastName,
		validAge,
		3,
		validCity,
		validHobbies)
	assert.Error(t, err)
}
