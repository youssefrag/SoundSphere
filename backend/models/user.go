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

func (u *User) Save() error {
  // 1) Use Postgres placeholders ($1, $2, $3)
  // 2) Add RETURNING id so we can grab the new primary key
  const query = `
		INSERT INTO users(name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

  // QueryRow + Scan is the simplest way to get the new id
  return db.DB.
    QueryRow(query, u.Name, u.Email, u.Password).
    Scan(&u.ID)
}

func GetAllUsers() []User {
 return users
}