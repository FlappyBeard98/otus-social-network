package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	userId   = 1
	friendId = 2
)

func TestNewFriendReturnsValidFriend(t *testing.T) {
	act, _ := NewFriend(userId, friendId)

	assert.Equal(t, userId, int(act.UserId))
	assert.Equal(t, friendId, int(act.FriendId))
}

func TestNewFriendWithSameUserIdReturnsError(t *testing.T) {
	_, err := NewFriend(userId, userId)

	assert.Error(t, ErrInvalidInput, err)
}

func TestDeleteFriendReturnsNotNil(t *testing.T) {
	sut, _ := NewFriend(userId, friendId)
	act := sut.DeleteFriend()
	assert.NotNil(t, act)
}

func TestInsertFriendReturnsNotNil(t *testing.T) {
	sut, _ := NewFriend(userId, friendId)
	act := sut.InsertFriend()
	assert.NotNil(t, act)
}
