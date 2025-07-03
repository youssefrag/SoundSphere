package models

import (
	"github.com/youssefrag/SoundSphere/db"
)

type User struct {
	ID int64
	Name string
	Email string
	Password string
}

var users = []User{}

func (u *User) Save() error{
	query := "INSERT INTO users(name, email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetAllUsers() []User {
 return users
}