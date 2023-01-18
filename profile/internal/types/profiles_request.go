package types

import (
	"social-network/lib/mysql"
)

// ProfilesRequest
type ProfilesRequest struct {
	FirstName *string `query:"firstName"` //user first name
	LastName  *string `query:"lastName"`  //user last name
	Age       *int32  `query:"age"`       //user age
	Gender    *int32  `query:"gender"`    //user gender
	City      *string `query:"city"`      //user city
	Hobbies   *string `query:"hobbies"`   //user hobbies
}

// ReadProfilesTotal returns new mysql.SqlQuery for selecting total count of user profiles by filter in ProfilesRequest
func (o *ProfilesRequest) ReadProfilesTotal() *mysql.SqlQuery {
	params := []any{
		o.FirstName,
		mysql.Like(o.FirstName, false, true),
		o.LastName,
		mysql.Like(o.LastName, false, true),
		o.Age,
		o.Age,
		o.Gender,
		o.Gender,
		o.City,
		o.City,
		o.Hobbies,
		mysql.Like(o.Hobbies, true, true),
	}
	return mysql.NewSqlQuery(`
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
		;`,
		params...)
}

// ReadProfilesPage returns new mysql.SqlQuery for selecting page of user profiles by filter in ProfilesRequest
func (o *ProfilesRequest) ReadProfilesPage(limit int, offset int) *mysql.SqlQuery {
	params := []any{
		o.FirstName,
		mysql.Like(o.FirstName, false, true),
		o.LastName,
		mysql.Like(o.LastName, false, true),
		o.Age,
		o.Age,
		o.Gender,
		o.Gender,
		o.City,
		o.City,
		o.Hobbies,
		mysql.Like(o.Hobbies, true, true),
		limit,
		offset,
	}

	return mysql.NewSqlQuery(`
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
		OFFSET ?;`,
		params...)
}
