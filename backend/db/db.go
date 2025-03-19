package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
    connStr := "user=postgres password=supersecret dbname=soundsphere host=database port=5432 sslmode=disable"
		var err error
    db, err = sql.Open("postgres", connStr)

		if err != nil {
			panic("Could not connect to database.")
		}

			// Test the connection
		err = db.Ping()
		if err != nil {
			panic(fmt.Sprintf("Could not ping the database: %v", err))
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)

		createTables()

		fmt.Println("Database Tables have been created")
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			imageUrl TEXT
		)
	`

	_, err := db.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	fmt.Println("Users table has been created!")
}