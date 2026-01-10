package resilience

import (
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("circuit breaker open")

type CircuitBreaker struct {
	mu sync.Mutex

	failures     int
	state        string // "closed", "open", "half-open"
	lastFailure  time.Time
	openTimeout  time.Duration
	failureLimit int
}

func NewCircuitBreaker(failureLimit int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        "closed",
		failureLimit: failureLimit,
		openTimeout:  openTimeout,
	}
}

func (cb *CircuitBreaker) Allow() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == "open" {
		if time.Since(cb.lastFailure) > cb.openTimeout {
			cb.state = "half-open"
			return nil
		}
		return ErrCircuitOpen
	}

	return nil
}

func (cb *CircuitBreaker) Success() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures = 0
	cb.state = "closed"
}

func (cb *CircuitBreaker) Failure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.lastFailure = time.Now()

	if cb.failures >= cb.failureLimit {
		cb.state = "open"
	}
}
