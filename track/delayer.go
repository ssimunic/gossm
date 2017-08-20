package track

import (
	"time"
)

// Delayer is used to delay some action
type Delayer interface {
	Delay() time.Duration
}
