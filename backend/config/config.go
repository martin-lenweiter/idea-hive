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
	log.Println("Loading configuration...")

	for _, env := range os.Environ() {
		log.Println(env)
	}

	dbURL := os.Getenv("DATABASE_URL")
	log.Println("Database URL from env:", dbURL)

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
