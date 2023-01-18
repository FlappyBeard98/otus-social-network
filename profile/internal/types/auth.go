package types

import (
	"fmt"
	"social-network/lib/mysql"
	"social-network/lib/utils"
)

type Auth struct {
	UserId   int64
	Login    string
	Password string
}

func NewAuth(login string, rawPassword string) (*Auth, error) {
	if len(login) < 8 || len(login) > 250 {
		return nil, fmt.Errorf("%w: login value must be between 8 and 250 characters", ErrInvalidInput)
	}

	if len(rawPassword) < 8 {
		return nil, fmt.Errorf("%w: password must be 8 characters long or more", ErrInvalidInput)
	}

	return &Auth{
		Login:    login,
		Password: utils.GetHash(rawPassword),
	}, nil
}

func (o *Auth) PasswordEquals(password string) bool {
	ph := utils.GetHash(password)
	return ph == o.Password
}

func (o *Auth) InsertAuth() *mysql.SqlQuery {
	return mysql.NewSqlQuery(
		`
		INSERT INTO social_network.auth(login, password)
		VALUES (?, ?);`,
		o.Login,
		o.Password)
}

func (o *Auth) UpdatePassword() *mysql.SqlQuery {
	return mysql.NewSqlQuery(
		`
		UPDATE social_network.auth SET
			password = ?
		WHERE
			user_id = ?;`,
		o.Password,
		o.UserId)
}
