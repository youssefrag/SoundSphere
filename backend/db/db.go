package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Could not open DB: %v", err))
	}

	for i := 0; i < 5; i++ {
		err = DB.Ping()
		if err == nil {
			break
		}
		fmt.Printf("DB not ready, retrying… (%v)\n", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(fmt.Sprintf("Could not ping the database: %v", err))
	}


	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
	fmt.Println("✅ Connected to Postgres and ensured tables exist.")
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

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	fmt.Println("Users table has been created!")
}