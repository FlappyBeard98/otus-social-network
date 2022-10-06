package query

import "social-network/common"

type ProfileQuery struct {
}

type ProfileHandler = common.Handler[ProfileQuery, interface{}]

type profileHandler struct {
}

func NewProfileHandler() ProfileHandler {
	return &profileHandler{}
}

func (h *profileHandler) Handle(arg ProfileQuery) (interface{}, error) {
	return nil, nil
}
