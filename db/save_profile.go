package db

import "social-network/common"

type SaveProfileQuery struct {
	UserId    int64
	FirstName string
	LastName  string
	Age       int32
	Gender    int32
	City      string
	Hobbies   string
}

func (receiver *SaveProfileQuery) Sql() string {
	return `
INSERT INTO social_network.profiles(user_id, first_name, last_name, age, gender, city, hobbies)
VALUES (?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE 
     first_name = ?
    ,last_name = ?
    ,age = ?
    ,gender = ?
    ,city = ?
    ,hobbies = ?
;
`
}

func (receiver *SaveProfileQuery) GetParams() []any {
	params := common.GetFieldsValuesAsSlice(receiver)
	return append(params,params[1:]...)
}
