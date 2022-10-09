package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type RemoveFriendCommand struct {
}

type RemoveFriendHandler = application.Handler[RemoveFriendCommand, interface{}]

type removeFriendHandler struct {
	db *sql.DB
}

func NewRemoveFriendHandler(db *sql.DB) RemoveFriendHandler {
	return &removeFriendHandler{db}
}

func (receiver *removeFriendHandler) Handle(ctx context.Context, arg RemoveFriendCommand) (interface{}, error) {
	return nil, nil
}
