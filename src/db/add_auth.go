package db

type AddAuthQuery struct {
	Login    string
	Password string
}

func (receiver *AddAuthQuery) Sql() string {
	return `
INSERT INTO social_network.auth(login, password)
VALUES ($1, $2)
;
SELECT 
     user_id
    ,login
    ,password
    ,token 
FROM social_network.auth
WHERE 
    login = $1
;
`
}
