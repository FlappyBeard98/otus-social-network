package types

import (
	"fmt"
	"net/http"
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
		FROM profiles.profiles
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
		FROM profiles.friends
		WHERE
			user_id = ?
		;`,
		o.UserId)
}

func (o *UserIdRequest) ReadUserFriendsProfiles(page *PageRequest) *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		WITH friends_page AS (
			SELECT
				friend_user_id AS user_id
			FROM profiles.friends
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
		JOIN profiles.profiles p ON p.user_id = f.user_id
		;`,
		o.UserId,
		page.Limit,
		page.Offset)
}


type GetProfileRequest  UserIdRequest 

func (o *GetProfileRequest) CreateRequest(host string) (*http.Request, error) {
	route := fmt.Sprintf("%s/%d",host, o.UserId)

	request, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, nil
}

type GetFriendsRequest struct {
	UserIdRequest
	PageRequest
}

func (o *GetFriendsRequest) CreateRequest(host string) (*http.Request, error) {
	route := fmt.Sprintf("%s/%d/friends?limit=%d&offset=%d",host, o.UserId, o.Limit, o.Offset)

	request, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, nil
}


