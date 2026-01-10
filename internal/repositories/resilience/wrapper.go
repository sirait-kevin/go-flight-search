package resilience

import (
	"go-flight-search/internal/domain"
	"go-flight-search/internal/repositories/providers"
	"time"
)

type ResilientProvider struct {
	inner providers.FlightProvider
	cb    *CircuitBreaker
	retry RetryConfig
}

func WrapProvider(p providers.FlightProvider) providers.FlightProvider {
	return &ResilientProvider{
		inner: p,
		cb:    NewCircuitBreaker(3, 10*time.Second),
		retry: RetryConfig{
			MaxAttempts: 3,
			BaseDelay:   p.BaseDelay(),
			MaxDelay:    p.MaxDelay(),
		},
	}
}

func (r *ResilientProvider) Name() string {
	return r.inner.Name()
}

func (r *ResilientProvider) BaseDelay() time.Duration {
	return r.inner.BaseDelay()
}

func (r *ResilientProvider) MaxDelay() time.Duration {
	return r.inner.MaxDelay()
}

func (r *ResilientProvider) Search(q domain.SearchQuery) (*[]domain.Flight, error) {
	// Check circuit breaker
	if err := r.cb.Allow(); err != nil {
		return nil, err
	}

	var result []domain.Flight

	err := Retry(r.retry, func() error {
		flights, err := r.inner.Search(q)
		if err != nil {
			return err
		}
		result = *flights
		return nil
	})

	if err != nil {
		r.cb.Failure()
		return nil, err
	}

	r.cb.Success()
	return &result, nil
}
