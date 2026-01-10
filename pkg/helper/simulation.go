package helper

import (
	"errors"
	"math/rand"
	"time"
)

func SimulateDelay(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func SimulateFailure(successRate float64) error {
	if rand.Float64() > successRate {
		return errors.New("provider temporary failure")
	}
	return nil
}
