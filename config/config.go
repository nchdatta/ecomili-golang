package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nchdatta/ecomili-golang/internal/helpers"
)

type Config struct {
	App      AppConfig
	Database DBConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Port        int
	Environment string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret   string
	ExpireAt string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// Initialize the Config instance with Env
	cfg := Config{
		App: AppConfig{
			Port:        helpers.GetIntFromEnv("PORT"),
			Environment: os.Getenv("ENVIRONMENT"),
		},
		Database: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     helpers.GetIntFromEnv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret:   os.Getenv("JWT_SECRET"),
			ExpireAt: os.Getenv("JWT_EXPIRE"),
		},
	}

	return &cfg, nil
}
