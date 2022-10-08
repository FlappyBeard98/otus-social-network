package db

type FindProfileByUserIdQuery struct {
	UserId int64
}

func (receiver *FindProfileByUserIdQuery) Sql() string {
	return `
SELECT 
	 user_id
    ,first_name
    ,last_name
    ,age
    ,gender
    ,city
    ,hobbies
FROM social_network.profile
WHERE 
	user_id = $1
;
`
}