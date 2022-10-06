package command

import (
	"social-network/common"
)

type AddFriendCommand struct {
}

type AddFriendHandler = common.Handler[AddFriendCommand, interface{}]

type addFriendHandler struct {
}

func NewAddFriendHandler() AddFriendHandler {
	return &addFriendHandler{}
}

func (h *addFriendHandler) Handle(arg AddFriendCommand) (interface{}, error) {
	return nil, nil
}
