package model

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

var db *sql.DB

func Setup() {
	// Load environment variables from .env file
	dotenv := ".local.env"
	if err := godotenv.Load(dotenv); err != nil {
		log.Fatal("Error loading .env file")
	}
}
