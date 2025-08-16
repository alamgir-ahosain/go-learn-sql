package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func Connect() {

	con := "root:mysql@tcp(127.0.0.1:3306)/goDB" //DSN (Data Source Name) formats for Go + MySQL.(connection string)
	var err error
	DB, err = sql.Open("mysql", con)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	// Test the connection immediately
	if err := DB.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Successfully connected to MySQL database")
}
