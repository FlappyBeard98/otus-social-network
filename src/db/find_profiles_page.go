package db

type FindProfilesPageQuery struct {
	Limit     int32
	Offset    int32
	FirstName *string
	LastName  *string
	Age       *int32
	Gender    *int32
	City      *string
	Hobbies   *string
}

func (receiver *FindProfilesPageQuery) Sql() string {
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
	($3 IS NULL OR first_name LIKE $3+'%') 
  	AND ($4 IS NULL OR last_name LIKE $4+'%') 
	AND ($5 IS NULL OR age = '$5') 
	AND ($6 IS NULL OR gender = $6) 
	AND ($7 IS NULL OR city LIKE $7+'%') 
	AND ($8 IS NULL OR hobbies LIKE '%'+$8+'%') 
LIMIT $1
OFFSET $2
;
`
}
