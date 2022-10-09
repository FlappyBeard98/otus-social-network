package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type LogoutCommand struct {
}

type LogoutHandler = application.Handler[LogoutCommand, interface{}]

type logoutHandler struct {
	db *sql.DB
}

func NewLogoutHandler(db *sql.DB) LogoutHandler {
	return &logoutHandler{db}
}

func (receiver *logoutHandler) Handle(ctx context.Context, arg LogoutCommand) (interface{}, error) {
	return nil, nil
}
