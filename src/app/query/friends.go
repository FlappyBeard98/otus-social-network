package query

import (
	"context"
	"database/sql"
	"social-network/common/application"
)

type FriendsQuery struct {
}

type FriendsHandler = application.Handler[FriendsQuery, interface{}]

type friendsHandler struct {
	db *sql.DB
}

func NewFriendsHandler(db *sql.DB) FriendsHandler {
	return &friendsHandler{db}
}

func (h *friendsHandler) Handle(ctx context.Context, arg FriendsQuery) (interface{}, error) {
	return nil, nil
}
