package types

import (
	"social-network/lib/pg"
)

// ProfilesRequest used for return user profiles data
type ProfilesRequest struct {
	FirstName *string `query:"firstName"` //user first name
	LastName  *string `query:"lastName"`  //user last name
	Age       *int32  `query:"age"`       //user age
	Gender    *int32  `query:"gender"`    //user gender
	City      *string `query:"city"`      //user city
	Hobbies   *string `query:"hobbies"`   //user hobbies
	PageRequest
}

// ReadProfilesTotal returns new pg.SqlQuery for selecting total count of user profiles by filter in ProfilesRequest
func (o *ProfilesRequest) ReadProfilesTotal() *pg.SqlQuery {
	params := []any{
		o.FirstName,
		pg.Like(o.FirstName, false, true),
		o.LastName,
		pg.Like(o.LastName, false, true),
		o.Age,
		o.Age,
		o.Gender,
		o.Gender,
		o.City,
		o.City,
		o.Hobbies,
		pg.Like(o.Hobbies, true, true),
	}
	return pg.NewSqlQuery(`
		SELECT
			COUNT(*) AS Total
		FROM profiles.profiles
		WHERE
				(cast($1 as varchar(100)) IS NULL OR first_name LIKE $2)
			AND (cast($3 as varchar(150)) IS NULL OR last_name LIKE $4)
			AND (cast($5 as int) IS NULL OR age = $6)
			AND (cast($7 as int) IS NULL OR gender = $8)
			AND (cast($9 as varchar(50)) IS NULL OR city = $10)
			AND (cast($11 as text) IS NULL OR hobbies LIKE $12)
		;`,
		params...)
}

// ReadProfilesPage returns new pg.SqlQuery for selecting page of user profiles by filter in ProfilesRequest
func (o *ProfilesRequest) ReadProfilesPage() *pg.SqlQuery {
	params := []any{
		o.FirstName,
		pg.Like(o.FirstName, false, true),
		o.LastName,
		pg.Like(o.LastName, false, true),
		o.Age,
		o.Age,
		o.Gender,
		o.Gender,
		o.City,
		o.City,
		o.Hobbies,
		pg.Like(o.Hobbies, true, true),
		o.Limit,
		o.Offset,
	}

	return pg.NewSqlQuery(`
		SELECT
			user_id
			,first_name
			,last_name
			,age
			,gender
			,city
			,hobbies
		FROM profiles.profiles
		WHERE
				(cast($1 as varchar(100)) IS NULL OR first_name LIKE $2)
			AND (cast($3 as varchar(150)) IS NULL OR last_name LIKE $4)
			AND (cast($5 as int) IS NULL OR age = $6)
			AND (cast($7 as int) IS NULL OR gender = $8)
			AND (cast($9 as varchar(50)) IS NULL OR city = $10)
			AND (cast($11 as text) IS NULL OR hobbies LIKE $12)
		LIMIT $13
		OFFSET $14;`,
		params...)
}
