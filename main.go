package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Counter uint64

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

func printer(done chan bool) {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-done:
			return
		case <-ticker:
			fmt.Printf("0x%016x\r", Counter)
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
