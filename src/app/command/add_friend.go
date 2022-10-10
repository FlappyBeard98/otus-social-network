package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
	"social-network/db"
)

type AddFriendCommand struct {
	UserId int64 `param:"userId"`
	FriendUserId int64 `json:"userId"`
}

type AddFriendHandler = application.Handler[AddFriendCommand, interface{}]

type addFriendHandler struct {
	db *sql.DB
}

func NewAddFriendHandler(db *sql.DB) AddFriendHandler {
	return &addFriendHandler{db}
}

func (receiver *addFriendHandler) Handle(ctx context.Context, arg AddFriendCommand) (interface{}, error) {


	r := db.NewRepository(receiver.db)

	_,err := r.RemoveFriend.Handle(ctx,&db.RemoveFriendQuery{
		UserId:       arg.UserId,
		FriendUserId: arg.FriendUserId,
		})

	if err!=nil {
		return nil, err
	}

	return nil, nil

}
