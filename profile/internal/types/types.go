package types

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInternal         = errors.New("internal error")
)

type RegisterRequest struct{
	Auth Auth `json:"auth"`
	Profile Profile `json:"profile"`
}
//TODO: move RegisterRequest to separate file and move mysql.SqlQuery's to it