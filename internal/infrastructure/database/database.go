package database

import (
	"fmt"
	"github.com/Evensee/user-service/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseConfig internal.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.DatabaseName,
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return connection
}
