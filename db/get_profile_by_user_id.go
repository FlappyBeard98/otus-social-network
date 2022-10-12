package db

type GetProfileByUserIdQuery struct {
	UserId int64
}

func (receiver *GetProfileByUserIdQuery) Sql() string {
	return `
SELECT
	 user_id
    ,first_name
    ,last_name
    ,age
    ,gender
    ,city
    ,hobbies
FROM social_network.profiles
WHERE
	user_id = ?
;
`
}
