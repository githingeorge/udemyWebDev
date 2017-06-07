package user

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

var dbUsers = map[string]User{}

func GetByUserName(username string) (User, bool) {
	user, ok := dbUsers[username]
	return user, ok

}
