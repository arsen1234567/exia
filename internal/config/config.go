package config

import (
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
	// Optionally load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()

	// Load config from environment variables
	var cfg Config
	cfg.Server.Address = os.Getenv("SERVER_ADDRESS")
	cfg.Database.Driver = os.Getenv("DB_DRIVER")
	cfg.Database.URL = os.Getenv("DB_URL")

	return cfg
}
