package types

import (
	"social-network/lib/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	validLogin    = "test_lgn"
	validPassword = "test_pwd"
)

func TestNewAuthReturnsValidAuth(t *testing.T) {
	hashedPassword := utils.GetHash(validPassword)

	act, _ := NewAuth(validLogin, validPassword)

	assert.NotEqual(t, validPassword, act.Password)
	assert.Equal(t, hashedPassword, act.Password)
	assert.Equal(t, validLogin, act.Login)
}

func TestNewAuthWithShortLoginReturnsError(t *testing.T) {
	_, err := NewAuth("test7ch", validPassword)
	assert.Error(t, err)
}

func TestNewAuthWithLongLoginReturnsError(t *testing.T) {
	long := gofakeit.LetterN(251)
	_, err := NewAuth(long, validPassword)
	assert.Error(t, err)
}

func TestNewAuthWithShortPasswordReturnsError(t *testing.T) {
	_, err := NewAuth(validLogin, "test7ch")
	assert.Error(t, err)
}

func TestInsertAuthReturnsNotNil(t *testing.T) {
	sut, _ := NewAuth(validLogin, validPassword)
	act := sut.InsertAuth()
	assert.NotNil(t, act)
}

func TestUpdatePasswordReturnsNotNil(t *testing.T) {
	sut, _ := NewAuth(validLogin, validPassword)
	act := sut.UpdatePassword()
	assert.NotNil(t, act)
}

func TestReadByLoginReturnsNotNil(t *testing.T) {
	sut, _ := NewAuth(validLogin, validPassword)
	act := sut.ReadByLogin()
	assert.NotNil(t, act)
}
