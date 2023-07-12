package types

import (
	"errors"
	"social-network/lib/pg"
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

// InsertFriend returns new pg.SqlQuery for inserting friend link in database
func (o *Friend) InsertFriend() *pg.SqlQuery {
	return pg.NewSqlQuery(`
		INSERT INTO profiles.friends(user_id, friend_user_id)
		VALUES ($1, $2)
		;`,
		o.UserId,
		o.FriendId)
}

// DeleteFriend returns new pg.SqlQuery for deleting friend link from database
func (o *Friend) DeleteFriend() *pg.SqlQuery {
	return pg.NewSqlQuery(`
		DELETE FROM profiles.friends
		WHERE
			user_id = $1
			AND friend_user_id = $2
		;`,
		o.UserId,
		o.FriendId)
}
