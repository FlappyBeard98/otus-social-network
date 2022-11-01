package db

import (
	"social-network/common/database"
)

type GetProfilesPageByFilterQuery struct {
	FirstName *string
	LastName  *string
	Age       *int32
	Gender    *int32
	City      *string
	Hobbies   *string
	Limit     int64
	Offset    int64
}

func (receiver *GetProfilesPageByFilterQuery) Sql() string {
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
		(? IS NULL OR first_name LIKE ?)
	AND (? IS NULL OR last_name LIKE ?)
	AND (? IS NULL OR age = ?)
	AND (? IS NULL OR gender = ?)
	AND (? IS NULL OR city = ?)
	AND (? IS NULL OR hobbies LIKE ?)
LIMIT ?
OFFSET ?
ORDER BY user_id
;`
}

func (receiver *GetProfilesPageByFilterQuery) GetParams() []any {
	params := make([]any, 0)
	params = append(params, receiver.FirstName)
	params = append(params, database.Like(receiver.FirstName, false, true))
	params = append(params, receiver.LastName)
	params = append(params, database.Like(receiver.LastName, false, true))
	params = append(params, receiver.Age)
	params = append(params, receiver.Age)
	params = append(params, receiver.Gender)
	params = append(params, receiver.Gender)
	params = append(params, receiver.City)
	params = append(params, receiver.City)
	params = append(params, receiver.Hobbies)
	params = append(params, database.Like(receiver.Hobbies, true, true))
	params = append(params, receiver.Limit)
	params = append(params, receiver.Offset)
	return params
}
