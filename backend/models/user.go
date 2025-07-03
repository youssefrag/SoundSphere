package models

type User struct {
	ID int64
	Email string
	Password string
}

var users = []User{}

func (u User) Save() {
	// later: add it to a database
	users = append(users, u)
}

func GetAllUsers() []User {
 return users
}