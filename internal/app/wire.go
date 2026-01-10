package app

import (
	"github.com/gorilla/mux"
	"go-flight-search/internal/handlers/httphandlers"
	"go-flight-search/internal/handlers/middlewares"
	"go-flight-search/internal/repositories/cache"
	"go-flight-search/internal/repositories/providers/garuda"
	"go-flight-search/internal/repositories/resilience"
	"go-flight-search/internal/usecases"
	"net/http"
)

func Run() *http.Server {
	// TODO:
	// - load config
	// - init redis
	// - init providers
	GarudaProvider := garuda.New("mock_data/garuda_indonesia_search_response.json", 50, 100)
	RedisCache, err := cache.New("redis:6379", "supersecretpassword")
	if err != nil {
		panic(err)
	}
	flightProviders := []usecases.FlightProvider{
		resilience.WrapProvider(GarudaProvider),
	}
	// - init usecase
	FlightSearchUsecase := usecases.SearchFlightsUsecase{
		Providers: flightProviders,
		Cache:     RedisCache,
	}
	// - init handler
	SearchHandler := &httphandlers.SearchHandler{
		SearchFlightUseCase: &FlightSearchUsecase,
	}
	// - init router
	router := mux.NewRouter()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.RecoverMiddleware)

	router.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/search", SearchHandler.Search).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
