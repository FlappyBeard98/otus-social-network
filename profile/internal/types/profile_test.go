package types

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	name = "test"
)

func TestUpsertProfileReturnsNotNil(t *testing.T) {
	sut, _ := NewProfile(
		1,
		name,
		name,
		1,
		1,
		name,
		name)
	act := sut.UpsertProfile()
	assert.NotNil(t, act)
}

func TestNewProfileReturnsNotNil(t *testing.T) {
	act, _ := NewProfile(
		1,
		name,
		name,
		1,
		1,
		name,
		name)
	assert.NotNil(t, act)
}

func TestNewProfileWithLongFirstNameReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		gofakeit.LetterN(101),
		name,
		1,
		1,
		name,
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithLongLastNameReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		gofakeit.LetterN(101),
		1,
		1,
		name,
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithLongCityReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		name,
		1,
		1,
		gofakeit.LetterN(51),
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithLongHobbiesReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		name,
		1,
		1,
		name,
		gofakeit.LetterN(5001))
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithLessThanZeroAgeReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		name,
		-1,
		1,
		gofakeit.LetterN(51),
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithLessThanZeroGenderReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		name,
		1,
		-1,
		gofakeit.LetterN(51),
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewProfileWithMoreThanTwoGenderReturnsError(t *testing.T) {
	_, err := NewProfile(
		1,
		name,
		name,
		1,
		3,
		gofakeit.LetterN(51),
		name)
	assert.ErrorIs(t, err, ErrInvalidInput)
}
