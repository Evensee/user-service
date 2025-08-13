package http

import (
	"net/http"
)

func NewRouter(userHandler *UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", userHandler.CreateUser)

	return mux
}
