package apistuff

import (
	"time"
)

// Used to stub out time in testing
type Clock interface {
	Now() time.Time
}

type SystemClock struct{}

func (s SystemClock) Now() time.Time {
	return time.Now()
}
