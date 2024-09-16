package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Using default environment variables.")
	}
}

func GetMongoDBURL() string {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}
	return mongoURI
}