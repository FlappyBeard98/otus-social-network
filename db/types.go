package db

// Profile struct in database
type Profile struct {
	UserId    int64  `db:"user_id"`    //user identifier
	FirstName string `db:"first_name"` //user first name
	LastName  string `db:"last_name"`  //user last name
	Age       int32  `db:"age"`        //user age
	Gender    int32  `db:"gender"`     //user gender
	City      string `db:"city"`       //user city
	Hobbies   string `db:"hobbies"`    //user hobbies
}

// Auth struct in database
type Auth struct {
	UserId   int64  `db:"user_id"`  //user identifier
	Login    string `db:"login"`    //user login
	Password string `db:"password"` // encrypted password
}
