package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type RegisterCommand struct {
}

type RegisterHandler = application.Handler[RegisterCommand, interface{}]

type registerHandler struct {
	db *sql.DB
}

func NewRegisterHandler(db *sql.DB) RegisterHandler {
	return &registerHandler{db}
}

func (h *registerHandler) Handle(ctx context.Context, arg RegisterCommand) (interface{}, error) {
	return nil, nil
}
