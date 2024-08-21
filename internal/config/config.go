package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Address string
	}
	Database struct {
		Driver string
		URL    string
	}
}

func LoadConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load config from environment variables
	var cfg Config
	cfg.Server.Address = os.Getenv("SERVER_ADDRESS")
	cfg.Database.Driver = os.Getenv("DB_DRIVER")
	cfg.Database.URL = os.Getenv("DB_URL")

	return cfg
}
