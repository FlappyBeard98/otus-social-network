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
