package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() {
	// Load environment variables from .env file if ENV is not set
	if _, exists := os.LookupEnv("ENV"); !exists {
		dotenv := ".local.env"
		if err := godotenv.Load(dotenv); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Retrieve values from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Construct the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not connect to db", err)
	} else {
		fmt.Println("Connected to db")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
	} else {
		fmt.Println("Successfully pinged to db")
	}
}
