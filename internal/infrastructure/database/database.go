package database

import (
	"fmt"

	"github.com/Evensee/user-service/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseConfig *internal.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.DatabaseName,
	)

	gormConfig := gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gormConfig)

	if err != nil {
		panic(err)
	}

	return connection
}
