package dbsqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "storage/goDB.db")

	if err != nil {
		log.Fatal(err)
	}

	//create users table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sid TEXT,
    name TEXT,
    cgpa DECIMAL(5,2)
);

    `
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'users' created successfully")

}
