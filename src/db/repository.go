package db

import (
	"github.com/georgysavva/scany/v2/sqlscan"
	"social-network/common/database"
)


type Repository struct {
	AddAuth database.DbHandler[*AddAuthQuery,any]
	GetAuthByLogin database.DbHandler[*GetAuthByLoginQuery,Auth]
	AddFriend any
	GetFriendsCountByUserId any
	GetFriendsPageByUserId any
	GetProfileByUserId any
	GetProfilesCountByFilter any
	GetProfilesPageByFilter any
	RemoveFriend any
	SavePassword any
	SaveProfile database.DbHandler[*SaveProfileQuery,any]
	SaveToken any
}



func NewRepository(connection sqlscan.Querier) *Repository {
	return &Repository{
		AddAuth:                  database.NewDbHandler[*AddAuthQuery,any](connection),
		GetAuthByLogin:			  database.NewDbHandler[*GetAuthByLoginQuery,Auth](connection),
		AddFriend:                nil,
		GetFriendsCountByUserId:  nil,
		GetFriendsPageByUserId:   nil,
		GetProfileByUserId:       nil,
		GetProfilesCountByFilter: nil,
		GetProfilesPageByFilter:  nil,
		RemoveFriend:             nil,
		SavePassword:             nil,
		SaveProfile:              database.NewDbHandler[*SaveProfileQuery,any](connection),
		SaveToken:                nil,
	}
}






