package models

import "github.com/youssefrag/SoundSphere/db"

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
	ImageUrl string
}

var users = []User{}

func (u *User) Save() error {

	var (
		query string
		args []interface{}
	)
	
	args = []interface{}{u.Name, u.Email, u.Password}

	if u.ImageUrl != "" {
		// include image_url
		query = `
			INSERT INTO users (name, email, password, image_url)
			VALUES ($1, $2, $3, $4)
			RETURNING id;
		`

		args = append(args, u.ImageUrl)
	} else {
		query = `
			INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)
			RETURNING id;
    `
  }

	return db.DB.QueryRow(query, args...).Scan(&u.ID)
}


// func (u *User) Save() error {

//   const query = `
// 		INSERT INTO users(name, email, password)
// 		VALUES ($1, $2, $3)
// 		RETURNING id;
// 	`

//   return db.DB.
//     QueryRow(query, u.Name, u.Email, u.Password).
//     Scan(&u.ID)
// }

func GetAllUsers() []User {
 return users
}