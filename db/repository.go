package db

import (
	"github.com/georgysavva/scany/v2/sqlscan"
	"social-network/common/database"
)

type Repository struct {
	AddAuth                  database.DbHandler[*AddAuthQuery, any]
	GetAuthByLogin           database.DbHandler[*GetAuthByLoginQuery, Auth]
	AddFriend                database.DbHandler[*AddFriendQuery, any]
	GetFriendsCountByUserId  database.DbHandler[*GetFriendsCountByUserIdQuery, int32]
	GetFriendsPageByUserId   database.DbHandler[*GetFriendsPageByUserIdQuery, Profile]
	GetProfileByUserId       database.DbHandler[*GetProfileByUserIdQuery, Profile]
	GetProfilesCountByFilter database.DbHandler[*GetProfilesCountByFilterQuery, int32]
	GetProfilesPageByFilter  database.DbHandler[*GetProfilesPageByFilterQuery, Profile]
	RemoveFriend             database.DbHandler[*RemoveFriendQuery, any]
	SavePassword             database.DbHandler[*SavePasswordQuery, any]
	SaveProfile              database.DbHandler[*SaveProfileQuery, any]             
}

func NewRepository(connection sqlscan.Querier) *Repository {
	return &Repository{
		AddAuth:                  database.NewDbQuerierHandler[*AddAuthQuery, any](connection),
		GetAuthByLogin:           database.NewDbQuerierHandler[*GetAuthByLoginQuery, Auth](connection),
		AddFriend:                database.NewDbQuerierHandler[*AddFriendQuery, any](connection),
		GetFriendsCountByUserId:  database.NewDbQuerierHandler[*GetFriendsCountByUserIdQuery, int32](connection),
		GetFriendsPageByUserId:   database.NewDbQuerierHandler[*GetFriendsPageByUserIdQuery, Profile](connection),
		GetProfileByUserId:       database.NewDbQuerierHandler[*GetProfileByUserIdQuery, Profile](connection),
		GetProfilesCountByFilter: database.NewDbQuerierHandler[*GetProfilesCountByFilterQuery, int32](connection),
		GetProfilesPageByFilter:  database.NewDbQuerierHandler[*GetProfilesPageByFilterQuery, Profile](connection),
		RemoveFriend:             database.NewDbQuerierHandler[*RemoveFriendQuery, any](connection),
		SavePassword:             database.NewDbQuerierHandler[*SavePasswordQuery, any](connection),
		SaveProfile:              database.NewDbQuerierHandler[*SaveProfileQuery, any](connection),
	}
}
