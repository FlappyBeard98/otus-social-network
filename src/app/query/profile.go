package query

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type ProfileQuery struct {
}

type ProfileHandler = application.Handler[ProfileQuery, interface{}]

type profileHandler struct {
	db *sql.DB
}

func NewProfileHandler(db *sql.DB) ProfileHandler {
	return &profileHandler{db}
}

func (h *profileHandler) Handle(ctx context.Context, arg ProfileQuery) (interface{}, error) {
	return nil, nil
}
