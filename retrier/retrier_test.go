package retrier

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRetrier_Failed(t *testing.T) {
	err := Retry(func() error {
		return fmt.Errorf("Simulating a failure")
	}, DefaultRetryConfig)

	if err != nil {
		log.Printf("Operation failed after retries: %v", err)
	} else {
		t.Errorf("Expected error, got nil")
	}
}

func TestRetrier_Success(t *testing.T) {
	err := Retry(func() error {
		return nil
	}, DefaultRetryConfig)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestCustomBackoff(t *testing.T) {
	err := Retry(func() error {
		return fmt.Errorf("Simulating a failure")
	}, RetryConfig{
		MaxRetries: 5,
		Interval:   time.Second,
		ShouldRetry: func(err error) bool {
			return true
		},
		BackoffFunc: func(attempt int, interval time.Duration) time.Duration {
			return interval * time.Duration(attempt)
		},
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
