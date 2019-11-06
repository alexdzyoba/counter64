// package counter64 provides uint64 incrementing counter
package counter64

import (
	"math/rand"
	"time"
)

// Counter type allows you to launch increment of the counter and read its
// value.
type Counter interface {
	Count(done chan bool)
	Read() uint64
}

// New creates a simple counter that increments by one as fast as it can.
func New() Counter {
	return &counter{}
}

// NewTicked returns ticked counter that increments in incr at each tick.
func NewTicked(incr uint64, tick time.Duration) Counter {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := time.NewTicker(tick)

	return &counterTicked{
		incr:   incr,
		ticker: t,
		rand:   r,
	}
}
