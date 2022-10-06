package command

import "social-network/common"

type LogoutCommand struct {
}

type LogoutHandler = common.Handler[LogoutCommand, interface{}]

type logoutHandler struct {
}

func NewLogoutHandler() LogoutHandler {
	return &logoutHandler{}
}

func (h *logoutHandler) Handle(arg LogoutCommand) (interface{}, error) {
	return nil, nil
}
