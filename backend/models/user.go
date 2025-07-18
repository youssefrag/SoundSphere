package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
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
		// include imageurl
		query = `
			INSERT INTO users (name, email, password, imageurl)
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

	err = db.DB.QueryRow(query, args...).Scan(&u.ID)
	if err != nil {
			// check if it's a Postgres error
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
					// 23505 = unique_violation
					if pqErr.Code == "23505" && pqErr.Constraint == "users_email_key" {

							return errors.New("email already exists")
					}
			}
			return err
	}
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, name, password, imageurl FROM users WHERE email = $1
	`
	row := db.DB.QueryRow(query, u.Email)

	var (
		retrievedPassword string
		nsImage           sql.NullString
	)

	
	err := row.Scan(&u.ID, &u.Name, &retrievedPassword, &nsImage)


	if err != nil {
		return errors.New("credentials invalid")
	}

	if nsImage.Valid {
    u.ImageUrl = nsImage.String
  } else {
    u.ImageUrl = "" 
	}

	

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func RemoveRefreshTokenByJTI(jti string) error {
  const q = `DELETE FROM refresh_tokens WHERE jti = $1`
  if _, err := db.DB.Exec(q, jti); err != nil {
    return fmt.Errorf("could not delete refresh token: %w", err)
  }
  return nil
}


func GetUserByID(id int64) (User, error) {
  const q = `
    SELECT id, name, email, imageurl
    FROM users
    WHERE id = $1
  `
	var (
		u     User
		nsImg sql.NullString
	)

  err := db.DB.QueryRow(q, id).Scan(&u.ID, &u.Name, &u.Email, &nsImg)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return u, errors.New("user not found")
    }
    return u, err
  }

	if nsImg.Valid {
		u.ImageUrl = nsImg.String
	} else {
		u.ImageUrl = ""  // or set a default avatar URL
	}
	
  return u, nil
}

func GetUserIDByEmail(email string) (int64, error) {
	var id int64
	err := db.DB.QueryRow(
			`SELECT id FROM users WHERE email = $1`,
			email,
	).Scan(&id)
	if err != nil {
			if err == sql.ErrNoRows {
					return 0, ErrUserNotFound
			}
			return 0, err
	}
	return id, nil
}

var ErrUserNotFound = fmt.Errorf("user not found")