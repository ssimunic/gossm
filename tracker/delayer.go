package tracker

import (
	"time"
)

type Delayer interface {
	Delay() time.Duration
}
