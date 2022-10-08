package db

type GetFriendsPageByUserIdQuery struct {
	Limit  int32
	Offset int32
	UserId int64
}

func (receiver *GetFriendsPageByUserIdQuery) Sql() string {
	return `
;WITH friends_page AS (
    SELECT
        right_user_id AS user_id
    FROM social_network.friends
    WHERE 
        left_user_id = $3
    LIMIT $1
    OFFSET $2
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
