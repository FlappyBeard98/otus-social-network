package db

type Profile struct {
	UserId    int64  `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int32  `db:"age"`
	Gender    int32   `db:"gender"`
	City      string `db:"city"`
	Hobbies   string `db:"hobbies"`
}

type Auth struct {
	UserId    int64     `db:"user_id"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
}

type Friend struct {
	UserId       int64  `db:"user_id"`
	FriendUserId string `db:"friend_user_id"`
}
