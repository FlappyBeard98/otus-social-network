package db

type UpdatePasswordQuery struct {
	UserId   int64
	Password string
}

func (receiver *UpdatePasswordQuery) Sql() string {
	return `
UPDATE social_network.auth SET
    password = $2
WHERE
    user_id = $1
;
`
}
