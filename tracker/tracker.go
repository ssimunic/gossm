package tracker

import (
	"time"
)

// TimeTracker is used to keep track of something
type TimeTracker struct {
	delayer  Delayer
	nextTime time.Time
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
	nextDelay := t.delayer.Delay()
	t.nextTime = time.Now().Add(nextDelay)
	return nextDelay, t.nextTime
}

// New returns pointer to new TimeTracker and sets its Delayer
func New(delayer Delayer) *TimeTracker {
	return &TimeTracker{delayer: delayer}
}
