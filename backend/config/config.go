package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	err := godotenv.Load() // Load .env file
	// log.Println name of environment variable and its value
	log.Println("ENVIRONMENT:", os.Getenv("ENVIRONMENT"))
	if err != nil {
		log.Println("Error loading .env file")
	}

	var dbURL, serverAddress string

	dbURL = os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if dbURL == "" || port == "" {
		log.Fatal("Missing DATABASE_URL or PORT environment variables")
	}

	serverAddress = ":" + port

	return &Config{
		ServerAddress: serverAddress,
		DatabaseURL:   dbURL,
	}, nil
}
