package db

type GetFriendsCountByUserIdQuery struct {
	UserId int64
}

func (receiver *GetFriendsCountByUserIdQuery) Sql() string {
	return `
SELECT
	COUNT(*) AS Total
FROM social_network.friends
WHERE
	user_id = ?
`
}
