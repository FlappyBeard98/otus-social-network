package types

type User struct {
	UserId  int64
	Auth    *Auth
	Profile *Profile
	Friends []Friend
}
