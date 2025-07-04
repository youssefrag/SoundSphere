package models

import (
	"github.com/youssefrag/SoundSphere/db"
	"github.com/youssefrag/SoundSphere/utils"
)

type User struct {
	ID int64
	Name string
	Email string `binding:"required"`
	Password string `binding:"required"`
	ImageUrl string
}

func (u *User) Save() error {

	var (
		query string
		args []interface{}
	)

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	
	args = []interface{}{u.Name, u.Email, hashedPassword}

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


