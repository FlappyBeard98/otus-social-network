package model

import "social-network/db"

type Profile struct {
	UserId    int64 `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
	Gender    int32   `json:"gender"`
	City      string `json:"city"`
	Hobbies   string `json:"hobbies"`
}

func NewProfileFromDb(profile db.Profile) Profile  {

	return Profile{
		UserId:    profile.UserId,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Age:       profile.Age,
		Gender:    profile.Gender,
		City:      profile.City,
		Hobbies:   profile.Hobbies,
		}
}

type PageInfo struct {
	From int `json:"from"`
	Count int `json:"count"`
	Total int `json:"total"`
}