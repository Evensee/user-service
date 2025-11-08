package http

import (
	"encoding/json"
	"net/http"

	"github.com/Evensee/user-service/internal/application/contract"
	"github.com/Evensee/user-service/internal/application/usecase"
)

type UserHandler struct {
	userUserCase *usecase.UserUseCase
}

func NewUserHandler(userUserCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUserCase: userUserCase}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input contract.CreateUserContract

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	output, err := handler.userUserCase.CreateUser(r.Context(), input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		return
	}
}
