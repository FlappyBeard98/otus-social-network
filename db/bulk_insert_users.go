package db

import (
	"fmt"
	"strings"
)

type BulkInsertUsersQuery struct {
	Users []struct{
		Profile
		Auth
	}
}

func (receiver *BulkInsertUsersQuery) Sql() string {

	var (
		placeholders []string
		vals         []interface{}
	)

	for idx, user := range receiver.Users {
		placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d)",
		idx*3+1,
		idx*3+2,
		idx*3+3,
		))

		vals = append(vals, user.FirstName, user.LastName, user.Email)
	}
	insertStatement := fmt.Sprintf("INSERT INTO contacts(first_name,last_name,email) VALUES %s", strings.Join(placeholders, ","))
	
	
	return insertStatement
}


