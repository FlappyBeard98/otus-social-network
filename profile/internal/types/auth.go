package types

import (
	"errors"

	"social-network/lib/mysql"
	"social-network/lib/utils"
)

// Auth define user authorization data
type Auth struct {
	UserId   int64  `json:"userId" db:"user_id"`    //user identifier
	Login    string `json:"login" db:"login"`       //user login
	Password string `json:"password" db:"password"` // hashed password
}

// NewAuth creates new valid Auth with hashed password or return error if validation failed
func NewAuth(login string, rawPassword string) (*Auth, error) {
	if len(login) < 8 || len(login) > 250 {
		return nil, errors.New("%w: login value must be between 8 and 250 characters")
	}

	if len(rawPassword) < 8 {
		return nil, errors.New("%w: password must be 8 characters long or more")
	}

	passwordHash := utils.GetHash(rawPassword)

	return &Auth{
		Login:    login,
		Password: passwordHash,
	}, nil
}

// InsertAuth returns new mysql.SqlQuery for inserting user auth in database
func (o *Auth) InsertAuth() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		INSERT INTO social_network.auth(login, password)
		VALUES (?, ?);`,
		o.Login,
		o.Password)
}

// UpdatePassword returns new mysql.SqlQuery for updating password in database
func (o *Auth) UpdatePassword() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		UPDATE social_network.auth SET
			password = ?
		WHERE
			user_id = ?;`,
		o.Password,
		o.UserId)
}

// ReadByLogin returns new mysql.SqlQuery that returns user auth from database
func (o *Auth) ReadByLogin() *mysql.SqlQuery {
	return mysql.NewSqlQuery(`
		SELECT
			user_id
			,login
			,password
		FROM social_network.auth
		WHERE
			login = ?;`,
		o.Login)
}
