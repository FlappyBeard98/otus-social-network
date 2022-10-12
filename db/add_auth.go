package db

type AddAuthQuery struct {
	Login    string
	Password []byte
}

func (receiver *AddAuthQuery) Sql() string {
	return `
INSERT INTO social_network.auth(login, password)
VALUES (?, ?)
;
`
}
