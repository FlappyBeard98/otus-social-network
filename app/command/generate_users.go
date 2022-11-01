package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
	"social-network/common/database"
	"social-network/db"

	"github.com/bxcodec/faker"
)

type GenerateUsersHandler = application.Handler[int, any]

type generateUsersHandler struct {
	db  *sql.DB
}

func NewClearUsersHandler(db *sql.DB) GenerateUsersHandler {
	return &generateUsersHandler{db}
}

func (receiver *generateUsersHandler) Handle(ctx context.Context, arg int) (any, error) {
	
	handler := database.NewDbHandler[*db.ClearUsersQuery](receiver.db)

	_,err := handler.Handle(ctx, &db.ClearUsersQuery{})

	if err !=nil {
		return nil,err
	}

	users := make([]RegisterCommand,0)
	for i := 0; i < arg; i++ {
		cmd, _ :=  generateRegisterCommand()
		users = append(users, cmd)		
	}

	handler := database.NewDbHandler[*db.ClearUsersQuery](receiver.db)
	
	return nil, nil
}

func generateRegisterCommand()(RegisterCommand,error){
	
	cmd := RegisterCommand{}
	
	err := faker.FakeData(&cmd)
	if err != nil {
		return cmd,err
	}

	return cmd,err
}