package main

import (
	"os"
	"social-network/lib/http"
	"social-network/lib/utils"
	"social-network/profile/internal/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	enviroment = "../docker-compose.yml"
	migrator   = "../deployments/docker-compose.yml"
	sleep      = 5 * time.Second
	host       = "http://localhost:1323"
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

	login := "test_login"
	password := "test_password"

	//register user with auth and profile
	registerRequest := types.RegisterRequest{
		Login:     login,
		Password:  password,
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Age:       18,
		Gender:    1,
		City:      "test_city",
		Hobbies:   "test_hobbies",
	}

	profile, err := http.GetHttpResponse[types.Profile](&registerRequest, host,nil)
	assert.NoError(t, err)
	assert.NotNil(t, profile)

	//read profiles
	profilesRequest := types.ProfilesRequest{
		PageRequest: types.PageRequest{
			Limit:  1,
			Offset: 0,
		},
	}

	profiles, err := http.GetHttpResponse[types.PageResponse[types.Profile]](&profilesRequest, host,nil)
	assert.NoError(t, err)
	assert.NotNil(t, profiles)

	//change profile
	profile.FirstName = "test_first_name_changed"
	profile,err = http.GetHttpResponse[types.Profile](profile, host,http.SetBasicAuth(login, password))
	assert.NoError(t, err)
	assert.NotNil(t, profile)

	//get profile
	getProfileRequest := types.GetProfileRequest{ 
		UserId: profile.UserId, 
	}
	profile,err = http.GetHttpResponse[types.Profile](&getProfileRequest, host,http.SetBasicAuth(login, password))
	assert.NoError(t, err)
	assert.NotNil(t, profile)

	//add friend
	addFriendRequest := types.AddFriendRequest{
		UserId: profile.UserId,
	}
	friend,err := http.GetHttpResponse[types.Friend](&addFriendRequest, host,http.SetBasicAuth(login, password))
	assert.NoError(t, err)
	assert.NotNil(t, friend)

    //get friends
	getFriendsRequest := types.GetFriendsRequest{
		PageRequest: types.PageRequest{
			Limit:  1,
			Offset: 0,
		},
		UserIdRequest: types.UserIdRequest{
			UserId : profile.UserId,
		},	
	}
	profiles,err = http.GetHttpResponse[types.PageResponse[types.Profile]](&getFriendsRequest, host,http.SetBasicAuth(login, password))
	assert.NoError(t, err)
	assert.NotNil(t, profiles)

	//remove friend
	removeFriendRequest := types.RemoveFriendRequest{
		UserId: profile.UserId,
	}
	friend,err = http.GetHttpResponse[types.Friend](&removeFriendRequest, host,http.SetBasicAuth(login, password))
	assert.NoError(t, err)
	assert.NotNil(t, friend)

}
