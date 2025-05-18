package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type DatabaseConfig struct {
	Host string `env:"POSTGRES_HOST"`
	Port string `env:"POSTGRES_PORT"`

	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`

	DatabaseName string `env:"POSTGRES_DB"`

	AutoMigrate bool `env:"AUTO_MIGRATE"`
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found or could not be loaded")
	}

	databaseConfig := &DatabaseConfig{}

	if err := env.Load(databaseConfig, nil); err != nil {
		return nil, err
	}

	return databaseConfig, nil
}
