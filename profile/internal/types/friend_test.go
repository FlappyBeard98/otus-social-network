package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validUserId   = 1
	validFriendId = 2
)

func TestNewFriendReturnsValidFriend(t *testing.T) {
	act, _ := NewFriend(validUserId, validFriendId)

	assert.Equal(t, validUserId, int(act.UserId))
	assert.Equal(t, validFriendId, int(act.FriendId))
}

func TestNewFriendWithSameUserIdReturnsError(t *testing.T) {
	_, err := NewFriend(validUserId, validUserId)

	assert.Error(t, err)
}

func TestDeleteFriendReturnsNotNil(t *testing.T) {
	sut, _ := NewFriend(validUserId, validFriendId)
	act := sut.DeleteFriend()
	assert.NotNil(t, act)
}

func TestInsertFriendReturnsNotNil(t *testing.T) {
	sut, _ := NewFriend(validUserId, validFriendId)
	act := sut.InsertFriend()
	assert.NotNil(t, act)
}

func TestAddFriendRequestCreateRequestReturnsNotNil(t *testing.T) {
	sut := AddFriendRequest{}
	act, _ := sut.CreateRequest(host)
	assert.NotNil(t, act)
}


func TestRemoveFriendRequestCreateRequestReturnsNotNil(t *testing.T) {
	sut := RemoveFriendRequest{}
	act, _ := sut.CreateRequest(host)
	assert.NotNil(t, act)
}
