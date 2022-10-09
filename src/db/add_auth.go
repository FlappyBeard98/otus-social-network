package db

type AddAuthQuery struct {
	Login    string
	Password string
}

func (receiver *AddAuthQuery) Sql() string {
	return `
INSERT INTO social_network.auth(login, password)
VALUES (?, ?)
;
`
}
