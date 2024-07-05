package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type (
	Config struct {
		Env string
		HttpService
		PG
	}

	HttpService struct {
		Address     string
		Timeout     time.Duration
		IdleTimeout time.Duration
	}

	PG struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
		SSLMode  string
	}
)

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	config := &Config{}

	config.Env = os.Getenv("ENV")

	config.HttpService.Address = os.Getenv("HTTP_ADDRESS")
	config.HttpService.Timeout, err = time.ParseDuration(os.Getenv("HTTP_TIMEOUT"))
	if err != nil {
		log.Fatalf("error loading .env timeout param: %v", err)
	}
	config.HttpService.IdleTimeout, err = time.ParseDuration(os.Getenv("HTTP_IDLE_TIMEOUT"))
	if err != nil {
		log.Fatalf("error loading .env timeout param: %v", err)
	}

	config.PG.Host = os.Getenv("PG_HOST")
	config.PG.Port = os.Getenv("PG_PORT")
	config.PG.Username = os.Getenv("PG_USERNAME")
	config.PG.Password = os.Getenv("PG_PASSWORD")
	config.PG.Database = os.Getenv("PG_DATABASE")
	config.PG.SSLMode = os.Getenv("PG_SSLMODE")

	return config, nil
}
