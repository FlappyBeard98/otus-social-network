package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type LoginCommand struct {
}

type LoginHandler = application.Handler[LoginCommand, interface{}]

type loginHandler struct {
	db *sql.DB
}

func NewLoginHandler(db *sql.DB) LoginHandler {
	return &loginHandler{db}
}

func (receiver *loginHandler) Handle(ctx context.Context, arg LoginCommand) (interface{}, error) {
	return nil, nil
}
