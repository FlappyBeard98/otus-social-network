package db

type AddFriendQuery struct {
	UserId       int64
	FriendUserId int64
}

func (receiver *AddFriendQuery) Sql() string {
	return `
INSERT INTO social_network.friends(user_id, friend_user_id)
VALUES (?, ?)
;
`
}
