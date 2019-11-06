package counter64

// counter is the most simple counter that is incrementing in a loop as fast as
// it can
type counter struct {
	i uint64
}

func (c *counter) Count(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			c.i++
		}
	}
}

func (c *counter) Read() uint64 {
	return c.i
}
