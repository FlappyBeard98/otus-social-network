package db

type ClearUsersQuery struct {

}

func (receiver *ClearUsersQuery) Sql() string {
	return `
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE social_network.friends;
TRUNCATE TABLE social_network.profiles;
TRUNCATE TABLE social_network.auth;
SET FOREIGN_KEY_CHECKS = 1;
`
}


