package db

type SaveTokenQuery struct {
	UserId int64
	Token  string
}

func (receiver *SaveTokenQuery) Sql() string {
	return `
UPDATE social_network.auth SET
    token = $2
WHERE
    user_id = $1
;
`
}
