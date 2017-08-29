package track

import (
	"time"
)

// TimeTracker is used to keep track of something
type TimeTracker struct {
	delayer  Delayer
	nextTime time.Time
	// counter is used to remember how many times has delayer been ran
	counter int
}

// IsReady checks if current time is after nextTime
// On first check before calling SetNext, always true due to nextTime having zero value
func (t *TimeTracker) IsReady() bool {
	if time.Now().After(t.nextTime) {
		return true
	}
	return false
}

// SetNext updates nextTime based on delayer implementation
// Returns delay and time at which will current time be after it
func (t *TimeTracker) SetNext() (time.Duration, time.Time) {
	t.counter++
	nextDelay := t.delayer.Delay()
	t.nextTime = time.Now().Add(nextDelay)
	return nextDelay, t.nextTime
}

// NewTracker returns pointer to new TimeTracker and sets its Delayer
func NewTracker(delayer Delayer) *TimeTracker {
	return &TimeTracker{delayer: delayer}
}

// HasBeenRan checks how many times has time delayer and returns true if ever ran
func (t *TimeTracker) HasBeenRan() bool {
	if t.counter > 0 {
		return true
	}
	return false
}
