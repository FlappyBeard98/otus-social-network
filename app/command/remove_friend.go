package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
	"social-network/db"
)

type RemoveFriendCommand struct {
	UserId       int64 `param:"userId"`
	FriendUserId int64 `param:"friendUserId"`
}

type RemoveFriendHandler = application.Handler[RemoveFriendCommand, interface{}]

type removeFriendHandler struct {
	db *sql.DB
}

func NewRemoveFriendHandler(db *sql.DB) RemoveFriendHandler {
	return &removeFriendHandler{db}
}

func (receiver *removeFriendHandler) Handle(ctx context.Context, arg RemoveFriendCommand) (interface{}, error) {

	r := db.NewRepository(receiver.db)

	_, err := r.RemoveFriend.Handle(ctx, &db.RemoveFriendQuery{
		UserId:       arg.UserId,
		FriendUserId: arg.FriendUserId,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
