package db

type GetAuthByLoginQuery struct {
	Login    string
}

func (receiver *GetAuthByLoginQuery) Sql() string {
	return `
SELECT
     user_id
    ,login
    ,password
    ,token
FROM social_network.auth
WHERE
    login = ?
;
`
}
