package db

type GetFriendsPageByUserIdQuery struct {
	UserId int64
	Limit  int64
	Offset int64
}

func (receiver *GetFriendsPageByUserIdQuery) Sql() string {
	return `
;WITH friends_page AS (
    SELECT
        friend_user_id AS user_id
    FROM social_network.friends
    WHERE 
        user_id = ?
    LIMIT ?
    OFFSET ?
)
SELECT
     user_id
    ,first_name
    ,last_name
    ,age
    ,gender
    ,city
    ,hobbies
FROM friends_page f 
JOIN social_network.profile p ON p.user_id = f.user_id
;
`
}
