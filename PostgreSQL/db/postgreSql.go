package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	//load .env file
	err1 := godotenv.Load("config/.env")
	if err1 != nil {
		log.Fatal("Error loading .env file:", err1)
	}

	// Read env variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	portStr := os.Getenv("DB_PORT")

	//convert port string to int
	port, err2 := strconv.Atoi(portStr)
	if err2 != nil {
		log.Fatal("err2")
	}

	con := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
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
