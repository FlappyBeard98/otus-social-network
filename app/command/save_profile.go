package command

import (
	"context"
	"database/sql"
	"social-network/common/application"
	"social-network/db"
)

type SaveProfileCommand struct {
	UserId int64 `param:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
	Gender    int32   `json:"gender"`
	City      string `json:"city"`
	Hobbies   string `json:"hobbies"`
}

type SaveProfileHandler = application.Handler[SaveProfileCommand, interface{}]

type saveProfileHandler struct {
	db *sql.DB
}

func NewSaveProfileHandler(db *sql.DB) SaveProfileHandler {
	return &saveProfileHandler{db}
}

func (receiver *saveProfileHandler) Handle(ctx context.Context, arg SaveProfileCommand) (interface{}, error) {

	r := db.NewRepository(receiver.db)

	_,err := r.SaveProfile.Handle(ctx,&db.SaveProfileQuery{
		UserId:    arg.UserId,
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


	return nil, nil
}
