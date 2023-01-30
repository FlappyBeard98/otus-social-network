package types

import (
	"errors"
	"fmt"
	"net/http"

	"social-network/lib/mysql"
)

// Friend stores data about friendship
type Friend struct {
	UserId   int64 `db:"user_id" param:"userId"`              //user identifier
	FriendId int64 `db:"friend_user_id" param:"friendUserId"` //user friend identifier
}

// NewFriend creates new valid Friend link or return error if validation failed
func NewFriend(userId int64, friendId int64) (*Friend, error) {

	if userId == friendId {
		return nil, errors.New("%w: user can not create frield link with himself")
	}

	return &Friend{
		UserId:   userId,
		FriendId: friendId,
	}, nil
}

// InsertFriend returns new mysql.SqlQuery for inserting friend link in database
func (o *Friend) InsertFriend() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		INSERT INTO profiles.friends(user_id, friend_user_id)
		VALUES (?, ?)
		;`,
		o.UserId,
		o.FriendId)
}

// DeleteFriend returns new mysql.SqlQuery for deleting friend link from database
func (o *Friend) DeleteFriend() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		DELETE FROM profiles.friends
		WHERE
			user_id = ?
			AND friend_user_id = ?
		;`,
		o.UserId,
		o.FriendId)
}



type AddFriendRequest  Friend

func (o *AddFriendRequest) CreateRequest(host string) (*http.Request, error) {
	route := fmt.Sprintf("%s/%d/friends/%d",host, o.UserId,o.FriendId)

	request, err := http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, nil
}

type RemoveFriendRequest  Friend

func (o *RemoveFriendRequest) CreateRequest(host string) (*http.Request, error) {
	route := fmt.Sprintf("%s/%d/%d",host, o.UserId,o.FriendId)

	request, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, nil
}