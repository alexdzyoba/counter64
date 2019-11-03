package counter64

import (
	"math/rand"
	"time"
)

// counterTicked implements Counter interface but is more light on resources
// because it adds increment in batch at every tick.
type counterTicked struct {
	i uint64

	incr   uint64
	ticker *time.Ticker

	rand *rand.Rand
}

func (c counterTicked) Count(done chan bool) {
	defer c.ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-c.ticker.C:
			coeff := c.rand.Float64() + 1 // random in [1.0, 2.0)
			c.i += uint64(float64(c.incr) * coeff)
		}
	}
}

func (c counterTicked) Read() uint64 {
	return c.i
}
