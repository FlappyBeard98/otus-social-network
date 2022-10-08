package model

type Profile struct {
	Id        int64
	UserId    int64
	FirstName string
	LastName  string
	Age       int32
	Gender    bool
	City      string
	Hobbies   string
	Friends   []Profile
}
