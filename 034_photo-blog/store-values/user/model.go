package user

// User ...User model
type User struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
}

//NewUser ...Create new user
func NewUser(FirstName, LastName, Username, Password string) User {
	return User{
		FirstName,
		LastName,
		Username,
		Password,
	}
}
