package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgresql://dev_user:dev_password@localhost:5432/idea_repository_dev?sslmode=disable"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
