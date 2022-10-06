package command

import "social-network/common"

type RemoveFriendCommand struct {
}

type RemoveFriendHandler = common.Handler[RemoveFriendCommand, interface{}]

type removeFriendHandler struct {
}

func NewRemoveFriendHandler() RemoveFriendHandler {
	return &removeFriendHandler{}
}

func (h *removeFriendHandler) Handle(arg RemoveFriendCommand) (interface{}, error) {
	return nil, nil
}
