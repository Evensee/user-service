package internal

import (
	"fmt"

	"github.com/evensee/go-tl/dotenv"
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

type AppConfig struct {
	Secret string `env:"APP_SECRET"`

	GrpcApiPort int `env:"USER_SERVICE_GRPC_API_PORT"`
	HttpApiPort int `env:"USER_SERVICE_HTTP_API_PORT"`

	AccessTokenLifetimeSeconds  int `env:"ACCESS_TOKEN_LIFETIME_SECONDS"`
	RefreshTokenLifetimeSeconds int `env:"REFRESH_TOKEN_LIFETIME_SECONDS"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	Db int `env:"REDIS_DB"`
}

func MustLoadConfig[C any]() *C {
	config := dotenv.MakeConfig(new(C))

	return config
}
