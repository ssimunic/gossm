package tracker

import (
	"time"
)

// TimeTracker is used to time notifications
type TimeTracker struct {
	delayer  Delayer
	nextTime time.Time
}

func (t *TimeTracker) IsReady() bool {
	if time.Now().After(t.nextTime) {
		return true
	}
	return false
}

func (t *TimeTracker) SetNext() (time.Duration, time.Time) {
	nextDelay := t.delayer.Delay()
	t.nextTime = time.Now().Add(nextDelay)
	return nextDelay, t.nextTime
}

func NewTimeTracker(delayer Delayer) *TimeTracker {
	return &TimeTracker{delayer: delayer}
}
