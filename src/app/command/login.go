package command

import (
	"social-network/common"
)

type LoginCommand struct {
}

type LoginHandler = common.Handler[LoginCommand, interface{}]

type loginHandler struct {
}

func NewLoginHandler() LoginHandler {
	return &loginHandler{}
}

func (h *loginHandler) Handle(arg LoginCommand) (interface{}, error) {
	return nil, nil
}
