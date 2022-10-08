package app

import (
	"database/sql"
	"social-network/app/command"
	"social-network/app/query"
)

type commands struct {
	Register     command.RegisterHandler
	Login        command.LoginHandler
	SaveProfile  command.SaveProfileHandler
	AddFriend    command.AddFriendHandler
	RemoveFriend command.RemoveFriendHandler
	Logout       command.LogoutHandler
}

type queries struct {
	Hello            query.HelloHandler //TODO remove
	ProfilesByFilter query.ProfilesByFilterHandler
	Profile          query.ProfileHandler
	Friends          query.FriendsHandler
}

type App struct {
	Commands commands
	Queries  queries
}

func NewApp(db *sql.DB) *App {
	return &App{
		Commands: commands{
			Register:     command.NewRegisterHandler(db),
			Login:        command.NewLoginHandler(db),
			SaveProfile:  command.NewSaveProfileHandler(db),
			AddFriend:    command.NewAddFriendHandler(db),
			RemoveFriend: command.NewRemoveFriendHandler(db),
			Logout:       command.NewLogoutHandler(db),
		},
		Queries: queries{
			Hello:            query.NewHelloHandler(),
			ProfilesByFilter: query.NewProfilesByFilterHandler(db),
			Profile:          query.NewProfileHandler(db),
			Friends:          query.NewFriendsHandler(db),
		},
	}
}
