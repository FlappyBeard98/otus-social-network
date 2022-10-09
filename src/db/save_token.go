package db

type SaveTokenQuery struct {
	Token  *string
	UserId int64
}

func (receiver *SaveTokenQuery) Sql() string {
	return `
UPDATE social_network.auth SET
    token = ?
WHERE
    user_id = ?
;
`
}
