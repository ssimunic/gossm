package track

import (
	"time"
)

// ExpBackoff is used as delayer, implements exponential backoff algorithm
type ExpBackoff struct {
	counter int
	base    int
}

func calculateExponential(base, counter int) int {
	if counter == 0 {
		return 1
	}
	return base * calculateExponential(base, counter-1)
}

// Delay returns seconds
func (e *ExpBackoff) Delay() time.Duration {
	e.counter++
	return time.Duration(calculateExponential(e.base, e.counter)) * time.Second
}

// NewExpBackoff returns pointer to new ExpBackoff
// base is meant to be seconds
func NewExpBackoff(base int) *ExpBackoff {
	return &ExpBackoff{
		base: base,
	}
}
