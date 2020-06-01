package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kofoworola/godelve-tutorial/timer"
)

func main() {
	runner := timer.New(3, 10*time.Second, closureGreeting())
	tickerResp := runner.Begin()
	timeout := time.After(40 * time.Second)

	select {
	case err := <-tickerResp:
		if err != nil {
			log.Fatalf("error running ticker: %s", err.Error())
		}
	case <-timeout:
		log.Printf("timed out while running ticker\n")
	}
}

func closureGreeting() func() error {
	persons := 0
	return func() error {
		persons++
		fmt.Printf("Greetings to %d person(s)\n", persons)
		return nil
	}
}
