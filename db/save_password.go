package db

type SavePasswordQuery struct {
	Password string
	UserId   int64
}

func (receiver *SavePasswordQuery) Sql() string {
	return `
UPDATE social_network.auth SET
    password = ?
WHERE
    user_id = ?
;
`
}
