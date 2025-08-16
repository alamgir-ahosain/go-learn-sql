package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	
	con := "root:mysql@tcp(127.0.0.1:3306)/testdb" //DSN (Data Source Name) formats for Go + MySQL.(connection string)
	db, err := sql.Open("mysql", con)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}
	defer db.Close() //close the DB connection when main exist
	
	// Test the connection immediately
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	fmt.Println("Successfully connected to MySQL database")
}
