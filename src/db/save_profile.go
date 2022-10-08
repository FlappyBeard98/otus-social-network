package db

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
INSERT INTO social_network.profile(user_id, first_name, last_name, age, gender, city, hobbies)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON DUPLICATE KEY UPDATE 
     first_name = $2
    ,last_name = $3
    ,age = $4
    ,gender = $5
    ,city = $6
    ,hobbies = $7
;
`
}
