package gossm

import (
	"time"
)

// ExpBackoff is used to delay notifications in exponential way
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

// NextDelay returns seconds until next iteration
func (e *ExpBackoff) NextDelay() time.Duration {
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
