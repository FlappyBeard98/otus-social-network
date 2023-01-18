package types

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	login    = "test_lgn"
	password = "test_pwd"
	short    = "test7ch"
)

func TestNewAuthReturnsHashedPassword(t *testing.T) {
	act, _ := NewAuth(login, password)
	assert.NotEqual(t, password, act.Password)
}

func TestNewAuthWithShortLoginReturnsError(t *testing.T) {
	_, err := NewAuth(short, password)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewAuthWithLongLoginReturnsError(t *testing.T) {
	long := gofakeit.LetterN(251)
	_, err := NewAuth(long, password)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestNewAuthWithShortPasswordReturnsError(t *testing.T) {
	_, err := NewAuth(login, short)
	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestPasswordEqualsWithSamePasswordReturnsTrue(t *testing.T) {
	sut, _ := NewAuth(login, password)
	act := sut.PasswordEquals(password)
	assert.True(t, act)
}

func TestPasswordEqualsWithOtherPasswordReturnsFalse(t *testing.T) {
	sut, _ := NewAuth(login, password)
	act := sut.PasswordEquals("other_password")
	assert.False(t, act)
}

func TestInsertAuthReturnsNotNil(t *testing.T) {
	sut, _ := NewAuth(login, password)
	act := sut.InsertAuth()
	assert.NotNil(t, act)
}

func TestUpdatePasswordReturnsNotNil(t *testing.T) {
	sut, _ := NewAuth(login, password)
	act := sut.UpdatePassword()
	assert.NotNil(t, act)
}
