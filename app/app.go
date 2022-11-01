package app

import (
	"database/sql"
	"social-network/app/command"
	"social-network/app/query"
)

type commands struct {
	Register     command.RegisterHandler
	SaveProfile  command.SaveProfileHandler
	AddFriend    command.AddFriendHandler
	RemoveFriend command.RemoveFriendHandler
	ClearUsers   command.ClearUsersHandler
}

type queries struct {
	IsValidAuth      query.IsValidAuthHandler
	ProfilesByFilter query.ProfilesByFilterHandler
	Profile          query.ProfileHandler
	Friends          query.FriendsHandler
}

type App struct {
	Commands commands
	Queries  queries
}

func NewApp(db *sql.DB, key string) *App {
	return &App{
		Commands: commands{
			Register:     command.NewRegisterHandler(db, key),
			SaveProfile:  command.NewSaveProfileHandler(db),
			AddFriend:    command.NewAddFriendHandler(db),
			RemoveFriend: command.NewRemoveFriendHandler(db),
			ClearUsers:   command.NewClearUsersHandler(db),
		},
		Queries: queries{
			IsValidAuth:      query.NewIsValidAuthHandler(db, key),
			ProfilesByFilter: query.NewProfilesByFilterHandler(db),
			Profile:          query.NewProfileHandler(db),
			Friends:          query.NewFriendsHandler(db),
		},
	}
}
