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
WHERE
	%s AND %s AND %s AND %s AND %s AND %s
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


