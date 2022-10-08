package query

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type ProfilesByFilterQuery struct {
}

type ProfilesByFilterHandler = application.Handler[ProfilesByFilterQuery, interface{}]

type profilesByFilterHandler struct {
	db *sql.DB
}

func NewProfilesByFilterHandler(db *sql.DB) ProfilesByFilterHandler {
	return &profilesByFilterHandler{db}
}

func (h *profilesByFilterHandler) Handle(ctx context.Context, arg ProfilesByFilterQuery) (interface{}, error) {
	return nil, nil
}
