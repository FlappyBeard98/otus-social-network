package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type SaveProfileCommand struct {
}

type SaveProfileHandler = application.Handler[SaveProfileCommand, interface{}]

type saveProfileHandler struct {
	db *sql.DB
}

func NewSaveProfileHandler(db *sql.DB) SaveProfileHandler {
	return &saveProfileHandler{db}
}

func (h *saveProfileHandler) Handle(ctx context.Context, arg SaveProfileCommand) (interface{}, error) {
	return nil, nil
}
