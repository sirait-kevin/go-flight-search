package app

import (
	"github.com/gorilla/mux"
	"go-flight-search/internal/handlers/middlewares"
	"net/http"
)

func Run() *http.Server {
	// TODO:
	// - load config
	// - init redis
	// - init providers
	// - init usecase
	// - init router

	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.RecoverMiddleware)

	router.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
