package main

import (
	"os"
	"social-network/lib/utils"
	"testing"
	"time"
)

var (
	enviroment = "../docker-compose.yml"
	migrator = "../deployments/docker-compose.yml"
	sleep = 5 * time.Second
)

func TestMain(m *testing.M) {
	utils.ComposeUp(enviroment)

	time.Sleep(sleep)

	exitVal := 1
	defer func() {
		utils.ComposeDown(enviroment)
		os.Exit(exitVal)
	}()

	go main()
	
	time.Sleep(sleep)

	exitVal = m.Run()
}

func TestApi(t *testing.T) {
//register
//login
//get profiles
//change profile
//get profile
//add friend
//get friends
//remove friend	

}

type api struct{
	testHost string
}


