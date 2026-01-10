package usecases

import "go-flight-search/internal/domain"

type RedisCache interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, ttlSeconds int) error
}

type FlightProvider interface {
	Name() string
	Search(query domain.SearchQuery) (*[]domain.Flight, error)
}

type SearchFlightsUsecase struct {
	Providers []FlightProvider
	Cache     RedisCache
}
