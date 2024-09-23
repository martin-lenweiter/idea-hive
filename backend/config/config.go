package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")

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
