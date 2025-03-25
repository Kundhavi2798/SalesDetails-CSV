package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

const (
	DB_USER = "admin"
	DB_PASS = "kundhavi"
	DB_NAME = "users"
	DB_HOST = "localhost"
	DB_PORT = 5432
)

func InitDB() {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	var err error
	DB, err = sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database is unreachable:", err)
	}

	fmt.Println("Database initialized successfully!")
}
