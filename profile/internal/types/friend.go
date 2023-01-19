package types

import (
	"errors"

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
		INSERT INTO social_network.friends(user_id, friend_user_id)
		VALUES (?, ?)
		;`,
		o.UserId,
		o.FriendId)
}

// DeleteFriend returns new mysql.SqlQuery for deleting friend link from database
func (o *Friend) DeleteFriend() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		DELETE FROM social_network.friends
		WHERE
			user_id = ?
			AND friend_user_id = ?
		;`,
		o.UserId,
		o.FriendId)
}
