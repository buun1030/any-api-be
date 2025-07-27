package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Default port
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://user:password@localhost:5432/anyapi_db?sslmode=disable" // Default for local development
	}

	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}
}