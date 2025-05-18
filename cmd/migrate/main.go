package main

import (
	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/infrastructure/database"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
	"log"
)

func main() {
	databaseConfig, err := internal.LoadDatabaseConfig()

	if err != nil {
		panic(err)
	}

	connection := database.Connect(*databaseConfig)

	if err = connection.AutoMigrate(&model.UserORMModel{}); err != nil {
		log.Fatalf("failed to run migration: %w", err)
	}

}
