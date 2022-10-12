package command

import (
	"context"
	"database/sql"
	"social-network/common"
	"social-network/common/application"
	"social-network/common/database"
	"social-network/db"
)

type RegisterCommand struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
	Gender    int32   `json:"gender"`
	City      string `json:"city"`
	Hobbies   string `json:"hobbies"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type RegisterCommandResult struct {
	UserId int64 `json:"userId"`
}

type RegisterHandler = application.Handler[RegisterCommand, *RegisterCommandResult]

type registerHandler struct {
	db *sql.DB
	key string
}

func NewRegisterHandler(db *sql.DB,key string) RegisterHandler {
	return &registerHandler{db,key}
}

func (receiver *registerHandler) Handle(ctx context.Context, arg RegisterCommand) (*RegisterCommandResult, error) {

	tx,err := receiver.db.BeginTx(ctx,nil)

	if err != nil {
		return nil, err
	}

	defer database.FixTx(tx,&err)

	r := db.NewRepository(tx)


	password, err := common.Encrypt([]byte(receiver.key),[]byte(arg.Password))

	_,err = r.AddAuth.Handle(ctx,&db.AddAuthQuery{
		Login:    arg.Login,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	auth,err := r.GetAuthByLogin.Handle(ctx,&db.GetAuthByLoginQuery{
		Login:    arg.Login,
		})

	if err != nil {
		return nil, err
	}

	_,err = r.SaveProfile.Handle(ctx,&db.SaveProfileQuery{
		UserId:    auth[0].UserId,
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
		Age:       arg.Age,
		Gender:    arg.Gender,
		City:      arg.City,
		Hobbies:   arg.Hobbies,
	})

	if err != nil {
		return nil, err
	}


	return &RegisterCommandResult{UserId: auth[0].UserId }, nil
}
