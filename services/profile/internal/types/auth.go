package types

import (
	"errors"

	"social-network/lib/pg"
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

// InsertAuth returns new pg.SqlQuery for inserting user auth in database
func (o *Auth) InsertAuth() *pg.SqlQuery {
	return pg.NewSqlQuery(`
		INSERT INTO profiles.auth(login, password)
		VALUES ($1, $2)
		RETURNING user_id;`,
		o.Login,
		[]byte(o.Password))
}

// UpdatePassword returns new pg.SqlQuery for updating password in database
func (o *Auth) UpdatePassword() *pg.SqlQuery {
	return pg.NewSqlQuery(`
		UPDATE profiles.auth SET
			password = $1
		WHERE
			user_id = $2;`,
		o.Password,
		o.UserId)
}

// ReadByLogin returns new pg.SqlQuery that returns user auth from database
func (o *Auth) ReadByLogin() *pg.SqlQuery {
	return pg.NewSqlQuery(`
		SELECT
			user_id
			,login
			,password
		FROM profiles.auth
		WHERE
			login = $1;`,
		o.Login)
}
