package main

import (
	"github.com/Evensee/user-service/internal"
	"github.com/Evensee/user-service/internal/application/usecase"
	http2 "github.com/Evensee/user-service/internal/delivery/http"
	"github.com/Evensee/user-service/internal/domain/user"
	"github.com/Evensee/user-service/internal/infrastructure/database"
	"github.com/Evensee/user-service/internal/infrastructure/database/repository"
	"log"
	"net/http"
)

func main() {
	dbConfig, err := internal.LoadDatabaseConfig()

	if err != nil {
		panic(err)
	}

	// Database connection
	db := database.Connect(*dbConfig)

	// Setup repositories
	userRepo := repository.NewUserRepository(db)

	// Setup domain service
	userService := user.NewUserDomainService(userRepo)

	// Setup use case
	userUserCase := usecase.NewUserUseCase(userService)

	// Setup HTTP handler
	mux := http.NewServeMux()
	userHandler := http2.NewUserHandler(userUserCase)

	mux.HandleFunc("/users", userHandler.CreateUser) // CreateUser has correct signature

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
