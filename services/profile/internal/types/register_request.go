package types

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

// NewAuth returns new Auth from RegisterRequest
func (o *RegisterRequest) NewAuth() (*Auth, error) {
	return NewAuth(o.Login, o.Password)
}

// NewProfile returns new Profile from RegisterRequest
func (o *RegisterRequest) NewProfile() (*Profile, error) {
	return NewProfile(o.FirstName, o.LastName, o.Age, o.Gender, o.City, o.Hobbies)
}
