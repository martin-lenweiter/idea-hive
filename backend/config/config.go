package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	log.Println("DB_URL:", dbURL)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverAddress := ":" + port

	return &Config{
		ServerAddress: serverAddress,
		DatabaseURL:   dbURL,
	}, nil
}
