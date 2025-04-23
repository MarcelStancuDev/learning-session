package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://admin:admin@postgres:5432/taskdb?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	log.Println("Connected to PostgreSQL")
}