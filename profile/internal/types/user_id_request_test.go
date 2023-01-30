package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProfileByUserIdReturnsNotNil(t *testing.T) {
	sut := UserIdRequest{}
	act := sut.ReadProfileByUserId()
	assert.NotNil(t, act)
}

func TestReadUserFriendsProfilesReturnsNotNil(t *testing.T) {
	sut := UserIdRequest{}
	pageRequest := PageRequest{}
	act := sut.ReadUserFriendsProfiles(&pageRequest)
	assert.NotNil(t, act)
}

func TestReadUserFriendsTotalReturnsNotNil(t *testing.T) {
	sut := UserIdRequest{}
	act := sut.ReadUserFriendsTotal()
	assert.NotNil(t, act)
}

func TestGetProfileRequestCreateRequestReturnsNotNil(t *testing.T) {
	sut := GetProfileRequest{}
	act, _ := sut.CreateRequest(host)
	assert.NotNil(t, act)
}


func TestGetFriendsRequestCreateRequestReturnsNotNil(t *testing.T) {
	sut := GetFriendsRequest{}
	act, _ := sut.CreateRequest(host)
	assert.NotNil(t, act)
}


