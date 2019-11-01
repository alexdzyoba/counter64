package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Counter uint64

// counter increments counter as fast as it can until it's notified that it's
// done
func counter(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			Counter++
		}
	}
}

// counterWithTick is the same as counter but accumulate counter by adding large
// increment instead of incrementing by one. This eases load on the system so
// core is not spinning in 100%.
func counterWithTick(done chan bool) {
	ticker := time.Tick(100 * time.Millisecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		select {
		case <-done:
			return
		case <-ticker:
			coeff := r.Float64() + 1 // random in [1.0, 2.0)
			Counter += uint64(float64(100000000) * coeff)
		}
	}
}

func printer(done chan bool) {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-done:
			return
		case <-ticker:
			// 1<<64-1 as decimal - 18446744073709551615
			fmt.Printf("%020d\r", Counter)
			// fmt.Printf("0x%016x\r", Counter)
		}
	}
}

func main() {
	done := make(chan bool)
	go counter(done)
	go printer(done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	done <- true
}
