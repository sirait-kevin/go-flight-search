package providers

import (
	"go-flight-search/internal/domain"
	"time"
)

type FlightProvider interface {
	Name() string
	Search(query domain.SearchQuery) (*[]domain.Flight, error)
	BaseDelay() time.Duration
	MaxDelay() time.Duration
}
