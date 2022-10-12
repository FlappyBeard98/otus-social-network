package db

import (
	"github.com/georgysavva/scany/v2/sqlscan"
	"social-network/common/database"
)


type Repository struct {
	AddAuth database.DbHandler[*AddAuthQuery,any]
	GetAuthByLogin database.DbHandler[*GetAuthByLoginQuery,Auth]
	AddFriend database.DbHandler[*AddFriendQuery,any]
	GetFriendsCountByUserId database.DbHandler[*GetFriendsCountByUserIdQuery,int32]
	GetFriendsPageByUserId database.DbHandler[*GetFriendsPageByUserIdQuery,Profile]
	GetProfileByUserId database.DbHandler[*GetProfileByUserIdQuery,Profile]
	GetProfilesCountByFilter database.DbHandler[*GetProfilesCountByFilterQuery,int32]
	GetProfilesPageByFilter database.DbHandler[*GetProfilesPageByFilterQuery,Profile]
	RemoveFriend database.DbHandler[*RemoveFriendQuery,any]
	SavePassword database.DbHandler[*SavePasswordQuery,any]
	SaveProfile database.DbHandler[*SaveProfileQuery,any]
}



func NewRepository(connection sqlscan.Querier) *Repository {
	return &Repository{
		AddAuth:                  database.NewDbHandler[*AddAuthQuery,any](connection),
		GetAuthByLogin:			  database.NewDbHandler[*GetAuthByLoginQuery,Auth](connection),
		AddFriend:                database.NewDbHandler[*AddFriendQuery,any](connection),
		GetFriendsCountByUserId:  database.NewDbHandler[*GetFriendsCountByUserIdQuery,int32](connection),
		GetFriendsPageByUserId:   database.NewDbHandler[*GetFriendsPageByUserIdQuery,Profile](connection),
		GetProfileByUserId:       database.NewDbHandler[*GetProfileByUserIdQuery,Profile](connection),
		GetProfilesCountByFilter: database.NewDbHandler[*GetProfilesCountByFilterQuery,int32](connection),
		GetProfilesPageByFilter:  database.NewDbHandler[*GetProfilesPageByFilterQuery,Profile](connection),
		RemoveFriend:             database.NewDbHandler[*RemoveFriendQuery,any](connection),
		SavePassword:             database.NewDbHandler[*SavePasswordQuery,any](connection),
		SaveProfile:              database.NewDbHandler[*SaveProfileQuery,any](connection),
	}
}






