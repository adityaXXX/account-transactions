package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Initializing DB globally
var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	// Database parameters
	DB_USER := "test"
	DB_NAME := "aditya"
	DB_PASSWORD := "assignment"
	DB_HOST := "localhost"
	DB_PORT := "5432"

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_PORT)
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Successfully connected to the database")
}
