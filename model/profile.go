package model

import "social-network/db"

// Profile user profile data
type Profile struct {
	UserId    int64  `json:"userId"`    //user identifier
	FirstName string `json:"firstName"` //user first name
	LastName  string `json:"lastName"`  //user last name
	Age       int32  `json:"age"`       //user age
	Gender    int32  `json:"gender"`    //user gender
	City      string `json:"city"`      //user city
	Hobbies   string `json:"hobbies"`   //user hobbies
}

// NewProfileFromDb creates Profile from db.Profile
func NewProfileFromDb(profile db.Profile) Profile {

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

// PageInfo represents requested page metadata
type PageInfo struct {
	From  int `json:"from"`  //page start position
	Count int `json:"count"` //number of items per page
	Total int `json:"total"` //total elements
}
