package db

import (
	"fmt"
	"social-network/common/database"
)

type GetProfilesCountByFilterQuery struct {
	FirstName *string
	LastName  *string
	Age       *int32
	Gender    *int32
	City      *string
	Hobbies   *string
}

func (receiver *GetProfilesCountByFilterQuery) Sql() string {
	return fmt.Sprintf(`
SELECT
COUNT(*) AS Total
FROM social_network.profiles
WHERE
(? IS NULL OR first_name LIKE ?)
AND (? IS NULL OR last_name LIKE ?)
AND (? IS NULL OR age = ?)
AND (? IS NULL OR gender = ?)
AND (? IS NULL OR city = ?)
AND (? IS NULL OR hobbies LIKE ?)
;
`)
}

func (receiver *GetProfilesCountByFilterQuery) GetParams() []any {
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

	return params
}
