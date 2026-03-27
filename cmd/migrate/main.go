package main

import (
	"log"

	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/infrastructure/database"
	"github.com/Evensee/user-service/internal/infrastructure/database/model"
)

func main() {
	databaseConfig := internal.MustLoadConfig[internal.DatabaseConfig]()

	db := database.Connect(databaseConfig)

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&model.UserORMModel{}); err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}

}
