package retrier

import (
	"log"
	"time"
)

type RetryFunc func() error
type RetryConfig struct {
	MaxRetries  int
	Interval    time.Duration
	ShouldRetry func(error) bool
	BackoffFunc func(int, time.Duration) time.Duration
}

var DefaultRetryConfig = RetryConfig{
	MaxRetries: 3,
	Interval:   time.Second,
	ShouldRetry: func(err error) bool {
		return true // Retry for any error by default
	},
	BackoffFunc: func(attempt int, interval time.Duration) time.Duration {
		return interval
	},
}

// Retry attempts to execute a RetryFunc until it succeeds or
// the maximum number of retries, defined in RetryConfig, is
// reached.
//
// The function fn is the operation to be retried and config
// contains retry parameters such as maximum retries, backoff
// strategy, and a condition to determine whether to retry
// after a failure.
// Returns an error if all attempts fail, nil otherwise.
func Retry(fn RetryFunc, config RetryConfig) error {
	var err error

	for attempt := 1; attempt <= config.MaxRetries; attempt++ {
		err = fn()

		if err == nil {
			// Operation succeeded, no need to retry
			return nil
		}

		log.Printf("Attempt %d failed: %v", attempt, err)

		if attempt < config.MaxRetries && config.ShouldRetry(err) {
			// Wait for the specified interval or until the operation completes
			<-time.After(config.BackoffFunc(attempt, config.Interval))
			// If the operation should be retried, continue to the next attempt
		}
	}

	return err
}
