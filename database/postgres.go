package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Wine defines the structure of a wine object in database
type Wine struct {
	gorm.Model
	Name            string `json:"name"`
	Domain          string `json:"domain"`
	Year            int64  `json:"year"`
	AppellationType string `json:"appellation_type"`
	AppellationName string `json:"appellation_name"`
}

// DatabaseConnection initializes the connection to database
func DatabaseConnection() {
	// Find .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Retrieve DB configuration
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	// Connect to DB and migrate structures
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Wine{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
