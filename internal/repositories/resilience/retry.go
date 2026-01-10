package resilience

import "time"

type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func Retry(cfg RetryConfig, fn func() error) error {
	var err error
	delay := cfg.BaseDelay

	for i := 0; i < cfg.MaxAttempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		time.Sleep(delay)

		delay *= 2
		if delay > cfg.MaxDelay {
			delay = cfg.MaxDelay
		}
	}

	return err
}
