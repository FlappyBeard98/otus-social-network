package query

import (
	"context"
	"database/sql"
	"social-network/common"
	"social-network/common/application"
	"social-network/db"
)

type IsValidAuthQuery struct {
	Login string
	Password string
}

type IsValidAuthHandler = application.Handler[IsValidAuthQuery, bool]

type isValidAuthQueryHandler struct {
	db *sql.DB
	key string
}

func NewIsValidAuthHandler(db *sql.DB,key string) IsValidAuthHandler {
	return &isValidAuthQueryHandler{db,key}
}

func (receiver *isValidAuthQueryHandler) Handle(ctx context.Context, arg IsValidAuthQuery) (bool, error) {

	r := db.NewRepository(receiver.db)

	auth, err := r.GetAuthByLogin.Handle(ctx,&db.GetAuthByLoginQuery{Login: arg.Login})

	if err != nil {
		return false, err
	}

	password, err := common.Decrypt([]byte(receiver.key), []byte(auth[0].Password) )

	if err != nil {
		return false, err
	}

	if string(password) != arg.Password {
		return false, nil
	}

	return true, nil
}
