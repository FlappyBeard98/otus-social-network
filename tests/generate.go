package tests

import (
	"context"
	"fmt"
	"social-network/app"
	"social-network/app/command"

	"github.com/bxcodec/faker"
)

func GenerateUsers(a *app.App,count int){

	ctx := context.Background()
	_, err := a.Commands.ClearUsers.Handle(ctx,nil)

	if err != nil {
		panic(err)
	}
	
	users := make([]command.RegisterCommand,0)
	for i := 0; i < count; i++ {
		users= append(users,  generateRegisterCommand())

		
	}

	_,err = a.Commands.Register.Handle(ctx,users[0])	
		if err != nil {
			fmt.Println(err)
		}
	
}

func generateRegisterCommand()command.RegisterCommand{
	cmd := command.RegisterCommand{}
	err := faker.FakeData(&cmd)
	if err != nil {
		fmt.Println(err)
	}

	return cmd
}