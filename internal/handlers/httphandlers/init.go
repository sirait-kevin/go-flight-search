package httphandlers

import (
	"context"
	"go-flight-search/internal/domain"
)

type SearchHandler struct {
	SearchFlightUseCase SearchFlightUseCase
}

type SearchFlightUseCase interface {
	Execute(ctx context.Context, q domain.SearchQuery) (*[]domain.Flight, bool, error)
}
