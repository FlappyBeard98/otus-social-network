package query

import (
	"social-network/common"
)

type FriendsQuery struct {
}

type FriendsHandler = common.Handler[FriendsQuery, interface{}]

type friendsHandler struct {
}

func NewFriendsHandler() FriendsHandler {
	return &friendsHandler{}
}

func (h *friendsHandler) Handle(arg FriendsQuery) (interface{}, error) {
	return nil, nil
}
