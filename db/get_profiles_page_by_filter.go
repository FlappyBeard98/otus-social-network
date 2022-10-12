package db

import (
	"fmt"
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
	return fmt.Sprintf(`
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
	%s AND %s AND %s AND %s AND %s AND %s
LIMIT ?
OFFSET ?
;
`,
database.NilOrExprMysql(receiver.FirstName,"first_name LIKE ? +'%'"),
database.NilOrExprMysql(receiver.LastName,"last_name LIKE ? +'%'"),
database.NilOrExprMysql(receiver.Age,"age = ?"),
database.NilOrExprMysql(receiver.Gender,"gender = ?"),
database.NilOrExprMysql(receiver.City,"city = ?"),
database.NilOrExprMysql(receiver.Hobbies,"hobbies like '%' + ? + '%'"),
)
}
