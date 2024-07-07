package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type (
	Config struct {
		Env string `env:"ENV,required"`
		HttpService
		PG
		ExternalApi
	}

	HttpService struct {
		Address     string        `env:"HTTP_ADDRESS,required"`
		Timeout     time.Duration `env:"HTTP_TIMEOUT,required"`
		IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT,required"`
	}

	PG struct {
		Host     string `env:"PG_HOST,required"`
		Port     string `env:"PG_PORT,required"`
		Username string `env:"PG_USERNAME,required"`
		Password string `env:"PG_PASSWORD,required"`
		Database string `env:"PG_DATABASE,required"`
		SSLMode  string `env:"PG_SSLMODE,required"`
	}

	ExternalApi struct {
		PeopleApi string `env:"PEOPLES_API"`
	}
)

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	config := &Config{}

	err = env.Parse(config)
	if err != nil {
		log.Fatalf("error to parse .env variables: %v", err)
	}

	return config, nil
}
