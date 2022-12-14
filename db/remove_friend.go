package db

type RemoveFriendQuery struct {
	UserId       int64
	FriendUserId int64
}

func (receiver *RemoveFriendQuery) Sql() string {
	return `
DELETE FROM social_network.friends
WHERE
    user_id = ?
    AND friend_user_id = ?
;
`
}
