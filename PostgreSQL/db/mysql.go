package db

import (
	"database/sql"
	"fmt"
	"log"
		// _ "github.com/go-sql-driver/mysql"
		_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	con := "postgres://postgres:postgresql@127.0.0.1:5432/goDB?sslmode=disable"
	//DSN (Data Source Name) formats for Go + PostgreSql.(connection string)
	var err error
	DB, err = sql.Open("postgres", con)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	// Test the connection immediately
	if err := DB.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Successfully connected to PostgreSql database")
}
