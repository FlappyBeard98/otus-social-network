package types

import (
	"social-network/lib/mysql"
)

// UserIdRequest used for http requests with userId
type UserIdRequest struct {
	UserId int64 `param:"userId"` //user identifier
}

// ReadProfilesTotal returns new mysql.SqlQuery for selecting total count of user profiles by filter in ProfilesRequest
func (o *UserIdRequest) ReadProfileByUserId() *mysql.SqlQuery {
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
			user_id = ?
		;`,
		o.UserId)
}

// ReadProfilesPage returns new mysql.SqlQuery for selecting page of user profiles by filter in ProfilesRequest
func (o *UserIdRequest) ReadUserFriendsTotal() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		SELECT
			COUNT(*) AS Total
		FROM social_network.friends
		WHERE
			user_id = ?
		;`,
		o.UserId)
}

func (o *UserIdRequest) ReadUserFriendsProfiles(limit int, offset int) *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		WITH friends_page AS (
			SELECT
				friend_user_id AS user_id
			FROM social_network.friends
			WHERE 
				user_id = ?
			LIMIT ?
			OFFSET ?
		)
		SELECT
			p.user_id
			,first_name
			,last_name
			,age
			,gender
			,city
			,hobbies
		FROM friends_page f 
		JOIN social_network.profiles p ON p.user_id = f.user_id
		;`,
		o.UserId,
		limit,
		offset)
}
