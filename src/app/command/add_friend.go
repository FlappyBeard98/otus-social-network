package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type AddFriendCommand struct {
}

type AddFriendHandler = application.Handler[AddFriendCommand, interface{}]

type addFriendHandler struct {
	db *sql.DB
}

func NewAddFriendHandler(db *sql.DB) AddFriendHandler {
	return &addFriendHandler{db}
}

func (receiver *addFriendHandler) Handle(ctx context.Context, arg AddFriendCommand) (interface{}, error) {
	return nil, nil
}
