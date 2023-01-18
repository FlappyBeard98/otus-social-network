package types

import (
	"fmt"
	"social-network/lib/mysql"
	"social-network/lib/utils"
)

// Profile defines user profile data
type Profile struct {
	UserId    int64  `json:"userId" db:"user_id"`       //user identifier
	FirstName string `json:"firstName" db:"first_name"` //user first name
	LastName  string `json:"lastName" db:"last_name"`   //user last name
	Age       int32  `json:"age" db:"age"`              //user age
	Gender    int32  `json:"gender" db:"gender"`        //user gender
	City      string `json:"city" db:"city"`            //user city
	Hobbies   string `json:"hobbies" db:"hobbies"`      //user hobbies
}

// NewProfile creates new valid Profile or return error if validation failed
func NewProfile(userId int64,
	firstName string,
	lastName string,
	age int32,
	gender int32,
	city string,
	hobbies string) (*Profile, error) {
	if len(firstName) > 100 {
		return nil, fmt.Errorf("%w: firstName must be 100 characters long or less", ErrInvalidInput)
	}

	if len(lastName) > 100 {
		return nil, fmt.Errorf("%w: lastName must be 100 characters long or less", ErrInvalidInput)
	}

	if len(city) > 50 {
		return nil, fmt.Errorf("%w: city must be 100 characters long or less", ErrInvalidInput)
	}

	if age < 0 {
		return nil, fmt.Errorf("%w: age must be greater than 0", ErrInvalidInput)
	}

	if gender < 0 || gender > 2 {
		return nil, fmt.Errorf("%w: gender takes 0 for women, 1 for men and 2 for everyone else", ErrInvalidInput)
	}

	if len(hobbies) > 5000 {
		return nil, fmt.Errorf("%w: hobbies must be 5000 characters long or less", ErrInvalidInput)
	}

	return &Profile{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Gender:    gender,
		City:      city,
		Hobbies:   hobbies,
	}, nil
}

// UpsertProfile returns new mysql.SqlQuery for upserting user profile in database
func (o *Profile) UpsertProfile() *mysql.SqlQuery {
	params := utils.GetFieldsValuesAsSlice(o)
	params = append(params, params[1:]...)
	return mysql.NewSqlQuery(`
		INSERT INTO social_network.profiles(user_id, first_name, last_name, age, gender, city, hobbies)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE 
			first_name = ?
			,last_name = ?
			,age = ?
			,gender = ?
			,city = ?
			,hobbies = ?;`,
		params...)
}
