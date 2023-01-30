package types

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RegisterRequest used for user registration
type RegisterRequest struct {
	Login     string `json:"login"`     //user login
	Password  string `json:"password"`  // hashed password
	FirstName string `json:"firstName"` //user first name
	LastName  string `json:"lastName"`  //user last name
	Age       int32  `json:"age"`       //user age
	Gender    int32  `json:"gender"`    //user gender
	City      string `json:"city"`      //user city
	Hobbies   string `json:"hobbies" `  //user hobbies
}

func (o *RegisterRequest) NewAuth() (*Auth, error) {
	return NewAuth(o.Login, o.Password)
}

func (o *RegisterRequest) NewProfile() (*Profile, error) {
	return NewProfile(o.FirstName, o.LastName, o.Age, o.Gender, o.City, o.Hobbies)
}

func (o *RegisterRequest) CreateRequest(host string) (*http.Request, error) {
	body, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, host+"/register", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, nil
}
